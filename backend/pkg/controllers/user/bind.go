package user

import "github.com/akifkadioglu/vocapedia/pkg/entities"

type _emailData struct {
	Header      string
	Description string
	Warning     string
	Device      string
}
type _updateUser struct {
	Name        string          `json:"name"`
	Username    string          `json:"username"`
	Description string          `json:"description"`
	Biography   string          `json:"biography"`
	Device      entities.Device `json:"device"`
}

type _streak struct {
	Count    int    `json:"count"`
	LastDate string `json:"lastDate"`
	Rewarded bool   `json:"rewarded"`
}
