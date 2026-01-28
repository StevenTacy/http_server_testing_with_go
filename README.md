# Testing with Go Practice

This repository contains my practice code and learnings from the "Learn Go with Tests" book. Each directory corresponds to a chapter or concept, focusing on Test-Driven Development (TDD) in Go.

## Learnings & Documentation

Through the development of this project, I have documented several key concepts and patterns in the code:

### Hello World
- **Test Helpers**: `t.Helper()` is used in helper functions (like `assertMsg`) to ensure that when a test fails, the error is reported at the line number of the test call, not inside the helper function itself.

### Integers
- **Testable Documentation**: `Example` functions (e.g., `ExampleAdd`) serve a dual purpose: they act as live documentation for godoc and are compiled/executed as tests to verify the code examples are correct. The `// Output: ...` comment is used to assert the expected standard output.

### Dependency Injection
- **Interface Injection**: Instead of depending on concrete types like `*os.File`, I learned to inject interfaces like `io.Writer`.
- **Flexibility**: This allows functions like `Greet` to write to `os.Stdout` in the main application and `bytes.Buffer` in tests, making the code easily testable.

### Reading Files (Blogposts)
- **File System Abstraction**: used `fs.FS` to decouple code from the real file system. This allows testing with in-memory implementations.
- **Scanning**: used `bufio.Scanner` to read files token by token (or line by line) using `.Scan()` and `.Text()`.

### Arrays & Slices (Sum)
- **DeepEqual**: `reflect.DeepEqual` is useful for comparing complex structures like slices, but it is **not type safe**.
- **Generics**: Implemented generic functions like `Reduce` and `Find` to work with slices of any type.

### Pointers & Errors
- **Pointer Receivers**: Used pointer receivers (e.g., `func (w *Wallet) Deposit`) to modify the state of a struct.
- **Custom Errors**: Defined sentinel errors (e.g., `ErrorMsgInsufficientAmount`) to handle specific logic failures cleanly.

### Mocking
- **Spies**: Implemented "Spy" structs to record calls (e.g., `SpySleeper`, `SpyTime`). This allows verifying *how* dependencies are used without executing their real side effects (like waiting for `time.Sleep`).
- **Configuration**: Built configurable sleepers to switch between real time waiting and mocked waiting.

### Concurrency & Select
- **Goroutines**: Used the `go` keyword to run functions concurrently.
- **Select**: Used `select` to wait on multiple channels. This is useful for:
    - Racing two operations (e.g., `Racer` function).
    - Implementing timeouts with `time.After`.
    - Handling context cancellation.

### Context
- **Cancellation**: Managed request lifecycle using `context.Context`.
- **Propagation**: Passed context down to long-running operations (like `Store.Fetch`) so they can be cancelled early if the client disconnects or times out.

### Sync
- **Mutex**: Used `sync.Mutex` to protect shared state (`Counter`) from race conditions when accessed by multiple goroutines.

### Reflection
- **Walking Structures**: Used the `reflect` package to implement a `Walk` function that can traverse arbitrary data structures (structs, slices, maps, etc.) dynamically.

### Assert Functions
- **Generic Assertions**: Created reusable, generic assertion helpers (`AssertEqual[T comparable]`, `AssertTrue`) to clean up test code and reduce duplication.

## Project Structure

- **arraysSlices**: Tests for array and slice manipulations.
- **assertFunctions**: Generic test helper functions.
- **blogrenderer**: HTML template rendering for blog posts.
- **concurrency**: Examples of goroutines and channel usage.
- **context**: Handling long-running processes and cancellation.
- **dependencyInjection**: Writing testable code by injecting dependencies.
- **helloWorld**: Basic introduction to testing in Go.
- **integers**: Basic arithmetic and example tests.
- **iteration**: Benchmarking and iteration patterns.
- **map**: Dictionary implementation using maps.
- **maths**: SVG clockface generation using math.
- **mocking**: Mocking dependencies (like time) for deterministic tests.
- **perimeter**: Interfaces and table-driven tests for shapes.
- **pointersAndErrors**: Managing state with pointers and handling errors.
- **property_based_test**: Roman numeral conversion using property-based testing ideas.
- **reading_files**: File system operations and parsing.
- **reflection**: Using reflection to traverse data.
- **select**: Coordinating channels with select.
- **sum**: Advanced slice operations (Reduce, Find) and Generics.
- **sync**: Synchronization primitives.
