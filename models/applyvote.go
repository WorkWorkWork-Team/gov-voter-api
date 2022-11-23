package model

type ApplyVote struct {
	ID        int    `db:"ID"`
	CitizenID string `db:"CitizenID"`
}
