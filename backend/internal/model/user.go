package model

/// User structs represents a client of DiscordX
type User struct {
	ID       string
	Username string
	Status   bool

	Guilds   []Guild
	Friends  []User
	Requests []User
}

/// Methods related to account operations
type UserImpl interface {
	Register() bool
	Login() bool

	Get(id string) (*User, error)
	GetByEmail(email string) (*User, error)
}

/// Methods related to internal database operations
type UserHandler interface {
	Create() bool
	Update() bool

	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
}
