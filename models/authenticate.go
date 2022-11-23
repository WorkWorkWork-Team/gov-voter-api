package model

type AuthenticateBody struct {
	CitizenID string `json:"citizenID"`
	LazerID   string `json:"lazerID"`
}
