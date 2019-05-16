# URL value parser

[![License](https://img.shields.io/badge/license-BSD--3--Clause-5B74AD.svg)](https://github.com/genert/urlvalueparser/blob/master/LICENSE)

Tool that lets you extract values from URL paths, leaving the other parts untouched.

## Features
By default, the following chunks are consider values that will be replaced:

- Decimal numbers
- Strings in UUID format.
    > Both the standard UUID forms of xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx and urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx are decoded as well as the Microsoft encoding {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx} and the raw hex encoding: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx. 
- Email
- Data URI
- IPV4 and IPV6
- Semver
- Ethereum address
- JWT token


## Usage

### ReplacePathValues(path: String, replacement: String): String

```go
result := urlvalueparser.ReplacePathValues("/legit/path/123/user", ":id")

// Following will be printed: /legit/path/:id/user
fmt.Println(result)
```