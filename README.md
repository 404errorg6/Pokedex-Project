# Pokedex-Project

## Overview
The **Pokedex-Project** is a command-line application written in Go that simulates a Pokémon Trainer's experience. It provides various commands to explore, catch, inspect, and manage Pokémon, as well as navigate through a virtual map.

## Features
- **Explore**: Discover new Pokémon in different areas.
- **Catch**: Attempt to catch Pokémon you encounter.
- **Inspect**: View details about the Pokémon you've caught.
- **Map Navigation**: Navigate through a virtual Pokémon world.
- **Pokedex Management**: Keep track of all the Pokémon you've encountered and caught.
- **Help Command**: Get a list of available commands and their descriptions.

## Project Structure
- `main.go`: Entry point of the application.
- `repl.go`: Handles the Read-Eval-Print Loop (REPL) for user interaction.
- `command_*.go`: Implements various commands like `catch`, `explore`, `inspect`, etc.
- `helper.go`: Contains utility functions used across the project.
- `internal/`: Contains internal logic, such as caching for Pokémon data.

## Getting Started

### Prerequisites
- Go 1.20 or later installed on your system.

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/404errorg6/Pokedex-Project.git
   cd Pokedex-Project
   ```
2. Build the project:
   ```bash
   go build
   ```

### Running the Application
Run the application using:
```bash
go run main.go
```

### Commands
Once the application starts, you can use the following commands:
- `explore`: Discover new Pokémon.
- `catch`: Attempt to catch a Pokémon.
- `inspect`: View details about a specific Pokémon.
- `map`: Navigate through the Pokémon world.
- `pokedex`: View your Pokedex.
- `help`: Display a list of commands.
- `exit`: Exit the application.

## Testing
Run the tests using:
```bash
go test ./...
```

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

## License
This project is licensed under the MIT License.
