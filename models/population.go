package model

import "time"

type Population struct {
	CitizenID   int       `db:"CitizenID"`
	LazerId     string    `db:"LazerID"`
	Name        string    `db:"Name"`
	Lastname    string    `db:"Lastname"`
	Birthday    time.Time `db:"Birthday"`
	Nationality string    `db:"Nationality"`
	DistricID   string    `db:"DistrictID"`
}
