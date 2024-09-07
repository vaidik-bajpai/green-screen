package model

type RecordID string
type RecordType string

var (
	RecordTypeMovie = RecordType("movie")
)

type UserID string

type RatingValue int

type Rating struct {
	ID         RecordID    `json:"id"`
	UserID     UserID      `json:"user_id"`
	RecordType RecordType  `json:"record_type"`
	Value      RatingValue `json:"value"`
}
