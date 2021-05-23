# go-template

## To init

To use this repository, you should press "Use this template" button. It will clone this repo to another location.

This repository is not ready after cloning, and you should follow a few steps:

1. Replace placeholder to match your new repo:
- `github.com/petuhovskiy/go-template`
- `github.com/<org>/<repo>`

2. Replace placeholder with your name:
- `Your Name`
- `Appropriate Name`

3. Replace the last placeholder with project name:
- `go-template`
- `your-project`

4. Remove this from README and write some code.

You can also check if everything is ok by running existing code:

```bash
# this will download dependencies
go mod download

# this will run existing code
go run main.go

# now program should be running without errors, until Ctrl+C is pressed
```

## To develop

Make sure you have:
- Go 1.16, [install](https://golang.org/doc/install)
- GoLand / VSCode / other IDE, [install goland](https://www.jetbrains.com/go/)
- golangci-lint 1.40, [install](https://golangci-lint.run/usage/install/)


### EnvFile plugin

EnvFile plugin for GoLand is useful for applying conf from .env files. Install [here](https://plugins.jetbrains.com/plugin/7861-envfile).

To use it:
- Open [Run configuration]
- Select EnvFile tab
- Add file .env from repo root
  * On macOS press shirt+cmd+. to display hidden files
