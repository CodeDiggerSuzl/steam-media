# Stream-media
$$Author: Suz1$$

Building a stream video website using golang, like Youku or netflix this kind of stream-media type websites. Unlike the real-time streaming website, like Douyu(which mainly focus on the real-time encode and decode), stream media is just using a stream to play the video.

## Getting start
### Again why golang?

- High development efficiency(Full packed develop engineer chain tools: test, benchmark, multiple build-in.etc )
- Easily deployed(Compile once, run everywhere) native code, compiled language.
- Superb native http lib and awesome template engine(Don't need other frameworks)
- Excellent concurrency programming model.
- and .etc.
### Steps

1. Basic golang usage introduction, by using a simple webserver to understand the go tool-chain and other usages, testing, debugging .etc.
2. A stream media website.
3. Deploy to the cloud.

### Key points
Go usages:
- Go tool-chain usages: fmt, build, install..etc.
- Go test and Go benchmark.

Building the website:
- Front-end and back-end system architecture design.
- Design and implementation of RESTful API.
- Use golang to achieve web service.
- Decoupling of system services(SOA decoupling).
- Practice and application of concurrency-model and channel.
- Using golang original template to achieve the web UI.

To the cloud:
- Refactor the business engineering framework by using the could service, such as ECS and ELB.
- Implement business deployment under cloud-native architecture on Alibaba Cloud(cloud-native: maximize use of infrastructure and basic services provided by the cloud).


---
## [Golang tool chains](./docs/01-golang-tools-chains.md)
- Useful commands.
- Testing.
- Benchmark

## Getting started

### Architecture of the stream-media website

# [Bug fix log](./docs/bug-fix-log.md)

A bug fix log of me during coding and learning thought this project.

This might help you to avoid some bugs.







---
refs:
- [ ] [a another golang server](https://blog.csdn.net/qq_44291044/article/details/99703150)
- [ ] [note](https://alanhou.org/golang-video-streaming/)
