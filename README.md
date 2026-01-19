# Wallet Transfer Service (Go)

A simple **wallet-to-wallet money transfer service** written in Go, designed with a strong focus on **dependency injection**, **clean architecture**, and **safe money handling**.

This project demonstrates how to structure business logic independently from storage concerns while avoiding common financial pitfalls such as floating-point precision errors.

---

## ✨ Features

- Wallet-to-wallet money transfers
- Dependency Injection via interfaces
- Safe money representation (no floating-point arithmetic)
- Clean separation of concerns
- Easily testable business logic
- Edge case handling (insufficient funds, self-transfers, repository failures)

---

## 🧱 Architecture Overview

The core business logic resides in the `WalletService`, which depends on a `WalletRepository` **interface** rather than a concrete storage implementation.

This design enables:
- Decoupling of business logic from persistence
- Easy unit testing with mocks or in-memory repositories
- Swappable storage implementations without modifying service logic

---

## 🔌 Dependency Injection

The repository is injected into the service using **constructor injection**:

```go
func NewWalletService(repo WalletRepository) *WalletService
````

### Benefits of This Approach

* Service logic is independent of storage
* Simplified unit testing
* Clear dependency boundaries
* Improved maintainability and scalability

---

## 💰 Money Representation

Money is represented using a dedicated `Money` type backed by `int64`, storing values in the smallest currency unit (e.g., cents).

### Why This Matters

* Prevents floating-point precision and rounding errors
* Disallows invalid transfers (zero or negative amounts)
* Centralizes money validation and arithmetic in the domain layer

---

## 🧪 Testing Strategy

The project includes unit tests that cover:

* Successful wallet transfers
* Transfers failing due to insufficient funds
* Transfers failing due to repository errors
* Prevention of self-transfers (edge case)

Tests rely on mock or in-memory repository implementations to ensure deterministic and isolated behavior.

---

## 📁 Project Structure (Example)

```
.
├── domain
│   ├── money.go
│   ├── wallet.go
│   └── repository.go
├── service
│   └── wallet_service.go
├── repository
│   └── in_memory_wallet.go
├── tests
│   └── wallet_service_test.go
└── README.md
```

---

## 🚀 Getting Started

1. Clone the repository
2. Run tests:

   ```bash
   go test ./...
   ```
3. Extend with your preferred persistence layer (SQL, Redis, etc.)

---

## 📌 Design Principles Used

* Clean Architecture
* Dependency Inversion Principle (DIP)
* Explicit domain modeling
* Fail-fast validation
* Test-driven design

---

```
```
