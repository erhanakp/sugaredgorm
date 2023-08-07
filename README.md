![Version](https://img.shields.io/badge/version-0.0.1-orange.svg)
![Go](https://img.shields.io/github/go-mod/go-version/erhanakp/sugaredgorm)
[![Documentation](https://godoc.org/github.com/erhanakp/sugaredgorm?status.svg)](https://pkg.go.dev/github.com/erhanakp/sugaredgorm)
[![Go Report Card](https://goreportcard.com/badge/github.com/erhanakp/sugaredgorm)](https://goreportcard.com/report/github.com/erhanakp/sugaredgorm)
![Go Build Status](https://github.com/erhanakp/sugaredgorm/actions/workflows/go-test.yml/badge.svg)
![GolangCI-Lint Status](https://github.com/erhanakp/sugaredgorm/actions/workflows/go-lint.yml/badge.svg)
[![codecov](https://codecov.io/gh/erhanakp/sugaredgorm/branch/main/graph/badge.svg?token=BTVK8VKVZM)](https://codecov.io/gh/erhanakp/sugaredgorm)

# Sugared Gorm

Custom GORM logger implementation enhanced with SugarLogger for improved log level management and standardized log outputs.

# Features
- Integrate GORM database logging with the powerful SugarLogger library.
- Easily manage log levels and control verbosity for debugging.
- Standardized and customizable log outputs for better readability and analysis.

## Installation

Now you can add this package via;

```bash
go get -u github.com/erhanakp/sugaredgorm
```

Install godoc for documentation;

```bash
go install golang.org/x/tools/cmd/godoc@latest
```

---

## Usage:
  
```go
	zap := zap.NewExample()

	gormLoggerValue := sugaredgorm.New(zap.Sugar(), sugaredgorm.Config{
		SlowThreshold:             200 * time.Millisecond,
		Colorful:                  true,
		IgnoreRecordNotFoundError: true,
		ParameterizedQueries:      true,
	})

	_, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost",
	}), &gorm.Config{
		Logger: gormLoggerValue,
	})
```

---

## Rake Tasks

```bash
rake -T

rake bump[revision]     # bump version, default is: patch
rake doc[port]          # run doc server
rake lint               # run golangci-lint
rake publish[revision]  # publish new version of the library, default is: patch
rake test               # run tests

```

---

## Contributor(s)

* [Erhan Akpınar](https://github.com/erhanakp) - Creator, maintainer
* [Uğur "vigo" Özyılmazel](https://github.com/vigo) - Contributor

---

## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/erhanakp/sugaredgorm/fork)
1. Create your `branch` (`git checkout -b my-feature`)
1. `commit` yours (`git commit -am 'add some functionality'`)
1. `push` your `branch` (`git push origin my-feature`)
1. Than create a new **Pull Request**!

This project is intended to be a safe, welcoming space for collaboration, and
contributors are expected to adhere to the [code of conduct][coc].

---

## License

This project is licensed under MIT

[coc]: https://github.com/erhanakp/sugaredgorm/blob/main/CODE_OF_CONDUCT.md