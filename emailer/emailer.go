package emailer

import (
	"fmt"
	"sync"
)

// Emailer for email sender
type Provider interface {
	Send(message *Message) error
}

var (
	providersMu sync.RWMutex
	providers   = make(map[string]Provider)
)

// Set the provider
func UseProvider(providerName string) (Provider, error) {
	providersMu.RLock()
	provideri, ok := providers[providerName]
	providersMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("unknown provider %q (forgotten import?)", providerName)
	}
	return provideri, nil
}

// Message defines the message body
type Message struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
}

// Register ...
func Register(providerName string, provider Provider) error {
	providersMu.Lock()
	defer providersMu.Unlock()
	if provider == nil {
		panic("register nil provider")
	}

	if _, dup := providers[providerName]; dup {
		panic("register twice for " + providerName)
	}
	providers[providerName] = provider

	return nil
}
