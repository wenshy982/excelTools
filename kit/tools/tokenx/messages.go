package tokenx

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (m Message) Str() string {
	return `{"role":"` + m.Role + `","content":"` + m.Content + `"},`
}

func (m Message) Str2() string {
	return `{"role":"` + m.Role + `","content":"` + m.Content + `"}` // 最后一条消息不加逗号
}

type Messages []*Message

func NewMessages() Messages {
	return make(Messages, 0)
}

func (m *Messages) Add(role, content string) {
	*m = append(*m, &Message{Role: role, Content: content})
}

func (m *Messages) AddSystem(content string) {
	m.Add("system", content)
}

func (m *Messages) AddUser(content string) {
	m.Add("user", content)
}

func (m *Messages) AddAssistant(content string) {
	m.Add("assistant", content)
}
