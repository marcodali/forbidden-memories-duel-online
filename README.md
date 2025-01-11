# forbidden-memories-duel-online
The classic forbidden memories for ps1 game but online p2p

## Project Structure
```go
/go-modules/
├── cmd/                    # Entry points for applications
├── internal/              # Private application code
│   ├── engine/           # Game Engine implementation
│   │   ├── domain/      # Core business logic
│   │   ├── integration/ # External service integration
│   │   └── transport/   # Communication layer
│   └── persistence/      # Data Persistence implementation
│       ├── domain/      # Data models
│       ├── integration/ # Database integration
│       ├── repository/  # Data access layer
│       │   ├── redis/   # Redis implementation
│       │   ├── mongodb/ # MongoDB implementation
│       │   ├── dynamodb/# DynamoDB implementation
│       │   └── wrapper/ # Database wrapper interfaces
│       └── transport/   # Data service communication
└── pkg/                  # Public shared code
    ├── models/           # Shared data models
    ├── proto/           # Protocol buffer definitions
    │   ├── engine/     # Game engine protos
    │   └── persistence/# Data persistence protos
    └── utils/           # Shared utilities
        ├── logger/     # Logging utilities
        └── config/     # Configuration management
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
