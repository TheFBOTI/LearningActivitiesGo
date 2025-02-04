# Learning Activities Go
A collection of activities ranging from beginner to intermediate for different languages, refer to ReadME for further information.

More to be added in time...
# Phase 1: Foundational Knowledge (2-4 weeks)
1.	Objective: Get comfortable with Go's syntax, tools, and basic features.
2.	Activities:
	  - Hello World CLI App
          -	Write a command-line tool that prints "Hello, World!" and takes optional flags (e.g., --name).
           - Skills: Basic syntax, flag package usage.
      - Temperature Converter
           - Create a program to convert temperatures between Fahrenheit, Celsius, and Kelvin.
           - Skills: I/O, conditionals, and simple math.
      - Basic File Manipulation
           - Write a script to read a text file and count the frequency of words.
           - Skills: File handling, string manipulation, and maps.
      - Go Playground Programs
           - Solve coding puzzles like Fibonacci sequence, palindrome checker, or prime number generator.
           - Skills: Loops, functions, recursion.
      - Simple Calculator
           - Build a CLI calculator that supports addition, subtraction, multiplication, and division.
           - Skills: Functions, error handling (e.g., division by zero).
      - JSON Parser
          - Write a program to read and parse JSON data (e.g., a list of users) and display it in a readable format.
          - Skills: encoding/json, structs, unmarshaling.
      - Basic HTTP Server
          - Create a simple web server with Go that serves static content like "Hello, Go!" on a specific port.
          - Skills: net/http, basic routing.
      - FizzBuzz
          - Solve the classic FizzBuzz problem with added flexibility for customizable numbers and phrases.
          - Skills: Logic, conditionals, and loops.
      - Random Quote Generator
           - Build a CLI app that fetches a random quote from a list and displays it.
           - Skills: Slices, randomization, and I/O  

3.	Tools to Learn: 
      - Go Playground.
      -	Go Modules (go mod init).


--------------------------------

# Phase 2: Intermediate Skills (4-6 weeks)
1.	Objective: Build practical projects to demonstrate problem-solving skills and familiarity with Go libraries.
2.	Activities:
      - Todo List API
        - Build a RESTful API with CRUD operations for managing a todo list.
        - Tools: net/http, json, mux (or Gin).
      - Concurrency Demo: Worker Pool
        - Implement a worker pool to process tasks concurrently.
        - Example: Download multiple URLs in parallel.
        - Skills: Goroutines, channels, sync primitives.
      - Database Integration
        - Create an app that connects to PostgreSQL/MySQL and performs basic operations.
        - Example: A user management system with signup/login.
        - Skills: database/sql, gorm.
      - Go CLI Tools
        - Build a command-line app like a URL shortener or a markdown-to-HTML converter.
        - Tools: Cobra or native flag package. 
      - Unit Tests and Benchmarks
        - Add test cases and benchmarks to your projects.
        - Skills: testing, go test.
      - URL Shortener Service
        -	Create a URL shortener that generates unique, short links and saves them in memory or a database.
        - Skills: net/http, hash functions, database integration.
      - Weather App with External API
        - Build a CLI or web app that fetches and displays weather data using an API like OpenWeatherMap.
        - Skills: HTTP client, JSON parsing, working with APIs.
      - Image Resizer
        - Create a CLI tool to resize images to a given dimension or percentage.
        - Tools: image package, file handling.
      - Basic Chatbot
        - Write a simple chatbot that responds to user inputs with predefined answers.
        - Example: A bot for FAQs or small talk.
        - Skills: String processing, decision trees.
      - Task Scheduler
        - Build a program that runs tasks at scheduled intervals.
        - Example: Fetch data from an API every minute and log it to a file.
        - Skills: Goroutines, time package.

3.	Tools to Learn:
      - REST APIs with net/http or Gin.
      - Testing (go test), and optionally, mock.

--------------------------------
# Phase 3: Advanced Projects (6-8 weeks)
1.	Objective: Showcase your ability to solve complex problems and work on production-ready systems.
2.	Activities:
      - Real-Time Chat Application
        - Develop a WebSocket-based chat app with real-time updates.
        - Skills: WebSockets, concurrency, server-client architecture.
      - Distributed Key-Value Store
       - Implement a simple distributed system for storing key-value pairs.
       - Skills: Consistent hashing, Goroutines, and networking.
      - Microservices Architecture
        - Build a microservice-based application with gRPC communication.
        - Example: E-commerce app with separate services for users, orders, and payments.
        - Skills: gRPC, protocol buffers, containerization (Docker).
      - Web Scraper
        - Write a program to scrape data from a website and store it in a database.
        - Example: Scrape product prices or news headlines.
        - Tools: colly, http, and database libraries.
      - Implement a Custom Cache
        - Create an in-memory caching solution with LRU or LFU eviction policies.
        - Skills: Data structures, algorithms, benchmarking.
      - Rate Limiter
        - Implement a rate-limiting middleware for a web server, allowing X requests per second.
        - Skills: Middleware, concurrency, and synchronization. 
      - Blockchain Prototype
        - Create a simple blockchain with proof-of-work and basic transactions.
        - Skills: Data structures, cryptography (crypto/sha256).
      - Search Engine
        - Build a lightweight search engine that indexes and retrieves text files.
        - Skills: Indexing algorithms, file I/O, string processing.
      - Kubernetes Controller
        - Write a custom Kubernetes controller in Go to automate tasks (e.g., scaling resources).
        - Skills: Kubernetes API, Go client libraries.
      - Real-Time Stock Tracker
        - Develop an app that tracks stock prices in real time using a WebSocket API.
        - Skills: WebSockets, data visualization, concurrency.

3.	Tools to Learn:
      -	Docker for containerization.
      -	Kubernetes basics for orchestration.
      -	Go profiling tools (pprof) for performance tuning.
-----------------------------

# Bonuses
 - Game Development
   - Create a terminal-based game like Snake or Tic-Tac-Toe. 
   - Skills: CLI design, game loops.
 - Markdown Blog Generator 
   - Build a tool that converts markdown files into a static blog site. 
   - Tools: File I/O, html/template.
 - Log Aggregator 
   - Create a tool to collect, parse, and display logs from multiple services. 
   - Skills: Streams, regular expressions, and log visualization.
 - Go Plugin System 
   - Implement a plugin-based architecture for extending functionality in a Go app. 
   - Example: Add new "commands" to a CLI dynamically. 
   - Tools: plugin package.
 - IoT Data Collector 
   - Write a service to collect and store IoT device data (e.g., temperature sensors). 
   - Skills: MQTT (Message Queuing Telemetry Transport) protocol, database integration.
