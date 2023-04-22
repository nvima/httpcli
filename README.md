# httpcli

`httpcli` is a versatile command-line tool for making HTTP requests using YAML configuration files.
You can easily define your API requests and customize the output to extract specific data from the JSON responses.
The tool supports environment variables, allowing you to store sensitive data, such as API keys, securely.
You can also pipe the stdin to your API Response with ${STDIN}.

## Example YAML Configuration

Here's an example of a YAML configuration file that includes two different API requests: `gitdiff` and `translate`.

```yaml PATH: ~/.httpcli.yaml
gitdiff:
  url: "https://api.openai.com/v1/chat/completions"
  header:
    - "Content-Type: application/x-www-form-urlencoded"
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

## Work in Progress
Please note that httpcli is a work in progress tool and may contain errors or incomplete features.
However, suggestions and merge requests are always welcome to help improve the tool.
