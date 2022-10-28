package model

type ApplyVote struct {
	ID        int `db:"ID"`
	CitizenID int `db:"CitizenID"`
}
