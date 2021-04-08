package domain

import (
	"errors"

	"github.com/fusion44/raspiblitz-backend/db/repositories"
	"github.com/fusion44/raspiblitz-backend/graph/model"
)

// App errors
var (
	ErrBadCredentials           = errors.New("Login credentials not valid")
	ErrUnauthenticated          = errors.New("Unauthenticated")
	ErrUnauthorized             = errors.New("Unauthorized")
	ErrInternalServer           = errors.New("Internal server error")
	ErrInvalidInput             = errors.New("Input not valid")
	ErrUnableToProcess          = errors.New("Unable to process FIT file")
	ErrDuplicateActivityForFile = errors.New("Duplicate activity for FIT file")
)

// Domain contains all business logic
type Domain struct {
	UsersRepo *repositories.UsersRepository
	InfoRepo  *repositories.BlitzInfoRepository
	SetupRepo *repositories.SetupRepository
}

// NewDomain creates a new Domain instance
func NewDomain(
	usersRepo repositories.UsersRepository,
	infoRepo repositories.BlitzInfoRepository,
) *Domain {
	return &Domain{
		UsersRepo: &usersRepo,
		InfoRepo:  &infoRepo,
		SetupRepo: repositories.New()}
}

// Ownable makes an object ownable by an user
type Ownable interface {
	IsOwner(user *model.User) bool
}
