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

    ```
    bash
    $ mkdir go-gin-api  
    $ cd go-gin-api     
    $ go mod init go-gin-api
    ```  
    *   `go mod init `creates a `go.mod` file in your project.
    *   The `go.mod` file keeps track of your project name(go-gin-api)
Your `go.mod` file will look like this:

    
    module go-gin-api
    go 1.18

(*your GO version might be different*)
* Step 3: Install gin   
    ```
    bash
    go get -u github.com/gin-gonic/gin
    ```
After the download make sure to import Gin in your code as shown in `main.go` file. 

Run:

```
bash
    go mod tidy
```    
This command adds missing packages to `go.mod`, removes unused ones and updates `go.sum`

`go.mod` should look like:

```
module go-gin-api
go 1.18
require github.com/gin-gonic/gin v1.10.1
```
And you'll also have a go.sum file listing all the exact dependency versions plus checksums.
## Minimal working Example:
This examples sets up a minimal Gin web server that listens on port 8080 and returns ```"Hello World!"``` when accessed at ```http://localhost:8080/```.

### Expected Output (Terminal):

```
bash
go run main.go
```    

```[GIN-debug] Listening and serving HTTP on :8080```  

### Expected Output (Browser)

```
json
{"message":"Hello, World! Oscar"}
```

## AI Prompt Journal
* *"Building a beginner's toolkit for an API using Go and Gin Framework."*  
* *"Give me a step-by-step guide and help understand the installation of gin".* 
*  *"For the beginner toolkit use Golang and Gin. Provide real-world example, system requirements, and setup instructions".*    
## License
Copyright (c) 2025 [Oscar Kamau](https://github.com/KamauDev-maker/).
