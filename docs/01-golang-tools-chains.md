
# Tool chains

## Commands

<details><summary>Go is a tool for managing Go source code.</summary>

```go
Usage:

	go <command> [arguments]

The commands are:

	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         update packages to use new APIs
	fmt         gofmt (reformat) package sources
	generate    generate Go files by processing source
	get         add dependencies to current module and install them
	install     compile and install packages and dependencies
	list        list packages or modules
	mod         module maintenance
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         report likely mistakes in packages
```

Frequently used commands:
- **build**
  - cross-OS build: `env GOOS=linux GOARCH=amd64 go build`
  - amd64 means the kernel of the OS.

- **install**
  - like build,also means compile the code, but it will build a package into the folder pkg.

- **get**
  - get 3rd party packages,normally will pull the latest packages from the github repository.
  - e.g. `go get -u github.com/go-sql-driver/msyql` `-u` means use the latest packages.

- fmt
  - reformat package sources.

- test
  - normally will test all the test in the current package.
  - `go test -v` means print detail infos of the test.
  - test files usually named like `xxx_test.go`(not necessary)
</details>



## go test

- `go test -v`: print much more info.
- A test file name must end with `_test.go`.
- Each test file must import testing.
- The function name in the test files must starts with `TestXxx`, otherwise the go will skip the test case and don't test the function.
- `t.Errorf` will print the error info, and will skip the test case.
- `t.SkipNow()` will skip the current test case, and will PASS the test case and continue to test the next test case.

Order:
- Go test will not test all test case in exact order,mostly in the Sequential execution.
- If you have to run multiple test cases in certain order,you can use the **subtests**,like this:

	```go
	func TetsPrint(t *testing.T){
		t.run("a1",func(t *testing.T){fmt.Println("a1")})
		t.run("a2",func(t *testing.T){fmt.Println("a3")})
		t.run("a3",func(t *testing.T){fmt.Println("a2")})
	}
	```

If you want to do some prepare work before all test, you can use the `TestMain`:
```go
func TestMain(t *testing.M){
	// some code
	m.run()
}
```
TestMain is for setup works before the tests,ues `m.Run()` to call other tests to complete some set-up testing, like database connection,open a file, login into a RESTful service.

If you didn't call the `m.Run()` in the TestMain test,other tests will not be executed.

#### benchmark in go test
Some tips to know:
- The benchmark function should start with `Benchmark`.
- The benchmark test case will run `b.N` time in every execution.
- During the execution process, the number of `b.N` will be increased to a stable state according to whether the actual case execution time reaches stability.

Example:

    ```go
    func Benchmark(b *testing.B) {
        for i := 0; i < b.N; i++ {
            fmt.Println(i)
        }
    }
    ```

Use command `go test -bench=.` to execute the benchmark.

- [ ] Benchmark test usage.

**You must make sure the function in the benchmark will reach stability, or the benchmark maybe never stop**.

Like this:
```go
func a(n int) int{
	for n > 0{
		n--
	}
	return n
}

// this benchmark will never stop !!!
func BenchmarkAll(b *testing.B){
	for n:= 0; n< b.N; n++{
		a(n)
	}
}
```
