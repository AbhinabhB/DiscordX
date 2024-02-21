package model

type Guild struct {
	ID      string
	OwnerID string

	Members []User
}

type GuildImpl interface {
	GetGuild() (*Guild, error)
	GetUser() (*User, error)

	CreateGuild() (*Guild, error)
	UpdateGuild() (*Guild, error)
	DeleteGuild() (*Guild, error)

	AddMember() bool
	RemoveMember() bool
}

type GuildHandler interface {
	Create() bool
	Update() bool
	Delete() bool

	FindByID() bool
	FindUserByID() bool
}
