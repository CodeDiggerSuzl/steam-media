# Stream-media

Building a stream video website using golang,like Youku or netflix this kind of stream-media type web sites. Unlike real-time streaming website,like Douyu(which mainly focus on the real-time encode and decode), stream media is just using a stream to play the video.

## Getting start
### Again why golang?

- High development efficiency(Full packed develop engineer chain tools: test,benchmark,multiple build-in.etc )
- Easily deployed(Compile once,run everywhere) native code,compiled language.
- Superb native http lib and awesome template engine(Don't need other frameworks)
- Excellent concurrency programming model.
- and .etc
### Steps

1. Basic golang usage introduction,by using a simple webserver to understand the go tool-chain and other usages,testing,debugging .etc.
2. A stream media website.
3. Deploy to the cloud.

### Key points
Go usages:
- Go tool-chain usages: fmt,build,install..etc.
- Go test and Go benchmark.

Building the website:
- Front-end and back-end system architecture design.
- Design and implementation of RESTful API.
- Use golang to achieve web service.
- Decoupling of system services(SOA decoupling).
- Practice and application of concurrency-model and channel.
- Using golang original template to achieve the web UI.

To the cloud:
- Refactor the business engineering framework by using the could service,such as ECS and ELB.
- Implement business deployment under cloud-native architecture on Alibaba Cloud(cloud-native: maximize use of infrastructure and basic services provided by the cloud)

## Basic use of golang

### A demo to go
#### Tool chains
##### Commands
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
  - get 3rd party packages,normally will pull the latest packages from github repository.
  - e.g. `go get -u github.com/go-sql-driver/msyql` `-u` means use the latest packages.

- fmt
  - reformat package sources.

- test
  - normally will test all the test in the current package.
  - `go test -v` means print detail infos of the test.
  - test files usually named like `xxx_test.go`(not necessary)
</details>


 
