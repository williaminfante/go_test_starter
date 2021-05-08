# Go Test Starter

This is the accompanying code repo for the [Golang testing with TDD](https://williaminfante.medium.com/golang-testing-with-tdd-e548d8be776) article.


This repository contains the following:

1. [starter package](starter.go) containg three functions
2. [starter_test package](starter_test.go) for the related tests

## Table of Contents

- [Go Test Starter](#go-test-starter)
  - [Table of Contents](#table-of-contents)
  - [Background](#background)
  - [Usage](#usage)
  - [Related Link](#related-link)
  - [Contributing](#contributing)

## Background

This is the accompanying code repo for the [Golang testing with TDD article](https://williaminfante.medium.com/golang-testing-with-tdd-e548d8wbe776).

We’ll introduce three kinds of functions in our starter package. The aim of these functions is to introduce go testing concepts:
- `SayHello()` — basics
- `TestPickAnInteger()` — subtests, refactoring, coverage
- `TestCheckHealth()` — use of more advanced testing libraries


## Usage

This requires a minimum version of Go 1.16 and testify 1.7.0.

You can also just execute:
```
go mod tidy
```

To run tests and optionally add `-v` for verbose:
```
go test -v
```

To add coverage to the desired output like `coverage.out`, simply:
```
go test -v -coverprofile coverage.out
```

To create an html-formatted coverage document:
```
go tool cover -html=coverage.out -o coverage.html
```

## Related Link

Follow-up article on [Golang Testing: Mocking and Error Handling](https://williaminfante.medium.com/golang-testing-mocking-and-error-handling-fbfe7f6008b9)


## Contributing

Please don't hesitate to [raise an issue](https://github.com/williaminfante/go_test_starter/issues/new) or submit a PR to improve this project. Thanks!
