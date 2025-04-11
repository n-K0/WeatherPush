package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

const version = "1.0.2"

type WeatherResponse struct {
	List []struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
		Dt_txt string `json:"dt_txt"`
	} `json:"list"`
}

const openWeatherAPIURL = "https://api.openweathermap.org/data/2.5/forecast"
const defaultPushURL = "https://push.aut-o-matic.com/message"

func getEnvOrArg(envName, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		value = defaultValue
	}

	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			arg := os.Args[i]
			if arg == "--"+envName {
				if i+1 < len(os.Args) {
					value = os.Args[i+1]
				}
				break
			}
		}
	}

	return value
}

func getWeatherCity(apiKey, city, apiURL string) (string, error) {
	client := resty.New()

	resp, err := client.R().
		SetQueryParam("q", city).
		SetQueryParam("appid", apiKey).
		SetQueryParam("units", "metric").
		SetQueryParam("lang", "fr").
		SetQueryParam("cnt", "4").
		Get(apiURL)

	if err != nil {
		return "", fmt.Errorf("failed to get weather: %s", err)
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(resp.Body(), &weatherResponse)
	if err != nil {
		return "", fmt.Errorf("failed to parse weather response: %s", err)
	}

	return formatWeatherInfo(weatherResponse), nil
}

func formatWeatherInfo(weatherResponse WeatherResponse) string {
	var description string
	for i := 0; i < len(weatherResponse.List); i++ {
		dtTxt, err := time.Parse("2006-01-02 15:04:05", weatherResponse.List[i].Dt_txt)
		if err != nil {
			return fmt.Sprintf("Error parsing date: %s", err)
		}
		weatherHour := dtTxt.Format("02/01 15:04")
		description += fmt.Sprintf("%s - %s - %.2f °C\n", weatherHour, weatherResponse.List[i].Weather[0].Description, weatherResponse.List[i].Main.Temp)
	}

	log.Print("\n" + description)
	return description
}

func sendPushNotif(title, message, pushKey string) error {
	type Body struct {
		Title   string `json:"title"`
		Message string `json:"message"`
	}
	body := Body{
		Title:   title,
		Message: message,
	}
	marshalled, err := json.Marshal(body)

	if err != nil {
		return fmt.Errorf("failed to marshal notification body: %s", err)
	}

	pushURL := getEnvOrArg("PUSH_URL", defaultPushURL)
	req, err := http.NewRequest("POST", pushURL+"?token="+pushKey, bytes.NewReader(marshalled))
	if err != nil {
		return fmt.Errorf("failed to build request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %s", err)
	}

	log.Print(res.StatusCode)
	log.Print("###################")

	return nil
}

func main() {
	apiKey := getEnvOrArg("OPENWEATHER_API_KEY", "")
	city := getEnvOrArg("CITY", "")
	pushKey := getEnvOrArg("PUSH_KEY", "")

	///////////////////////////
	start := time.Now()
	log.Printf("-- Start Script\n")
	///////////////////////////

	weatherInfo, err := getWeatherCity(apiKey, city, openWeatherAPIURL)
	if err != nil {
		errorMsg := fmt.Sprintf("Error getting weather: %s", err)
		log.Println(errorMsg)
		err := sendPushNotif("Weather Error", errorMsg, pushKey)
		if err != nil {
			log.Printf("Error sending push notification: %s", err)
		}
	} else {
		date := time.Now()
		weatherTitle := "Weather : " + city + " - " + date.Format("2006-01-02")

		err := sendPushNotif(weatherTitle, weatherInfo, pushKey)
		if err != nil {
			errorMsg := fmt.Sprintf("Error sending push notification: %s", err)
			log.Println(errorMsg)
			sendPushNotif("Push Notification Error", errorMsg, pushKey)
		}
	}

	///////////////////////////
	elapsed := time.Since(start)
	log.Printf("-- Elapsed time : %v\n", elapsed)
	///////////////////////////
}
