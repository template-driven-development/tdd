# TDD

Template Driven Development is a command-line tool that uses Go's [text/template](https://pkg.go.dev/text/template) engine to populate templates from a configuration file.

And yes, we know—TDD is usually about testing, but this time it's all about templates!

## Getting Started

### From Release

Download the latest release from [GitHub](https://github.com/template-driven-development/tdd/releases). Make sure to choose the correct binary for your operating system and architecture. For convenience, you can rename the file and add it to your system's PATH.

Next, create a JSON configuration file to define your inputs, output, and the data used to populate placeholders:

```json
[
    {
        "input": "path/to/input",
        "output": "path/to/output",
        "data": {
            "key": "value"
        }
    }
]
```

Finally, execute the tool with the path to your configuration file:

```sh
./tdd path/to/configuration.json
cat path/to/output
```

### From Source

Clone the repository and run one of the provided example:

```sh
go run . examples/helloworld/configuration.json
cat examples/helloworld/template.out.go
```
