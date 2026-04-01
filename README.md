# EchoDelta

A simple Go web server that converts text to Morse code.

## Features

- HTTP endpoint to convert text to Morse code
- Supports letters A-Z, numbers 0-9, and spaces
- Case-insensitive input

## Usage

1. Build the project:
   ```
   go build
   ```

2. Run the server:
   ```
   ./echodelta
   ```

3. Make a request:
   ```
   curl "http://localhost:8080/morse?text=HELLO"
   ```

   Or in browser: `http://localhost:8080/morse?text=HELLO`

## Example

Input: `HELLO`
Output: `.... . .-.. .-.. ---`

## API

- **GET /morse?text=<text>**
  - Converts the provided text to Morse code
  - Returns plain text response

## Requirements

- Go 1.21 or later