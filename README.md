# forbidden-memories-duel-online
The classic forbidden memories videogame for ps1 but with an amazing new feature; MULTIPLAYER ONLINE P2P. So then current rules for Yu Gi Oh! do not apply for this game simulation.

## Project Structure
```bash
/go-modules/
├── cmd/                    # Entry points for applications
├── internal/               # Private Application code
│   ├── engine/             # Game Engine implementation
│   │   ├── domain/         # Core business logic
│   │   ├── integration/    # Integration tests with other modules
│   │   └── transport/      # gRPC Communication layer
│   └── persistence/        # Data Persistence implementation
│       ├── domain/         # Core business logic
│       ├── integration/    # Integration tests with other modules
│       ├── repository/     # Data access layer
│       └── transport/      # gRPC Communication layer
└── pkg/                    # Public Application code
    ├── models/             # Data Models
    ├── proto/              # Protobuf definitions
    └── utils/              # Shared utilities
```

## Backend Components

- **Achievements and Rewards System (Python):** Manages player achievements and rewards.
  > Python's extensive data analysis libraries make it ideal for tracking and analyzing player progress and behavior.

- **WebSocket Handler (Rust):** Handles real-time communication between players.
  > Rust's performance and memory safety make it perfect for handling high-throughput, low-latency WebSocket connections.

- **Card and Deck Validation System (Rust):** Manages cards and verifies deck validity.
  > Rust's strict type system and performance characteristics ensure reliable and fast card/deck validation.

- **Matchmaking System (Go):** Matches players for duels.
  > Go's excellent concurrency model with goroutines and channels is perfect for managing multiple concurrent matchmaking requests efficiently.

- **Game Engine (Go):** Core game logic and managing flow.
  > Go's simplicity, strong concurrency support, and fast execution make it ideal for core game logic.

- **Data Persistence (Go):** Manages storage and retrieval of game data.
  > Go's strong standard library and excellent database drivers provide flexible and efficient data persistence.

## Game Entities

- **Card:** Represents a monster or magic in the game.
- **Player:** The human playing the game.
- **Board:** The playing area where cards are placed.
- **Deck:** A set of cards a player can use.
- **Game:** Represents the battle between 2 players.
- **Turn:** Manages the turn cycle in the game.
- **Event:** Handles events occurring during the game.
- **Engine:** Manages the game's logic and flow.

## Run Unit Tests
By default, tests within the same package are executed sequentially. Tests from different packages are executed in parallel. To force parallelism between test functions within the same package, use `t.Parallel()`.

```bash
cd go-modules/
go test -cover -coverprofile=./coverage.out ./pkg/models/ -v # Verbose output
go test -cover -coverprofile=./coverage.out ./pkg/models/ # Concise output
go tool cover -html=./coverage.out

## UI
We use Vercel, Redux, Next, Tailwind and GreenSock