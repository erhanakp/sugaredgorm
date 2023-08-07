![Version](https://img.shields.io/badge/version-1.0.0-orange.svg)
![Go](https://img.shields.io/github/go-mod/go-version/erhanakp/sugaredgorm)
[![Documentation](https://godoc.org/github.com/erhanakp/sugaredgorm?status.svg)](https://pkg.go.dev/github.com/erhanakp/sugaredgorm)
[![Go Report Card](https://goreportcard.com/badge/github.com/erhanakp/sugaredgorm)](https://goreportcard.com/report/github.com/erhanakp/sugaredgorm)
![Go Build Status](https://github.com/erhanakp/sugaredgorm/actions/workflows/go-test.yml/badge.svg)
![GolangCI-Lint Status](https://github.com/erhanakp/sugaredgorm/actions/workflows/go-lint.yml/badge.svg)
[![codecov](https://codecov.io/gh/erhanakp/sugaredgorm/branch/main/graph/badge.svg?token=BTVK8VKVZM)](https://codecov.io/gh/erhanakp/sugaredgorm)

# Sugared Gorm

A wrapper for Gorm logger for structured logging. It's using sugared logger of [zap]

## Installation

Now you can add this package via;

```bash
go get github.com/erhanakp/sugaredgorm
```

---

## Usage:
  

```go
sugaredLogger := sugaredgorm.New(sugerlogger)
gormDB, _ := gorm.Open(postgres.New(postgres.Config{}), &gorm.Config{Logger: sugaredLogger})
```

---
## Rake Tasks

```bash
rake -T

rake bump[revision]     # bump version, default is: patch
rake default            # default task
rake doc[port]          # run doc server
rake mockery            # run mockery
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