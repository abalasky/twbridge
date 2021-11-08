package domain

// EventType represents an event type.
type EventType string

const (
	StartEventType       EventType = "start"
	LoginEventType       EventType = "login"
	TextMessageEventType EventType = "text_message" // whatsapp only
	ReplyEventType       EventType = "reply"        // telegram only
)

// Event represents a generic event API.
type Event interface {
	Type() EventType
}

// StartEvent represents an initial event.
type StartEvent struct {
	// ChatID is telegram bot chat identifier.
	ChatID int64

	// FromUser is a telegram username of the client that interacts with the bot.
	FromUser string
}

func (se *StartEvent) Type() EventType {
	return StartEventType
}

// LoginEvent represents a login event.
type LoginEvent struct {
	// ChatID is telegram bot chat identifier.
	ChatID int64

	// FromUser is a telegram username of the client that interacts with the bot.
	FromUser string
}

func (le *LoginEvent) Type() EventType {
	return LoginEventType
}

// TextMessageEvent represents an incoming text message event.
type TextMessageEvent struct {
	// WhatsappRemoteJid is a whatsapp client identifier that sent the message.
	WhatsappRemoteJid string

	// WhatsappSenderName is a whatsapp client's name that sent the message.
	WhatsappSenderName string

	// Text is a text message body.
	Text string
}

func (te *TextMessageEvent) Type() EventType {
	return TextMessageEventType
}

// ReplyEvent represents a message reply event.
type ReplyEvent struct {
	// ChatID is telegram bot chat identifier.
	ChatID int64

	// FromUser is a telegram username of the client that interacts with the bot.
	FromUser string

	// Reply is a reply text message body.
	Reply string

	// TODO: add some identifier to know whom to reply
}

func (re *ReplyEvent) Type() EventType {
	return ReplyEventType
}
