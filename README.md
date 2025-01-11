# forbidden-memories-duel-online
The classic forbidden memories for ps1 game but online p2p

## Project Structure
```go
/go-modules/
├── cmd/                    # Entry points for applications
├── internal/              # Private application code
│   ├── engine/           # Game Engine implementation
│   └── persistence/      # Data Persistence implementation
└── pkg/                  # Public shared code
    ├── models/           # Shared data models
    ├── proto/            # Protocol buffer definitions
    └── utils/            # Shared utilities
```

## Backend Components

- **Achievements and Rewards System (Python):** Manages player achievements and rewards.
  > Python's extensive data analysis libraries make it ideal for tracking and analyzing player progress and behavior.

- **WebSocket Handler (Rust):** Handles real-time communication between players.
  > Rust's performance and memory safety make it perfect for handling high-throughput, low-latency WebSocket connections.

- **Matchmaking System (Go):** Matches players for duels.
  > Go's excellent concurrency model with goroutines and channels is perfect for managing multiple concurrent matchmaking requests efficiently.

- **Game Engine (Go):** Core game logic, managing flow and rules.
  > Go's simplicity, strong concurrency support, and fast execution make it ideal for core game logic.

- **Card and Deck Validation System (Rust):** Manages cards and verifies deck validity.
  > Rust's strict type system and performance characteristics ensure reliable and fast card/deck validation.

- **Rules System (C#):** Defines and applies game rules.
  > C#'s strong OOP features and LINQ capabilities make it excellent for implementing complex rule systems and their interactions.

- **Data Persistence (Go):** Manages storage and retrieval of game data.
  > Go's strong standard library and excellent database drivers provide flexible and efficient data persistence.
