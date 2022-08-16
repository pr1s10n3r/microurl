package url

import (
	"math/rand"
	"time"
)

const CodeLength = 6

type (
	URL struct {
		ID        uint64    `json:"id"`
		Code      string    `json:"code"`
		Value     string    `json:"value"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Repository interface {
		Save(value string) (*URL, error)
		FindByCode(code string) (*URL, error)
	}
)

func NewCode() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	code := make([]rune, CodeLength)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}

	return string(code)
}
