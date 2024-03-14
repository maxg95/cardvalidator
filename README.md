# API Server

API server implemented in Go.

## Usage

### Running the Server

To run the server locally, execute the following command:
make run/api

To build the server binary, execute the following command:
make build/api

## Docker
![Images](https://drive.google.com/drive/folders/1o6SjALp6WVfXzCy8CqL9uthHcem_BW2A?usp=share_link)

## Dependencies

github.com/joeljunstrom/go-luhn - Package for Luhn algorithm
github.com/julienschmidt/httprouter - HTTP request router

## Configuration

port (default: 4000): Port number for the API server.