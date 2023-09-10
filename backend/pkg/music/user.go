package music

type User struct {
	ID        string
	Name      string
	Followers int
}

type UserInformationProvider interface {
	CurrentUser() (*User, error)
}
