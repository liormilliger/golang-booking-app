
# Go Booking App – Learning Project

This repository is a **hands-on Golang learning project** that evolves step-by-step into a functional command-line booking application.

Each **chapter in the course corresponds to a Git commit**, allowing you to:
- Learn Go incrementally
- Inspect changes using `git diff`
- Understand how concepts build on top of each other

---

## Project Overview

The application simulates a **conference ticket booking system** and is used as a vehicle to learn core Go concepts, including:

- Go project structure and modules
- Variables, constants, and data types
- User input handling
- Control flow (if/else, loops, switch)
- Arrays, slices, maps, and structs
- Functions and return values
- Packages and scope rules
- Concurrency with goroutines and synchronization

---

## Repository Structure

```

.
├── main.go
├── helper/          # Helper package (introduced later in the course)
│   └── helper.go
├── go.mod
└── README.md

```

> The structure evolves during the course. Use Git history to follow the progression.

---

## Learning Flow (Mapped to Git Commits)

### 1. Project Initialization
- Create `main.go`
- Initialize Go module using:
```
  go mod init booking-app
```

* Understand `package main` and the `main()` entry point

---

### 2. Variables and Constants

* Variable declaration styles:

  ```
  var name string
  name := "Go Conference"
  ```
* Constants vs variables
* Basic numeric and string types

---

### 3. Formatting Output

* Using `fmt.Printf` with placeholders:

  ```
  fmt.Printf("Tickets left: %v\n", remainingTickets)
  ```
* Printing variable types using `%T`

---

### 4. User Input

* Reading input with `fmt.Scan`
* Using pointers (`&`) to store user input

---

### 5. Booking Logic

* Ticket reduction logic
* Type consistency (`uint` for ticket counts)
* Preventing overbooking

---

### 6. Arrays and Slices

* Fixed-size arrays
* Dynamic slices
* Appending elements
* Why slices are preferred in Go

---

### 7. Loops

* Infinite loops using `for {}`
* Conditional loops:

  ```
  for remainingTickets > 0 {}
  ```
* Iteration control using `break` and `continue`

---

### 8. For-Each (Range) Loops

* Iterating over slices
* Ignoring unused variables with `_`
* String manipulation using `strings.Fields`

---

### 9. Conditional Logic

* `if`, `else if`, `else`
* Boolean expressions
* Validating user input
* Preventing invalid bookings

---

### 10. Switch Statements

* City-based booking logic
* Grouping multiple cases
* Default case handling

---

### 11. Functions

* Extracting logic into reusable functions
* Function parameters
* Returning values
* Clean separation of concerns

---

### 12. Package-Level Variables

* Sharing state across functions
* Difference between local and package scope
* When to avoid passing variables explicitly

---

### 13. Packages and Code Organization

* Splitting code into multiple files
* Creating custom packages
* Importing internal packages using module path
* Exporting functions via capitalization

---

### 14. Maps

* Using `map[string]string` to store user data
* Initializing maps with `make`
* Converting numeric types using `strconv`

---

### 15. Structs

* Creating custom data types with `struct`
* Replacing maps with strongly-typed structs
* Cleaner and safer data modeling

---

### 16. Concurrency (Goroutines)

* Running functions concurrently using `go`
* Non-blocking ticket sending simulation
* Understanding Go’s concurrency model

---

### 17. Synchronization with WaitGroups

* Using `sync.WaitGroup`
* Preventing premature application exit
* Coordinating goroutines safely

---

## How to Run the App

```
go run .
```

> This runs all files in the module directory.

---

## Key Concepts Reinforced

* Go is **explicit**, **type-safe**, and **concurrency-first**
* Prefer **composition and simplicity**
* Use **Git history** as a learning tool
* Structure code early to avoid refactoring pain later

---

## Who This Project Is For

* Developers new to Go
* DevOps / SRE engineers learning Go fundamentals
* Anyone who prefers **learning by building**
* Engineers who want to understand Go idioms, not just syntax

---

## Suggested Learning Method

1. Checkout the first commit
2. Run the code
3. Read the changes
4. Move to the next commit
5. Use `git diff` aggressively


