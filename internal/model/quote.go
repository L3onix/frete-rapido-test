package model

type Quote struct {
	Recipient Recipient `json:"recipient"`
	Volumes   []Volume  `json:"volumes"`
}
