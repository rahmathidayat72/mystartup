package user

import "time"

type User struct {
	Id             int
	Name           string
	Occupation     string
	Email          string
	Password_Hash  string
	AvatarFileName string
	Role           string
	Token          string
	Created_At     time.Time
	Updated_At     time.Time
}
