package interfaces

import "scm.bluebeam.com/stu/golang-template/models"

type (
	// UsersPersister interface for user repositories
	UsersPersister interface {
		GetUser(userID int) (*models.User, error)
		CreateUser(user *models.User) (*models.User, error)
	}
)
