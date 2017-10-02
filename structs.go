package main

import (
	"time"
)

type Status struct {
	CreatedAt time.Time `json:"timestamp"`
	Sleep     float32   `json:"sleep"`
	Mood      int       `json:"mood"`
}

type MailgunToken struct {
	Timestamp time.Time
	Token     string
	Signature string
}
