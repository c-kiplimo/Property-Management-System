# README.md - Updated for DDD Architecture
# Property Management System - Domain-Driven Design

A comprehensive property management system built with Domain-Driven Design (DDD) principles using Go, Gin, GORM, and PostgreSQL.

## Architecture Overview

This system follows Domain-Driven Design (DDD) architecture with clear separation of concerns:

### Domain Boundaries

1. **Identity Domain** - Authentication, authorization, and landlord management
2. **Property Domain** - Physical properties and units management
3. **Tenant Domain** - Tenant information and relationships
4. **Leasing Domain** - Rental agreements and occupancy
5. **Financial Domain** - Payments, rent collection, and expenses
6. **Maintenance Domain** - Work orders and property maintenance
7. **Document Domain** - File storage and management

### Architecture Layers

```
┌─────────────────────────────────────────┐
│           Interface Layer               │
│    (HTTP Handlers, Routes, Middleware)  │
├─────────────────────────────────────────┤
│          Application Layer              │
│   (Commands, Queries, Command Handlers) │
├─────────────────────────────────────────┤
│            Domain Layer                 │
│  (Entities, Value Objects, Events)     │
├─────────────────────────────────────────┤
│         Infrastructure Layer           │
│   (Repositories, External Services)    │
└─────────────────────────────────────────┘
```

## Key Features

### Domain-Driven Benefits
- **Business Focus**: Code structure mirrors business domains
- **Scalability**: Independent domain evolution
- **Maintainability**: Clear boundaries and responsibilities
- **Testability**: Isolated domains for easier testing
- **Event-Driven**: Loose coupling through domain events

### Technical Features
- **CQRS Pattern**: Separate commands and queries
- **Domain Events**: Cross-domain communication
- **Value Objects**: Rich domain modeling
- **Repository Pattern**: Data access abstraction
- **Clean Architecture**: Dependency inversion

## Project Structure

```
property-management/
├── cmd/api/main.go                    # Application entry point
├── internal/
│   ├── shared/                        # Shared kernel
│   │   ├── domain/                    # Shared domain concepts
│   │   ├── infrastructure/            # Shared infrastructure
│   │   └── application/               # Shared application layer
│   ├── domains/                       # Business domains
│   │   ├── identity/                  # Identity & Access
│   │   ├── property/                  # Property Management
│   │   ├── tenant/                    # Tenant Management
│   │   ├── leasing/                   # Lease Management
│   │   ├── financial/                 # Financial Management
│   │   ├── maintenance/               # Maintenance Management
│   │   └── document/                  # Document Management
│   └── interfaces/                    # External interfaces
└── pkg/                              # Public packages
```

Each domain follows the same structure:
```
domain/
├── domain/                           # Domain layer
│   ├── entities/                     # Domain entities
│   ├── valueobjects/                 # Value objects
│   ├── events/                       # Domain events
│   ├── repositories/                 # Repository interfaces
│   └── services/                     # Domain services
├── application/                      # Application layer
│   ├── commands/                     # Command objects
│   ├── queries/                      # Query objects
│   └── handlers/                     # Command/Query handlers
├── infrastructure/                   # Infrastructure layer
│   └── repositories/                 # Repository implementations
└── interfaces/                       # Interface layer
    └── http/                         # HTTP handlers
```

## Setup Instructions

1. **Clone and setup**:
```bash
git clone <repository-url>
cd property-management
cp .env.example .env
# Edit .env with your configuration
```

2. **Install dependencies**:
```bash
make deps
```

3. **Set up database**:
```bash
# Create PostgreSQL database
createdb property_management
```

4. **Configure Google OAuth**:
- Create OAuth credentials in Google Cloud Console
- Add credentials to .env file

5. **Run the application**:
```bash
make run
# or
make build && ./bin/property-management
```

## API Documentation

### Domain APIs

#### Identity Domain
- `GET /api/auth/login` - Get Google OAuth URL
- `GET /api/auth/callback` - Handle OAuth callback
- `GET /api/me` - Get current user

#### Property Domain
- `POST /api/properties` - Create property
- `GET /api/properties` - List properties
- `POST /api/properties/:id/units` - Add unit

#### Tenant Domain
- `POST /api/tenants` - Create tenant
- `GET /api/tenants` - List tenants

#### Leasing Domain
- `POST /api/leases` - Create lease
- `GET /api/leases` - List leases

#### Financial Domain
- `POST /api/payments` - Record payment
- `GET /api/payments` - List payments
- `POST /api/expenses` - Record expense

## Development Commands

```bash
make build        # Build application
make run          # Run application  
make test         # Run tests
make test-coverage # Run tests with coverage
make fmt          # Format code
make lint         # Lint code
make clean        # Clean build artifacts
```

## Domain Events Flow

Example: When a lease is signed
```
1. Leasing Domain → LeaseSignedEvent
2. Financial Domain → Generates payment schedule
3. Property Domain → Marks unit as occupied
4. Document Domain → Creates lease document record
```

## Testing Strategy

- **Unit Tests**: Test individual domain entities and value objects
- **Integration Tests**: Test repository implementations
- **Application Tests**: Test command/query handlers
- **API Tests**: Test HTTP endpoints

## Migration from Previous Architecture

See `docs/migration-guide.md` for detailed steps to migrate from the previous flat structure.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Follow DDD principles
4. Write tests for new domains/features
5. Submit a pull request

## License

This project is licensed under the MIT License.
