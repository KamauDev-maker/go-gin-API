# Beginner's Toolkit for an API using Golang and Gin

## 1. Title and Objective

**Title**:Building a beginner's Toolkit for and API using Golang and Gin

### Objective:
To create a simple toolkit that helps beginners build REST APIs using **Golang** and the **Gin framework**. The toolkit introduces API concepts, provides installation and setup instructions, and includes a working "Hello world" example.

## Prerequisites
Before diving into this API toolkit, you should have basic GO programming knowledge.

## Technology used

Golang(Go) Is fast, simple and widely used in building high-performance APIs.

Gin (web framework) Is a lightweight, high performance web framework for Go that makes routing, middleware and JSON handling easy.

The end goal is to give beginners a ready-to-use toolkit for creating APIs in Go using Gin, with a working example and troubleshooting guide.

## Quick Summary of the Technology
Go is an open-source programming language developed by Google, designed for efficiency, concurrency and simplicity.

Gin is a web framework for Go, similar to Express.js for Node.js or Flask for python. Is widely used to build REST APIs, microservices and web application e.g on Github for open-source API projects.

## System Requirements
*   **Operating System**: Linux / macOS / Windows
*   **Tools /Editors Required**:
    *   Text editor (Vs Code, GoLand, Sublime Text, or Vim)
    *   Go compiler (latest stable Version)
*   **Packages Required**:
    * gin-gonic/gin (Gin framework)

## Installation & Setup Instructions
### step 1: Install Go
*   Download and install from:https://go.dev/doc/install
*   Verify installation:

    ```$ go version ```

*   Step 2 : setup Project Directory

    ```$ mkdir go-gin-api```  
    ```$ cd go-gin-api```     
    ```$ go mod init go-gin-api```  
* Step 3: Install gin   
    ```go get -u github.com/gin-gonic/gin```
## Minimal working Example:
This examples sets up a minimal Gin web server that listens on port 8080 and returns ```"Hello World!"``` when accessed at ```http://localhost:8080/```.

### Expected Output (Terminal):

```go run main.go```    

```[GIN-debug] Listening and serving HTTP on :8080```  

### Expected Output (Browser)

```{"message":"Hello, World! Oscar"}```

## AI Prompt Journal
* *"Building a beginner's toolkit for an API using Go and Gin Framework."*  
* *"Give me a step-by-step guide and help understand the installation of gin".* 
*  *"For the beginner toolkit use Golang and Gin. Provide real-world example, system requirements, and setup instructions".*    
## License
Copyright (c) 2025 [Oscar Kamau](https://github.com/KamauDev-maker/).
