package converters

import (
	"scm.bluebeam.com/stu/golang-template/apis/dtos"
	domainModels "scm.bluebeam.com/stu/golang-template/models"
)

// ToUser converts domain User to api User
func ToUser(serviceUser *domainModels.User) *dtos.User {
	return &dtos.User{ID: serviceUser.ID, Name: serviceUser.Name}
}

// FromUser converts api User to domain User
func FromUser(apiUser *dtos.User) *domainModels.User {
	return &domainModels.User{ID: apiUser.ID, Name: apiUser.Name}
}
