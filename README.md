# Universal Gravitation Simulator

This project is a Universal Gravitation Simulator implemented in Go, following the Ports and Adapters (Hexagonal) architecture. It models the gravitational interactions between bodies, allowing extensibility and easy integration with different interfaces and frameworks.

## Features

- Simulates gravitational forces between multiple bodies
- Modular design using Ports and Adapters architecture
- Written in Go for performance and simplicity
- Easily extendable for new input/output adapters

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone <repo-url>
   cd <repo-folder>
   ```
   
1. **Install dependencies:**
   ```bash
   go mod tidy
   ```
   
1. **Run the simulator:**
   ```bash
   go run main.go
    ```
## Project Structure
- domain/ - Core business logic and entities 
- adapters/ - Input/output adapters (e.g., CLI, web)
- ports/ - Interfaces for communication between core and adapters 
- main.go - Application entry point

## Architecture
This project uses the Ports and Adapters (Hexagonal) architecture to separate business logic from technical details, making it easy to test and maintain.
   
## License
MIT