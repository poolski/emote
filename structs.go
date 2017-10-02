package main

import (
	"time"
)

type Status struct {
	CreatedAt time.Time `json:"timestamp"`
	Sleep     float32   `json:"sleep"`
	Mood      int       `json:"mood"`
}

type Config struct {
	MailgunKey string `json:"mailgun_key"`
}
