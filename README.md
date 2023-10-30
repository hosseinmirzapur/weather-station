# Weather Station IoT Project

This is a weather station IoT project developed using Golang. The project collects data from various sensors to monitor and display real-time weather information.

## Features

- Temperature Monitoring: Displays the current temperature and temperature trends over time.
- Humidity Monitoring: Shows the current humidity level and provides insights into humidity variations.
- Atmospheric Pressure Monitoring: Displays the atmospheric pressure and historical pressure data.
- Weather Forecast: Retrieves weather forecast data from an API and displays it on the interface.
- Wind Speed and Direction: Displays the current wind speed and direction, including wind gusts and trends.
- Rainfall Measurement: Measures and displays rainfall data, including rate and historical records.
- UV Index Monitoring: Provides information on UV radiation levels and safety recommendations.
- Web Interface: A user-friendly web-based interface to view real-time and historical weather data.
- Mobile App Integration: Allows users to monitor weather data and receive notifications remotely.
- Data Logging and Analytics: Stores collected weather data for analysis and visualization.
- Notifications and Alerts: Sends notifications or alerts based on specific weather conditions.
- Geolocation and Mapping: Displays the weather station's location on a map.

## Requirements

- Golang (version 1.15 and higher)

## Installation

1. Clone the repository:

```
git clone https://github.com/hosseinmirzapur/weather-station.git
```

2. Install the dependencies:

```
cd weather-station
go mod download
```

3. You can run this code on any Raspberry Pi, PC or your system of choice

4. Creata a .env file based on the .env.example file which is included in the project root directory and place your own values

## Usage

1. Run the application:

```
go run main.go

# or go build && ./weatherstation   (weatherstation.exe => on windows)
```

2. Access the weather station interface by opening a web browser and navigating to `http://localhost:3000`.

3. You can also change the default port to whatever you want in the server.go file included in "server" folder

## Contributing

Contributions are welcome! If you would like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with descriptive messages.
4. Push your changes to your forked repository.
5. Submit a pull request to the main repository.

## License

MIT

## Contact

For any questions or inquiries, please contact Hossein Mirzapur at [hosseinmirzapur@gmail.com].

Feel free to customize the README file based on your specific project requirements, additional features, or team information.
