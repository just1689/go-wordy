package model

const WordTablename = "Word"
const FileName = "words.txt"

//Word describes a word
type Word struct {
	ID string `json:"_id,omitempty" bson:"_id,omitempty"`
}
