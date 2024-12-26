# Go LLM API

A simple REST API built with Go and Gin framework that interfaces with Groq's LLM API to process text requests.

## Features

- RESTful endpoint for LLM interactions
- Uses Groq's API for text generation
- Environment-based configuration
- JSON request/response handling

## Prerequisites

- Go 1.x
- Groq API key

## Setup

1. Clone the repository:
   ```
   git clone git@github.com:nikitazuevblago/LLMApi.git
   cd LLMApi
   ```

2. Initialize Go module and install dependencies:
   ```
   go mod init github.com/nikitazuevblago/LLMApi
   go mod tidy
   ```

3. Create a `.env` file in the root directory with:
   ```
   GROQ_API_KEY=your_api_key_here
   GROQ_MODEL=your_preferred_model
   ```

## Running the Application

Start the server:
```
go run main.go
```

The server will start on `localhost:8080`

## API Endpoints

### POST /LLMResponse

Send a message to the LLM and get a response.

Request body:
```json
{
    "message": "Your message here"
}
```

Response:
```json
{
    "response": "LLM generated response"
}
```

## Environment Variables

- `GROQ_API_KEY`: Your Groq API key
- `GROQ_MODEL`: The Groq model to use for completions