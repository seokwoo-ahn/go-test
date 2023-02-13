package types

type Configs struct {
	Port     int    `json:"port"`
	Url      string `json:"url"`
	Database Database
}

type Database struct {
	Name       string `json:"name"`
	Collection string `json:"collection"`
}
