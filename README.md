# GO-GO-GO

Golang's benefits are:

- **Simple Syntax**: Easy to learn and read.
- **Fast**: Compiles quickly and runs efficiently.
- **Concurrency**: Handles multiple tasks well.
- **Modern Use**: Ideal for microservices and cloud apps.
- **Garbage Collection**: Manages memory automatically.
- **Rich Standard Library**: Useful built-in tools and functions.
- **Cross-Platform**: Runs on various operating systems.
- **Built-In Tools**: Includes tools for testing and formatting.
- **Strong Community**: Active support and resources.

## Command

- IDE recommend: **Visual Studio Code (VS Code)**, **GoLand**, **IntelliJ IDEA**
- Install and initial the project: Run `go install ./path/to/your/package` command.
- Clean and add the `go.mod`, `go.sum` files: Run `go mod tidy` command.
- Build and compile the project: Run `go build ./path/to/your/package` command.
- Execute go file: `go run xxxx.go`
- Formate go file: `gofmt -w xxxx.go`

**Typically Project structure**

```text
my-go-project/
├── go.mod // dependencies data
├── go.sum // accurate dependencies record
├── cmd/   // default entry
│   └── myapp/
│       └── main.go
├── pkg/   // public code repo
│   └── mypackage/
│       └── myfile.go
├── internal/ // private code repo
│   └── helper.go
└── README.md
```

## Step by Step

Tour of Go,Go by example, Web Dev, Create a module, CLIs

1. [Download](https://go.dev/dl/) and install the go.
2. Lesson 1: Basic.
3. Lesson 2: Packages, variables and functions.
4. Lesson 3: Flow control statements: for, if, else, switch and defer
5. Lesson 4: More types: structs, slices, and maps.
6. Lesson 5: Methods and interfaces
7. Lesson 6: Generics
8. Lesson 7: Concurrency

### Reference:

- [GO](https://go.dev/)
- [GO Doc](https://go.dev/doc/)
