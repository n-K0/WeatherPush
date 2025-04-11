# WeatherPush

WeatherPush est une application Go qui permet de recevoir des notifications push avec les prévisions météorologiques pour une ville spécifique.

WeatherPush is a Go application that allows you to receive push notifications with weather forecasts for a specific city.

## Fonctionnalités / Features

- Récupération des prévisions météorologiques via l'API OpenWeatherMap
- Envoi de notifications push via le service push.aut-o-matic.com
- Affichage des températures en degrés Celsius
- Prévisions sur 4 périodes
- Interface en français
- Support des variables d'environnement et des arguments en ligne de commande

- Weather forecast retrieval via OpenWeatherMap API
- Push notification sending via push.aut-o-matic.com service
- Temperature display in Celsius
- 4-period forecasts
- French interface
- Environment variables and command line arguments support

## Prérequis / Prerequisites

- Go 1.x
- Une clé API OpenWeatherMap / An OpenWeatherMap API key
- Une clé d'API push.aut-o-matic.com / A push.aut-o-matic.com API key

## Installation / Installation

1. Clonez le dépôt / Clone the repository:
```bash
git clone https://github.com/votre-utilisateur/WeatherPush.git
cd WeatherPush
```

2. Installez les dépendances / Install dependencies:
```bash
go mod download
```

## Configuration / Configuration

Vous pouvez configurer l'application de deux façons / You can configure the application in two ways:

### Via les variables d'environnement / Using environment variables

```bash
export OPENWEATHER_API_KEY="votre_clé_api_openweather"
export CITY="votre_ville"
export PUSH_KEY="votre_clé_push"
export PUSH_URL="https://votre-service-push.com/message"  # Optionnel, par défaut : https://push.aut-o-matic.com/message / Optional, default: https://push.aut-o-matic.com/message
```

### Via les arguments en ligne de commande / Using command line arguments

```bash
./WeatherPush --OPENWEATHER_API_KEY votre_clé_api_openweather --CITY votre_ville --PUSH_KEY votre_clé_push --PUSH_URL https://votre-service-push.com/message
```

## Utilisation / Usage

Exécutez simplement le binaire / Simply run the binary:

```bash
./WeatherPush
```

L'application va / The application will:
1. Récupérer les prévisions météorologiques pour la ville spécifiée / Retrieve weather forecasts for the specified city
2. Formater les informations (date, heure, température, description) / Format the information (date, time, temperature, description)
3. Envoyer une notification push avec ces informations / Send a push notification with this information

## Format de la notification / Notification Format

La notification contiendra / The notification will contain:
- Un titre avec le nom de la ville et la date / A title with the city name and date
- Un message avec les prévisions sur 4 périodes, incluant / A message with 4-period forecasts, including:
  - Date et heure / Date and time
  - Description météorologique / Weather description
  - Température en degrés Celsius / Temperature in Celsius

## Docker

Le projet inclut un Dockerfile pour faciliter le déploiement / The project includes a Dockerfile for easy deployment:

```bash
docker build -t weatherpush .
docker run -e OPENWEATHER_API_KEY=votre_clé -e CITY=votre_ville -e PUSH_KEY=votre_clé weatherpush
```

## Version / Version

Version actuelle / Current version: 1.0.2

## Licence / License

Ce projet est sous licence MIT / This project is licensed under MIT.

## Auteur / Author

n-K0_

## Remerciements / Acknowledgments

- OpenWeatherMap pour l'API météorologique / OpenWeatherMap for the weather API
- push.aut-o-matic.com pour le service de notifications push / push.aut-o-matic.com for the push notification service 