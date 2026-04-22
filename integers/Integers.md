Go source files can only have 1 package per directory. Thus, if you are building a program with several packages, each for a specific purpose, then organize them into seperate directories.

The verb (formatter) for an int is `%d`.

In Go, the visibility of a function (private or public) means package-level visibility not file-level.

Named return values should only be used when what is being returned is not clear.
- [[Hello, World]]

Adding comments at the top of functions will appear in the function's Go Docs.
- To look at Go Docs, run the following command in the terminal: `pkgsite --open`
	- This page contains documentation for standard library packages, packages you have installed, and documentation of your own code locally (from the directory in which you ran it from).

**Testable Examples** in Go demonstrate how to use a function while optionally verifying its behavior.
- Testable Examples go in `*_test.go` files.
- They appear in a function’s Go Docs (`pkgsite --open`).
- If the example becomes invalid after code changes, the test run will fail.

- To verify output, include a special comment:
  `// Output: [Result]`
	- Go executes the example and compares printed output to this comment.
	- Functions as an assertion.
	- If it matches → passes silently.
	- If it doesn’t → fails during `go test`.

- If `// Output:` is omitted:
	- The example is not checked.
	- It serves as documentation / demonstration of how to use the function only.

- Syntax:
	- No `t *testing.T` parameter is needed.
	- Use `fmt.Println` to produce output.

Note: Use `go test -v` to see which tests and examples are being run.