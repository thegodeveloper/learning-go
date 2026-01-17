
# TUI-based Example Runner Architecture

## Overview

This document outlines the architecture of the TUI-based example runner for the `learning-go` project. The system is composed of two main packages: `internal/registry` and `internal/boxui`. Together, they provide a simple and effective way to navigate and execute the various Go code examples within this repository.

- `internal/registry`: Acts as a central database for all runnable code examples.
- `internal/boxui`: Provides a Terminal User Interface (TUI) to browse and run the examples registered in the `registry`.

## `internal/registry` Package

### Purpose

The `registry` package is the heart of the example-running functionality. It serves as a static registry that maps package names to their corresponding execution functions. This allows the TUI to dynamically discover and run examples without being tightly coupled to each individual example package.

### Configuration

The configuration of the `registry` package is done through two main variables declared in `internal/registry/registry.go`:

1.  **`Packages` (map[string]func(bool))**: This map holds the primary entry point for each example package.
    -   The `key` is the string name of the package (e.g., `"arrays"`).
    -   The `value` is the `Run` function of that package.

2.  **`PackageRegistry` (map[string]PackageFunctions)**: This map is used for packages that have multiple sub-examples.
    -   The `key` is the string name of the package.
    -   The `value` is a `PackageFunctions` map (`map[string]func(bool)`), which maps a sub-function name to its corresponding function.

All registrations are done manually by adding entries to these maps in `internal/registry/registry.go`.

## `internal/boxui` Package

### Purpose

The `boxui` package provides a user-friendly TUI for interacting with the code examples. It is built using the `github.com/rivo/tview` and `github.com/gdamore/tcell/v2` libraries.

### Features

-   Lists all available example packages from the `registry`.
-   Allows navigation into packages that have sub-functions.
-   Executes the selected example or function in a separate goroutine to keep the UI responsive.
-   Captures the standard output of the executed code and displays it in the output panel.
-   Implements a timeout to prevent long-running or interactive examples from blocking the application.
-   Includes panic recovery to ensure the TUI remains stable even if an example crashes.

## Integration

The `boxui` and `registry` packages are designed to work together seamlessly.

1.  **Initialization**: When the `boxui` application starts, it calls `registry.Names()` to get a sorted list of all registered package names.
2.  **Display**: The `boxui` then displays this list to the user.
3.  **Execution**:
    -   When the user selects a package, `boxui` checks if it has sub-functions using `registry.HasSubFunctions()`.
    -   If it has sub-functions, `boxui` calls `registry.GetFunctionNames()` and displays them.
    -   When the user selects a package or a function to run, `boxui` retrieves the corresponding function from the `Packages` or `PackageRegistry` map in the `registry` package.
    -   The retrieved function is then executed, and its output is displayed in the TUI.

## How to Add a New Package

To add a new runnable example to the TUI, you need to register it in the `internal/registry` package. Follow these steps:

1.  **Create Your Package**:
    -   Create a new directory under `internal/` (e.g., `internal/mynewexample`).
    -   Your package must contain a public `Run(show bool)` function that will be the main entry point for your example.
    -   If you have sub-functions, they must also have the `func(bool)` signature.

    *Example (`internal/mynewexample/mynewexample.go`):*
    ```go
    package mynewexample

    import "fmt"

    func Run(show bool) {
        if !show {
            return
        }
        fmt.Println("This is my new example!")
    }
    ```

2.  **Edit `internal/registry/registry.go`**:
    -   Open the `internal/registry/registry.go` file.

3.  **Import Your Package**:
    -   Add your new package to the `import` section.
    ```go
    import (
        // ... other imports
        "github.com/thegodeveloper/learning-go/internal/mynewexample"
    )
    ```

4.  **Register Your Package**:
    -   Add an entry for your package in the `Packages` map.
    ```go
    var Packages = map[string]func(bool){
        // ... other packages
        "mynewexample": mynewexample.Run,
    }
    ```

5.  **(Optional) Register Sub-functions**:
    -   If your package has sub-functions, add an entry to the `PackageRegistry` map.

    *Example:*
    ```go
    var PackageRegistry = map[string]PackageFunctions{
        // ... other packages
        "mynewexample": {
            "subExample1": mynewexample.SubExample1,
            "subExample2": mynewexample.SubExample2,
        },
    }
    ```

Once you have completed these steps, your new package and its functions will automatically appear in the TUI the next time you run the application.
