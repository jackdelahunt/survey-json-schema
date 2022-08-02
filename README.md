# survey-json-schema

Library for building interactive terminal inputs by using JSON schema.

![surveyjson](./img/surveyjson.gif)

## How this works

Library can be used to build interactive mode for golang terminal applications like Cobra (<https://github.com/spf13/cobra>) etc.
Library uses JSON schema as specification for questions and survey go library to execute them (<https://github.com/AlecAivazis/survey>).
Developers will get JSON files containing answers from the user that conform to specified JSON schema. 

In typical execution library will:

- Read and process input JSON Schema file
- Traverse JSON schema to create number of questions expecting specific types as answer
- Validate user input against schema spec
- Merge supplied defaults with the answers provided by users

## Usage as CLI

Project provides an CLI that reads JSON schema as file input.
To run cli with your local json schema execute

```bash
go get github.com/jackdelahunt/survey-json-schema
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
go get github.com/jackdelahunt/survey-json-schema
```

Add this sample code to your cobra handler method

```go
    import("github.com/jackdelahunt/survey-json-schema/pkg/surveyjson")

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

## Source

Library is basing on JSONSchema processor code originally contributed by Pete Muir for JenkinsX project.
