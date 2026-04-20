c
Everything in Go is organized into packages.
- A **package** is the real unit of Go code organization.
	- Files are just organizational units inside a package, meaning a package can be defined by multiple files (for readability) or just 1.
- `package main` is a special type of package that will make the current Go file act as a executable. Thus, we must define a `main` function.

Go is a statically typed language, meaning every variable has a fixed type known at compile time. It supports type inference, so you often don’t need to explicitly write the type, but the type is still determined at compile time.
- Type inference is done at compile time.

We use `import` to define dependencies that this package uses. Once we define the dependencies that the package will use in 1 `.go` file, we do not need to define it again in another.

To access a function within a package, we use the following syntax:
	`package_name.function_name(args)`

## Testing
To enable for easier testing, we distinguish 2 aspects of a program:
1. **Domain Code**: The pure *logical* and *data*
2. **Side Effects**: Interactions with the *outside world* (i.e. files, network, user input, databases, or anything outside your program’s memory)

Note: Calling a function is not a side effect. Whether a side effect happens depends entirely on **what that function does internally**

This design allows for **Test Driven Development** by separating domain logic from side effects by placing them in different functions.
- Main function contains side effects and we split our domain logic into seperate organized functions (i.e. retrieve data,  perform computation on data like multiply)
[[hello_world.go]]

This allows us to strictly test the domain logic (the code we wrote) and avoiding unnecessary noise, resulting in simpler tests. 
[[hello_test.go]]

Formatters (i.e. `%q, %d, ...`) are called **verbs**.
- `%q` is the formatter for string

`go test` automatically finds test functions, creates a test environment, and runs them, passing in a helper object (`*testing.T`) that lets you report results back.
- You need to create a mod file in order to run tests, which is done through `go mod init [URL]/` 
	- **Modules** are used to manage dependencies in go.
		- Required in order to run `go test` or `go build`
		- Named as URLs to serve as a unique identifier
		- mod files tell you the go version and the exact version dependencies

Rules for writing tests:
1. Needs to be in a file with a name like `xxx_test.go`
2. Test function must start with `Test`
3. Test function only takes in 1 argument (`t *testing.T`)
4. Must import `"testing"` package to use the `testing.T` type

**Best practice**: Write tests first before the actual code logic. If tests fail, we must refactor accordingly based on the compiler errors.

We use **constants** to capture the meaning of values and also to aid in performance.

You can ==group related tests== together using a language feature called **subtests**. The 'T' struct in the testing package has a Run function, which takes:
1. test function name (string) 
2. test function definition, which takes parameter `*testing.T` (just like a normal test would) and is where the actual code for the sub test is written.

`testing.TB` is an interface that `*testing.T` and `*testing.B` (`B` for benchmark) both satisfy.
- Note that variables of type interface accept pointer objects to valid implementations of the interface.

**When you want to add a new behavior to a function, start with a failing test**. If the test is passing, then it means:
1. The behavior is already implemented in the function, OR
2. The test is too weak (or incorrect)

Test Driven Development Cycle:
3. Write a test
4. Run tests and resolve any compiler errors
5. Run the test again to see that it fails and the error message is meaningful
6. Write enough code for the test to pass
7. Refactor (DRY principle)

We rely on tests to check our code rather than manually doing it ourselves. Additionally, we set up our tools (`go test`) so that running tests is simple.

In Go, switch statements are not limited to only integral data types. They work on any comparable data type, including strings!

You can **name the return values** in the function signature.
```go 
func divide(a, b int) (result int) {
    result = a / b
    return
}
```
- Note: `result` is considered as a local variable in the function and it is default initialized.
- Note: We do not have to specify what is being returned. 

**To make a function public**, the function name should start with a capital letter.
**To make a function private**, the function name should start with a lower case letter.

Notes on *initializing* a variable:
- `:=` can only be used locally 
- `var variable_name data_type = value` can be used anywhere
	- `data_type` field is optional. You can opt out and make use of the Go's compiler ==type inference==.
	- Modifiers (such as const) replace the `var` keyword

To *assign* (updating) to an existing variable: `=`

