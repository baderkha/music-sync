package music

// StreamingServiceProvider : an interface that has cloud methods for fetching user related playlist information
type StreamingServiceProvider interface {
	PlayListManager
	SongManager
	UserInformationProvider
	WithAuthorizer(a Authorizer) StreamingServiceProvider
}
