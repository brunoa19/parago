package models

type UserInfo struct {
	Org    string `json:"org,omitempty" bson:"org"`
	User   string `json:"user,omitempty" bson:"user"`
	UserId string `json:"userId,omitempty" bson:"userId"`
	Token  string `json:"token" bson:"token"`
}

func (u *UserInfo) IsGuest() bool {
	return len(u.UserId) == 0
}
