package models

type Task struct{
		Name string `json:name`
		Priority int `json:priority`
		Due time.Time `json:time`
		Status bool `json:status`
}


