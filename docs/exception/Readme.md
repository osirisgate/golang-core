# 🌐 Osirisgate Core Exception Library for Go

> The Osirisgate core exception library for Go provides robust and extensible error handling utilities based on
typed HTTP status codes. Designed for seamless integration into API responses and logging within modern Go applications.

-----

## 📦 Installation

To integrate this library into your Go project, use `go get`.

1.  **Initialize a Go module** in your project (if you haven't already):

    ```bash
    go mod init your_module_path/your_project_name
    ```

    *(Replace `your_module_path/your_project_name` with your actual module path, e.g., `github.com/myuser/myapp`)*

2.  **Download the library**:

    ```bash
    go get github.com/osirisgate/golang-core@latest
    ```

## ✅ Requirements

* **Go 1.18 or higher**: This library leverages Go 1.18+ features.

## ✨ Features

* Typed custom errors with HTTP status codes.
* `StatusCode` and `Status` (success/error string) enums for clarity.
* Consistent interface (`CoreInterface`) for unified error handling with more context.
* Implementation via composition (struct embedding) for code reuse.
* Utility methods for API response formatting and logging.

## 🧱 Project Structure in Go

The library is structured into Go packages.

```
module/ root
    ├── enum/       # Contains enum definitions and HTTP status codes
    │   └── status.go
    │   └── status_code.go
    └── exception/  # Contains interfaces and some exceptions implementations
        ├── exception.go
        └── bad_function_call.go # Example of a concrete exception
    ├── docs/     # Documentation files
        └── exception/ # Contains the main documentation for the exception library
            ├── Readme.md
    └── tests/     # Contains unit tests for the library
        ├── enum/
            ├── status_code_test.go
            ├── status_test.go
        └── exception/
            ├── exception_test.go
    go.mod // go.sum # Go module files
```

## 🧠 Key Concepts in Golang

In Go, OOP concepts like class inheritance, explicit interface implementation keywords, or traits do not exist directly. They are replaced by more idiomatic mechanisms:

* **Implicit Interfaces**: A type implements an interface if it provides all the methods declared in that interface, without an explicit `implements` keyword.
* **Composition (`Struct Embedding`)**: Instead of inheriting, a `struct` can embed another `struct`. The "parent" `struct` then gains access to the fields and methods of the embedded `struct`. This is how code reuse is achieved for `CoreException`.
* **Error Handling by Return Values**: Go prefers returning an `error` as the last return value of a function, rather than throwing exceptions with `try/catch`. The built-in `error` type is itself an interface (`type error interface { Error() string }`).
* **Constructors (`New...` functions)**: `structs` do not have automatic constructors. We use `New` prefixed functions (e.g., `NewInstance`) to create and initialize instances.

-----

## 🧑‍💻 Basic Usage

### 1\. `github.com/osirisgate/golang-core/enum` Package

This package provides clear types for statuses and HTTP codes.

```go
package main

import (
	"fmt"
	status "github.com/osirisgate/golang-core/enum" // Import the enum package
)

func main() {
	// Using Status constants (success/error)
	operationStatus := status.SUCCESS
	fmt.Printf("Operation Status: %s\n", operationStatus.GetValue()) // Output: "Operation Status: success"

	// Using StatusCode constants (HTTP codes)
	httpCode := status.NotFound
	fmt.Printf("HTTP Code: %d - Description: %s\n", httpCode.GetValue(), httpCode.GetDescription()) // Output: "HTTP Code: 404 - Description: Not Found"

	// Creating a StatusCode from an integer value
	code, ok := status.NewStatusCode(200)
	if ok {
		fmt.Printf("Created Code: %d - %s\n", code.GetValue(), code.GetDescription()) // Output: "Created Code: 200 - OK"
	}
}
```

### 2\. `github.com/osirisgate/golang-core/exception` Package

This package contains the core error handling logic.

#### **`CoreInterface`: The Contract for Rich Errors**

```go
package my_exception

// CoreInterface defines the methods to access structured data from an error.
type CoreInterface interface {
	error // Embeds the standard Go error interface (like \Throwable in PHP)
	Format() map[string]interface{}
	GetErrors() map[string]interface{}
	GetDetails() map[string]interface{}
	GetDetailsMessage() string
	GetStatusCode() int
	GetErrorsForLog() map[string]interface{}
}
```

#### **Create a Custom Exception**

A custom exception is a `struct` that embeds `CoreException`.

```go
// File: your_module_path/core/exception/resource_not_found.go
package exception

import (
	status "github.com/osirisgate/golang-core/enum"
	"github.com/osirisgate/golang-core/exception"
)

// ResourceNotFound represents a 404 Not Found exception.
type ResourceNotFound struct {
	CoreException // Embedding the base struct.
}

// NewResourceNotFound is the constructor for ResourceNotFound.
func NewResourceNotFound(errors map[string]interface{}) *ResourceNotFound {
	base := NewInstance(errors, status.NotFound) // Initialize with 404 status
	return &ResourceNotFound{CoreException: *base}
}

// File: your_module_path/exception/bad_function_call.go
package exception

import "github.com/osirisgate/golang-core/enum"

// BadFunctionCall represents an error due to an incorrect function call (400).
type BadFunctionCall struct {
	CoreException
}

// NewBadFunctionCall is the constructor for BadFunctionCall.
func NewBadFunctionCall(errors map[string]interface{}) *BadFunctionCall {
	base := NewInstance(errors, enum.BadRequest) // Initialize with 400 status
	return &BadFunctionCall{CoreException: *base}
}
```

#### **Trigger (Return) an Exception in Your Code**

In Go, functions return errors as their last return value.

```go
package main

import (
	"fmt"
	"github.com/osirisgate/golang-core/exception"
)

// GetUser simulates fetching a resource that might not exist.
// It returns an error that implements the CoreInterface.
func GetUser(id int) (string, error) {
	if id == 123 {
		return "Alice", nil // Nil for error means success
	}
	// Return an instance of your custom exception.
	return "", exception.NewResourceNotFound(map[string]interface{}{
		"message": "The requested user was not found.",
		"details": map[string]interface{}{
			"resource": "user",
			"id":       id,
		},
	})
}

// PerformAction simulates an action that might fail due to incorrect arguments.
func PerformAction(args []string) error {
	if len(args) < 2 {
		return exception.NewBadFunctionCall(map[string]interface{}{
			"message": "Missing required arguments.",
			"details": map[string]interface{}{
				"expected_args": 2,
				"received_args": len(args),
			},
		})
	}
	return nil
}
```

#### **Retrieve Exception Data**

Use an error check (`if err != nil`) and a **type assertion** to access the specific methods of your `CoreInterface`.

```go
func main() {
	// --- Example using ResourceNotFound ---
	_, errUser := GetUser(456)
	if errUser != nil {
		if customErr, ok := errUser.(exception.CoreInterface); ok {
			fmt.Println("\n--- Structured Exception Data (ResourceNotFound) ---")
			fmt.Printf("Message: %s\n", customErr.Error())             // Corresponds to getMessage()
			fmt.Printf("HTTP Code: %d\n", customErr.GetErrorsForLog()["status_code"]) // Corresponds to getCode()
			
			fmt.Printf("Errors: %+v\n", customErr.GetErrors()) // Corresponds to getErrors()
			fmt.Printf("Details: %+v\n", customErr.GetDetails())          // Corresponds to getDetails()
			
			fmt.Printf("Detail Message: %s\n", customErr.GetDetailsMessage()) // Corresponds to getDetailsMessage()

			fmt.Println("\n--- Formatted API Response (ResourceNotFound) ---")
			fmt.Printf("%+v\n", customErr.Format()) // Corresponds to format()
		}
	}

	// --- Example using BadFunctionCall ---
	errAction := PerformAction([]string{"arg1"})
	if errAction != nil {
		if badCallErr, ok := errAction.(exception.CoreInterface); ok {
			fmt.Println("\n--- Structured Exception Data (BadFunctionCall) ---")
			fmt.Printf("Message: %s\n", badCallErr.Error())
			fmt.Printf("HTTP Code: %d\n", badCallErr.GetErrorsForLog()["status_code"])
			fmt.Printf("Details: %+v\n", badCallErr.GetDetails())
			// ... you can call any other CoreInterface methods
		}
	}
}
```

#### **Example Outputs**

These outputs illustrate what you will get by using the methods on an instance of your `CoreException` or a custom exception (like `ResourceNotFound`).

```text
// For a ResourceNotFound instance created as follows:
// exception.NewResourceNotFound(map[string]interface{}{
//     "message": "The requested user was not found.",
//     "details": map[string]interface{}{"user_id": 456},
// })

// customErr.Error()
// Output: "The requested user was not found."

// customErr.GetErrorsForLog()["status_code"] (or customErr.StatusCode.GetValue())
// Output: 404

// customErr.GetErrors()
/*
map[string]interface{}{"details":map[string]interface{}{"user_id":456}}
*/

// customErr.GetDetails()
/*
map[string]interface{}{"user_id":456}
*/

// customErr.GetDetailsMessage()
// Output: "" (since there is no direct "error" key in the details for this example)

// customErr.GetStatusCode()
// Output: 400

// customErr.GetErrorsForLog()
/*
map[string]interface{}{"errors":map[string]interface{}{"details":map[string]interface{}{"user_id":456}}, "message":"The requested user was not found.", "status_code":404}
*/

// customErr.Format()
/*
map[string]interface{}{"details":map[string]interface{}{"user_id":456}, "enum":"error", "error_code":404, "message":"The requested user was not found."}
*/
```

-----

## 🧪 Run Tests
Execute the following command to run all tests across all packages:

```bash
go test ./...
```

For more verbose output:

```bash
go test -v ./...
```

## 🧰 Developer Tools

Use standard Go tools to maintain code quality:

* **Code Formatting**:
  ```bash
  go fmt ./...
  ```
* **Static Analysis (Linter)**:
  ```bash
  go vet ./...
  ```

## 📜 License

This package is licensed under the [MIT License](https://www.google.com/search?q=LICENSE).

## 👤 Author

**Ulrich Geraud AHOGLA** | Software Engineer

For any questions or feedback, please contact me at [developer@osirisgate.com](mailto:developer@osirisgate.com).