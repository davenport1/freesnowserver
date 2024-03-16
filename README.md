# FreeSnow Backcountry and Resort Web Server 

## Overview

This server provides information on ski resorts, snow reports, weather data, routes, and 
avalanche forecasts for popular backcountry zones. It aims to serve both enthusiasts 
seeking ski resort information and those planning backcountry adventures. The server is pumped information
from a serverless scraper using the scrapy libraries which scrapes resorts and their APIs as well as 
avalanche centers and API calls to NOAA.

## Features

- **Ski Resort Information**: Get details about ski resorts, including location, trail status, lift status, and trail maps.
- **Snow Reports**: Access up-to-date snow reports for ski resorts.
- **Weather Data**: Retrieve current weather data and forecasts for ski resorts and backcountry zones.
- **Routes**: Explore recommended routes for backcountry skiing and snowboarding.
- **Avalanche Forecasts**: Stay informed about avalanche forecasts and safety recommendations for backcountry areas.

## Usage

To run the server, follow these steps:

1. **Install Dependencies**: Make sure you have Go installed on your system. If not, download and install it from [here](https://golang.org/dl/).

2. **Clone the Repository**: Clone this repository to your local machine.

```bash 
git clone <repository_url>
```

4. **Navigate to the Project Directory**: Change directory to the cloned repository.

```bash
cd <repository_directory>
```

5. **Build the Server**: Build the server using the following command.

```bash
go build
```

5. **Run the Server**: Execute the server binary to start the server.

```bash
./<binary_name>
```
**or run in development mode**
```bash
go run .
```

**Make sure to add the FREESNOW_DB_DSN environment variable to your system environment variables,
for example with mac osx:**
```bash
export FREESNOW_DB_DSN=postgres://<postgres-user>:<postgres-password>@localhost:5433/freesnow_db?sslmode=disable
```

6. **Access the Server**: Once the server is running, you can access it through the 
specified port (default port: 8080).

```http
http://localhost:8080
```

## Configuration

You can configure the server using the following command line flags and environment variables:

    CLI -p -port: Specify the port number for the server (default: 8080).
    CLI -e -env: Set the environment (dev, stage, prod) for the server (default: dev).
    Environment Varialbe "FREESNOW_DB_DSN": Set the PostgreSQL database DSN for connecting to the database.

## Dependencies

    Go: The server is written in Go, a programming language.
    SQL Database: PostgreSQL database is used for storing ski resort and weather data.
    Libraries: The server relies on the pq postgres library for Go

## Contributing

Contributions to this project are welcome. If you find any bugs, have feature requests,
or want to contribute code, feel free to submit a pull request or open an issue on GitHub.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
