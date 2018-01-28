package main

type story struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []options `json:"options"`
}

type options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
