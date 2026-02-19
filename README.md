# 🍴 Golang Food Order CLI Application

A Command Line Interface (CLI) based food ordering application built using the Go programming language. This project is a migration from a JavaScript application to GoLang, implementing fundamental to advanced features within the Go ecosystem.

## 🚀 Key Features

- **Asynchronous Data Loading**: Fetches food menu data directly from the API/GitHub Raw in real-time.
- **Food Search**: Search for menus based on specific keywords.
- **Cart System**: Add items with dynamic variation options and pricing.
- **Checkout Process (Concurrency)**: Utilizes Goroutines and Channels to simulate asynchronous transaction processes.
- **Transaction History**: Records all successful transactions performed while the application is running.
- **Error Handling**: Implementation of `panic` and `recover` to prevent application crashes during fatal data connection errors.

## 🛠️ Implemented GoLang Concepts

This application utilizes 11 core GoLang topics:
1.  **Data Types & Variables**: Use of strict typing for data security.
2.  **Functions**: Decomposition of code logic into modular functions.
3.  **Looping**: Iterating through menu data and the main menu display using `for`.
4.  **Arrays, Slices, & Structs**: Storage for menu data, cart items, and transactions.
5.  **Modules**: Organizing code into `services` and `utils` packages.
6.  **Defer & Exit**: Safely closing HTTP response bodies and exiting the application.
7.  **Panic & Recover**: Runtime error handling within the `loadData` function.
8.  **Pointers**: Memory efficiency by referencing the addresses of selected menu data.
9.  **Methods**: Business logic attached to the `FoodService` struct.
10. **Interfaces**: Contract abstraction for cart management via `CartManager`.
11. **Goroutines & Channels**: Synchronizing the Checkout process to run concurrently.

## 📂 Folder Structure

```text
koda-b6-golang/
├── main.go             # Application entry point & UI Logic
├── services/
│   └── cart_service.go # Business logic for cart & transactions
└── utils/
    └── models.go       # Struct definitions & data models
```

## 💻 How to Run

1. **Clone the Repository**
   ```bash
   git clone https://github.com/Vincentius31/koda-b6-golang
   cd koda-b6-golang
   ```

2. Initialize the Module (If go.mod is not yet present)
   ```bash
   go mod init koda-b6-golang
   ```

3. Run the application
   ```
   go run main.go
   ```