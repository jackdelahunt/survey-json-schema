# survey-json-schema

Library for building interactive terminal inputs by using JSON schema.

![surveyjson](./img/surveyjson.gif)

## How this works

Library uses JSON schema and survey (<https://github.com/AlecAivazis/survey>)
to build set of the questions for the users to answer. It enables you to dynamically load various
json schema files and ask terminal users to fill all required data in order to build specific objects.

Library can be used to build interactive mode for tools like Cobra (<https://github.com/spf13/cobra>) CLI tool.

## Usage as CLI

To showcase how library works it provides an CLI that does accept JSON schema as file input.

To run cli with your local json schema execute

```bash
go get github.com/wtrocki/survey-json-schema
askjschema --file=jsonschema.json
```

You can also use remote json schema

```bash
 askjschema --file https://raw.githubusercontent.com/wtrocki/survey-json-schema/main/pkg/surveyjson/test_data/basicTypes.test.schema.json
```

 CLI binary files are also available from Github releases.

## Usage as library

Add library to dependencies

```bash
go get github.com/wtrocki/survey-json-schema
```

Add this sample code to your cobra handler method

```go
    import("github.com/wtrocki/survey-json-schema/pkg/surveyjson")

    // Creates JSONSchema based of
    options := surveyjson.JSONSchemaOptions{
        Out:                 os.Stdout,
        In:                  os.Stdin,
        OutErr:              os.Stderr,
        AskExisting:         true,
        AutoAcceptDefaults:  false,
        NoAsk:               false,
        IgnoreMissingValues: false,
    }

    yourSchema := ""
    initialValues := make(map[string]interface{})
    result, err := options.GenerateValues(yourSchema, initialValues)
    if err != nil {
        return err
    }
    fmt.Fprint(os.Stdin, string(result))
```

For fully functional example see [pkg/cmd/ask.go](pkg/cmd/ask.go) file

## Features

### Supported JSON Schema properties

- String (including specific formating for dates, emails etc.)
- Boolean
- Number
- Null
- Object
- Enum
- AllOf
- If

### Defaults from environment variables

Library can fetch default values for properties from environment variables.
Env var should be constructed as SURVEY_VALUE_{JSON_SCHEMA_PATH}. 
For example:

```bash
SURVEY_VALUE_USER_AGE=3
```
