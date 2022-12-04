package main

type message struct {
	Status  string `json:"status"`
	Channel string `json:"channel"`
	ID      int    `json:"ID"`
	Body    string `json:"body"`
	Queue   string `json:"queue"`
}
