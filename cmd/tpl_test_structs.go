package cmd

type GptMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGptPayload struct {
	Model    string       `json:"model"`
	Messages []GptMessage `json:"messages"`
}
