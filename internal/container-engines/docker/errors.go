package docker

type CollectionError int

const (
	ErrorCollectionNotFound CollectionError = iota
	ErrorCollectionAlreadyExists
	ErrorCollectionNotCreated
	ErrorCollectionNotDeleted
	ErrorCollectionNotUpdated
	ErrorCollectionNotStarted
	ErrorCollectionNotStopped
	ErrorCollectionNotRestarted
	ErrorCollectionNotPaused
	ErrorCollectionNotUnPaused
	ErrorCollectionFileNotFound
	ErrorCollectionFileAlreadyExists
	ErrorCollectionFileNotCreated
	ErrorCollectionFileNotDeleted
	ErrorCollectionFileNotUpdated
)

func (e CollectionError) Error() string {
	switch e {
	case ErrorCollectionNotFound:
		return "collection not found"
	case ErrorCollectionAlreadyExists:
		return "collection already exists"
	case ErrorCollectionNotCreated:
		return "collection not created"
	case ErrorCollectionNotDeleted:
		return "collection not deleted"
	case ErrorCollectionNotUpdated:
		return "collection not updated"
	case ErrorCollectionNotStarted:
		return "collection not started"
	case ErrorCollectionNotStopped:
		return "collection not stopped"
	case ErrorCollectionNotRestarted:
		return "collection not restarted"
	case ErrorCollectionNotPaused:
		return "collection not paused"
	case ErrorCollectionNotUnPaused:
		return "collection not un paused"
	case ErrorCollectionFileNotFound:
		return "collection file not found"
	case ErrorCollectionFileAlreadyExists:
		return "collection file already exists"
	case ErrorCollectionFileNotCreated:
		return "collection file not created"
	case ErrorCollectionFileNotDeleted:
		return "collection file not deleted"
	case ErrorCollectionFileNotUpdated:
		return "collection file not updated"
	}
	return ""
}

type ServiceError int

const (
	ErrorServiceNotFound ServiceError = iota
	ErrorServiceAlreadyExists
	ErrorServiceNotCreated
	ErrorServiceNotDeleted
	ErrorServiceNotUpdated
	ErrorServiceNotStarted
	ErrorServiceNotStopped
	ErrorServiceNotRestarted
	ErrorServiceNotPaused
	ErrorServiceNotUnPaused
)

func (e ServiceError) Error() string {
	switch e {
	case ErrorServiceNotFound:
		return "service not found"
	case ErrorServiceAlreadyExists:
		return "service already exists"
	case ErrorServiceNotCreated:
		return "service not created"
	case ErrorServiceNotDeleted:
		return "service not deleted"
	case ErrorServiceNotUpdated:
		return "service not updated"
	case ErrorServiceNotStarted:
		return "service not started"
	case ErrorServiceNotStopped:
		return "service not stopped"
	case ErrorServiceNotRestarted:
		return "service not restarted"
	case ErrorServiceNotPaused:
		return "service not paused"
	case ErrorServiceNotUnPaused:
		return "service not un paused"
	}
	return ""
}

type UserError int

const (
	ErrorUserNotFound UserError = iota
	ErrorUserAlreadyExists
	ErrorUserNotCreated
	ErrorUserNotDeleted
	ErrorUserNotUpdated
)

func (e UserError) Error() string {
	switch e {
	case ErrorUserNotFound:
		return "user not found"
	case ErrorUserAlreadyExists:
		return "user already exists"
	case ErrorUserNotCreated:
		return "user not created"
	case ErrorUserNotDeleted:
		return "user not deleted"
	case ErrorUserNotUpdated:
		return "user not updated"
	}
	return ""
}
