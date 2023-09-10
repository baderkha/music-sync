package music

type Authorizer interface {
	GetHeaders() map[string]string
	RefreshAuth() error
}
