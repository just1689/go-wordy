package model

//Result describes what happened in response to a request
type Result struct {
	Exists bool  `json:"exists" bson:"exists"`
	Error  error `json:"error" bson:"error"`
}
