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

type OperationNotSupported struct {
	Provider  string
	Operation string
}

func (s *OperationNotSupported) Error() string {
	return fmt.Sprintf("This operation (%s) is not supported by %s", s.Operation, s.Provider)
}

type ResourceAlreadyExists struct {
	Provider  string
	Operation string
}

func (s *ResourceAlreadyExists) Error() string {
	return fmt.Sprintf("This operation (%s) is not supported by %s as the resource already exists", s.Operation, s.Provider)
}

type PlayListCreateFailed struct {
	Item  *PlayList
	OgErr error
}

func (s *PlayListCreateFailed) Error() string {
	return fmt.Sprintf("This playlist (%s) could not be created due to  %s", s.Item.Title, s.OgErr.Error())
}
