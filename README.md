# Interactsh API Server

This is a simple Go web server that serves as an API for interacting with the interactsh tool. The server provides endpoints to fetch URLs and interactions from an interactsh server.

## Requirements

- Go (Golang)
- interactsh-client (Make sure it's installed and accessible in your system)

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/your_username/interactsh-api-server.git
   ```

2. Navigate to the project directory:

   ```bash
   cd interactsh-api-server
   ```

3. Build and run the server:

   ```bash
   go build
   ./interactsh-api-server
   ```

4. The server will start running on port `8080`.

## Usage

### 1. Fetch URL

- **Endpoint:** `/api/getURL`
- **Method:** GET
- **Description:** Fetches a URL from the interactsh server.
- **Response:** Returns the URL in the format `[INF] **url**`.

### 2. Fetch Interactions

- **Endpoint:** `/api/getInteractions`
- **Method:** GET
- **Description:** Fetches interactions from the interactsh server based on the previously fetched URL.
- **Response:** Returns interactions in the format `[INF] Interactions:` followed by each interaction's details.

## Additional Information

- The server automatically cleans up session data for inactive sessions.
- Make sure to set proper authorization headers when making requests to the server.
- The server uses the Gin framework for routing and handling HTTP requests.
- Session data is stored in memory, so it will be lost if the server restarts.

[![Screenshot 1](screenshots/Screenshot%202024-02-25%20172022.png)](screenshots/Screenshot%2024-02-25%20172022.png)


[![Screenshot 2](screenshots/Screenshot%202024-02-25%20170156.png)](screenshots/Screenshot%202024-02-25%20170156.png)


## Troubleshooting

- If you encounter any issues with the server, check the console output for error messages.
- Ensure that the interactsh-client is properly installed and accessible in your system environment.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.# interact
