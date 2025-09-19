package profile

import "time"

// public facing user, the internal user will never get exposed
type Profile struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
}
