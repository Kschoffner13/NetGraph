map object documentation: https://visgl.github.io/react-map-gl/

my-go-api/
├── cmd/
│ └── api/
│ └── main.go # Main application entry point
├── internal/
│ ├── config/ # Environment and app configuration
│ ├── database/ # Database connection and migrations
│ ├── handlers/ # HTTP request/response logic (Controllers)
│ ├── middleware/ # Auth, logging, and CORS middlewares
│ ├── models/ # Data objects and domain structures
│ └── services/ # Business logic and use cases
├── pkg/
│ └── validator/ # Custom reusable helper utilities
├── .env # Local environment configurations
├── Dockerfile # Container configuration
├── go.mod # Dependency management manifest
├── go.sum # Dependency checksums
└── Makefile # Shortcut tooling commands
