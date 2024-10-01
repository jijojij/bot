package users

type User struct {
	id int64
}

func (u *User) ID() int64 {
	return u.id
}

func NewUser(chatId int64) *User {
	return &User{id: chatId}
}
