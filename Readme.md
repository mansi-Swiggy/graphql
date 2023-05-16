# README

## Quick start
1. [Initialise a new go module](https://golang.org/doc/tutorial/create-module)

       mkdir example
       cd example
       go mod init example

2. Add `github.com/99designs/gqlgen` to your [project's tools.go](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module)

       printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go

       go mod tidy

3. Initialise gqlgen config and generate models

       go run github.com/99designs/gqlgen init

4. Start the graphql server

       go run server.go

More help to get started:
 - [Getting started tutorial](https://gqlgen.com/getting-started/) - a comprehensive guide to help you get started
 - [Real-world examples](https://github.com/99designs/gqlgen/tree/master/_examples) show how to create GraphQL applications
 - [Reference docs](https://pkg.go.dev/github.com/99designs/gqlgen) for the APIs


###  Some Useful Links & Commands

- Install graphql Library:
  - Command: `go install github.com/99designs/gqlgen@v0.17.25`
- See gqlgen version
  - Command `go run github.com/99designs/gqlgen@v0.17.25 version`
- Generate the schema:
  - Command: `go run github.com/99designs/gqlgen@v0.17.25 generate`. 