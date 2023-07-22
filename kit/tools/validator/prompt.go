package validator

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

var PromptValidator = newValidator()

//go:embed banned.json
var bannedData []byte

type Validator interface {
	FilterPrompt(prompt string) error
}

func newValidator() Validator {
	// Parse bannedData into a slice of strings
	var list []string
	if err := json.Unmarshal(bannedData, &list); err != nil {
		// This should never happen
		panic(err)
	}
	lookup := make(map[string]struct{})
	for _, word := range list {
		lookup[word] = struct{}{}
	}
	return &validator{banned: lookup}
}

type validator struct {
	banned map[string]struct{}
}

// FilterPrompt Check prompt content if any words are banned.
func (v *validator) FilterPrompt(prompt string) error {
	// Check if prompt is empty
	if prompt == "" {
		return nil
	}

	// Convert prompt to lowercase
	prompt = strings.ToLower(prompt)

	// Split prompt into words using any whitespace or punctuation as a delimiter
	words := strings.FieldsFunc(prompt, func(r rune) bool {
		return r == ' ' || r == ',' || r == '.' || r == '!' || r == '?' ||
			r == ';' || r == ':' || r == '-' || r == '_' || r == '(' ||
			r == ')' || r == '[' || r == ']' || r == '{' || r == '}' ||
			r == '"' || r == '\'' || r == '/' || r == '\\' || r == '|' ||
			r == '@' || r == '#' || r == '$' || r == '%' || r == '^' ||
			r == '&' || r == '*' || r == '+' || r == '=' || r == '<' ||
			r == '>' || r == '~' || r == '`' || r == '\t' || r == '\n'
	})

	// 单个词的校验
	for _, word := range words {
		if _, ok := v.banned[word]; ok {
			return fmt.Errorf("word [%q] is banned", word)
		}
	}
	// 多词的校验
	for key, _ := range v.banned {
		if !strings.Contains(key, " ") {
			continue
		}
		if containsPattern(prompt, fmt.Sprintf(`\b%s\b`, key)) {
			return fmt.Errorf("[%q] is banned", key)
		}
	}
	return nil
}

func containsPattern(str, pattern string) bool {
	return regexp.MustCompile(pattern).MatchString(str)
}
