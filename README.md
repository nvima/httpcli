# httpcli

`httpcli` is a versatile command-line tool for making HTTP requests using YAML configuration files.
You can easily define your API requests and customize the output to extract specific data from the JSON responses.
The tool supports environment variables, allowing you to store sensitive data, such as API keys, securely.
You can also pipe the stdin to your API Response with ${STDIN}.

- [Example YAML Configuration](#example-yaml-configuration)
- [Usage Examples](#usage-examples)
  - [Translate Text](#translate-text)
  - [Generate Git Commit Message](#generate-git-commit-message)
- [YAML Attributes](#yaml-attributes)
- [Passing Arguments to httpcli Commands](#passing-arguments-to-httpcli-commands)
- [CLI Flags](#cli-flags)

## Example YAML Configuration
Here's an example of a YAML configuration file that includes three different API requests: `gitdiff`, `translate` and `ask`.
~/.httpcli.yaml
```
gitdiff:
  url: "https://api.openai.com/v1/chat/completions"
  header:
    - "Content-Type: application/json"
    - "Authorization: Bearer ${OPENAI_API_KEY}"
  data:
    model: "gpt-3.5-turbo"
    messages:
      - role: "user"
        content: |
          I am working on a project and have made some changes to the code.
          I used the "git diff" command to see the differences.
          Based on the following git diff output, 
          please help me write a descriptive and informative git commit message:
          ${STDIN}
  output: "choices[0].message.content"
  statuscode: 200
  env:
    - "OPENAI_API_KEY"
translate:
  url: "https://api-free.deepl.com/v2/translate"
  header:
    - "Content-Type: application/json"
    - "Authorization: DeepL-Auth-Key ${DEEPL_AUTH}"
  data:
    text:
      - "${STDIN}"
    target_lang: "FR"
  output: "translations[0].text"
  env:
    - "DEEPL_AUTH"
ask:
  url: "https://api.openai.com/v1/chat/completions"
  header:
    - "Content-Type: application/json"
    - "Authorization: Bearer ${OPENAI_API_KEY}"
  data:
    model: "gpt-3.5-turbo"
    messages:
      - role: "user"
        content: ${ARG1}
  output: "choices[0].message.content" 
  env:
    - "OPENAI_API_KEY"
```

## Usage Examples
Using the above YAML configuration, you can easily make API requests and customize the output.

### Translate Text
To translate the text "Apple" to French using the translate API request:
```
$ echo "Apple" | httpcli translate
Pomme
```

### Generate Git Commit Message
To generate a git commit message based on a git diff output using the gitdiff API request:
```
$ git diff --staged | httpcli gitdiff
Added initial implementation of main.go file with necessary package imports and function call to execute CLI command.
```

### Ask a Question
To ask a random Question with an argument
```
$ httpcli ask "How many times Germany became world champion"
Germany has become world champion four times: in 1954, 1974, 1990, and 2014.
```

## YAML Attributes
In your YAML configuration file, the following attributes can be used to define your API requests:

| Attribute  | Required/Optional | Description                                                                                                         |
|------------|------------------|---------------------------------------------------------------------------------------------------------------------|
| `url`      | Required         | The URL to which the API request will be sent.                                                                      |
| `header`   | Optional         | Headers to be included in the API request.                                                                          |
| `data`     | Optional         | The payload to be sent with the API request                                                                         |
| `statuscode` | Optional       | The expected status code for the API response. If a different response code is received, the CLI tool will return an error and not proceed with parsing the response. |
| `output`   | Optional         | Specifies how the JSON response should be parsed, e.g., "translations[0].text". If omitted, the entire JSON response will be written to stdout without parsing. |
| `env`      | Optional         | Searches for environment variables in other YAML attributes and replaces them with the corresponding values using this pattern: ${API_KEY}. |


This structure allows you to easily define and customize your API requests within the YAML configuration file.

## Passing Arguments to httpcli Commands

With `httpcli`, you can pass dynamic values to your commands using the `${ARG1}`, `${ARG2}`, and so on, placeholders in your YAML configuration file. These placeholders will be replaced with the respective command line arguments provided when invoking `httpcli`.

For example, let's say you have a translation API configured in your YAML file, and you want to translate a given word or phrase in a target_lang. You can use the `${ARG1}` and `${ARG2}` placeholder in your YAML configuration to represent the text you want to translate:

```yaml
translate:
  url: "https://api.example.com/translate?text=${ARG1}&target_lang={ARG2}"
  output: "translations[0].text"
```
Now, when you call httpcli with the translate command and provide a text argument, the ${ARG1} placeholder will be replaced with the provided text:
```
httpcli translate "A green Apple" fr
```

## CLI Flags

`httpcli` supports various flags to customize its behavior and provide additional functionality. These flags can be passed along with the command to modify the tool's behavior:

- `--config string`: Specify a custom configuration file. By default, the tool uses the `$HOME/.httpcli.yaml` file. You can provide a different file path using this flag (e.g., `httpcli --config=./my-config.yaml`).
- `--debug`: Enable debug mode to display more detailed error messages and logs for troubleshooting purposes.
- `-h, --help`: Display help information for the `httpcli` command, including a list of available flags and usage instructions.

Use these flags in combination with your API request commands to customize the behavior of `httpcli` according to your needs.

## Work in Progress
Please note that httpcli is a work in progress tool and may contain errors or incomplete features.
However, suggestions and merge requests are always welcome to help improve the tool.
