# Template Driven Development

Template Driven Development is a command-line tool that uses Go's [text/template](https://pkg.go.dev/text/template) engine to populate templates from a configuration file.

And yes, we know—TDD is usually about testing, but this time it's all about templates!

## Example

```sh
go run . examples/helloworld/args.json
cat examples/helloworld/main.go.out
```
