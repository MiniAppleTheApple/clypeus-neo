package main

// struct that storage settings
type Settings struct {
	Token           string `json:"token"`      // bot's token
	DataManagerType string `json:"datamanger"` // way that we gonna to storage own data
	Prefix          string `json:"prefix"`     // command prefix
}

var settings *Settings = &Settings{}

func GetSettings() *Settings {
	return settings
}

func SetSettings(v *Settings) {
	settings = v
}
