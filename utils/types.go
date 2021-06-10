package utils

type SendMessageParams struct {
	ChatID   int64
	Text     string
	Keyboard [][]string
}

type EditMessageParams struct {
	SendMessageParams
	MessageID int
	ParseMode string
}
