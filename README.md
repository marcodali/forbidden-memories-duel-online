# forbidden-memories-duel-online
The classic forbidden memories videogame for ps1 but with an amazing new feature; MULTIPLAYER ONLINE P2P. So then current rules for Yu Gi Oh! do not apply for this game simulation.

## Project Structure
```bash
/go-modules/
├── cmd/                    # Entry points for applications
├── internal/               # Private application code
│   ├── engine/             # Game Engine implementation
│   │   ├── domain/         # Core business logic
│   │   ├── integration/    # External service integration
│   │   └── transport/      # Communication layer
│   └── persistence/        # Data Persistence implementation
│       ├── domain/         # Core business logic
│       ├── integration/    # Database integration tests
│       ├── repository/     # Data access layer
│       └── transport/      # Data service communication
└── pkg/                    # Public shared code
    ├── models/             # Shared data models
    ├── proto/              # Protocol buffer definitions
    └── utils/              # Shared utilities
```

## Backend Components

- **Achievements and Rewards System (Python):** Manages player achievements and rewards.
  > Python's extensive data analysis libraries make it ideal for tracking and analyzing player progress and behavior.

- **Rules System (Python):** Defines and applies game rules.
  > Python's flexibility and readability make it excellent for implementing and maintaining complex game rules.

- **WebSocket Handler (Rust):** Handles real-time communication between players.
  > Rust's performance and memory safety make it perfect for handling high-throughput, low-latency WebSocket connections.

- **Matchmaking System (Go):** Matches players for duels.
  > Go's excellent concurrency model with goroutines and channels is perfect for managing multiple concurrent matchmaking requests efficiently.

- **Game Engine (Go):** Core game logic, managing flow and rules.
  > Go's simplicity, strong concurrency support, and fast execution make it ideal for core game logic.

- **Card and Deck Validation System (Rust):** Manages cards and verifies deck validity.
  > Rust's strict type system and performance characteristics ensure reliable and fast card/deck validation.

- **Data Persistence (Go):** Manages storage and retrieval of game data.
  > Go's strong standard library and excellent database drivers provide flexible and efficient data persistence.

## Game Entities

- **Card:** Represents a card in the game.
- **Player:** Represents a player in the game.
- **Board:** The playing area where cards are placed.
- **Engine:** Manages the game's logic and flow.
- **Deck:** A set of cards a player can use.
- **Game:** Represents an instance of the game.
- **Turn:** Manages the turn cycle in the game.
- **Event:** Handles events occurring during the game.

## Run unit tests

```bash
cd go-modules/
go test -cover -coverprofile=./coverage.out ./pkg/models/ -v #verbose output
go test -cover -coverprofile=./coverage.out ./pkg/models/ #silence mode
go tool cover -html=./coverage.out
```
