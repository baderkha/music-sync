package music

import "fmt"

var (
	_ error = &ResourceNotFoundError{}
)

type ResourceNotFoundError struct {
	ResourceType string
	ID           string
	Provider     string
}

func (s *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("Your requested %s with id / name %s could not be found with (%s) provider", s.ResourceType, s.ID, s.Provider)
}

type UnexpectedFatalError struct {
	OgError  error
	Provider string
}

func (s *UnexpectedFatalError) Error() string {
	return fmt.Sprintf("Unexpected fatal error casued by %s from the provider (%s)", s.OgError.Error(), s.Provider)
}
