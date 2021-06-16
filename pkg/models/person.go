package models

type Person struct {
	metadata Metadata
	FirstName, MiddleName, Surname, Nickname string
	Birthday Celebration
	Notes []Note
	Celebrations []Celebration
	Activities []Activity
}
