package entity

import (
	"fmt"
	"unicode"
)

type Key struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

func (k *Key) Validate() error {
	if len(k.Key) < 1 || len(k.Key) > 40 {
		return fmt.Errorf("key must be 1-40 symbols")
	}

	for _, r := range k.Key {
		if !unicode.Is(unicode.Latin, r) && !unicode.Is(unicode.Number, r) {
			return fmt.Errorf("key must contain only latin symbols or numbers ")
		}
	}

	if len(k.Text) < 1 || len(k.Text) > 40 {
		return fmt.Errorf("text must be 1-40 symbols")
	}

	for _, r := range k.Text {
		if !unicode.Is(unicode.Latin, r) && !unicode.Is(unicode.Number, r) {
			return fmt.Errorf("text must contain only latin symbols or numbers ")
		}
	}
	return nil

}
