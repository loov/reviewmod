// reviewmod.cue - Example configuration
//
// Minimal config - only LLM settings are required.
// All analysis passes use builtin prompts by default.

llm: {
	provider:    "openai"
	base_url:    "http://localhost:1234/v1"
	model:       "qwen/qwen3-next-80b"
	max_tokens:  262144
	temperature: 0.1
}

// Optional: customize which analyses to run
// analyses: [
// 	{name: "summary"},
// 	{name: "security"},
// 	{name: "testing", enabled: false},  // disable specific pass
// 	{name: "custom", prompt: "path/to/custom.txt"},  // custom prompt
// ]
