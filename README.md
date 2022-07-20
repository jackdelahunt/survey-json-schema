# survey-json-schema

Library for building interactive terminal inputs by using JSON schema.

![surveyjson](./img/surveyjson.gif)

## How this works

Library uses JSON schema and survey (<https://github.com/AlecAivazis/survey>)
to build set of the questions for the users to answer. It enables you to dynamically load various
json schema files and ask terminal users to fill all required data in order to build specific objects.

Library can be used to build interactive mode for tools like Cobra (<https://github.com/spf13/cobra>) CLI tool.

## Usage as CLI

```bash
go get github.com/wtrocki/survey-json-schema
askjschema --file=jsonschema.json
```

## Usage as library

```go
    import("github.com/wtrocki/survey-json-schema/pkg/surveyjson")

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
