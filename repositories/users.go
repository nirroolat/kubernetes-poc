package repositories

import (
	"database/sql"

	"scm.bluebeam.com/stu/golang-template/models"
)

const sqlNotFound = "sql: no rows in result set"

type (
	// UsersRepository represents a repository for user information
	UsersRepository struct {
		DB *sql.DB
	}
)

// GetUser by ID
func (repository *UsersRepository) GetUser(userID int) (*models.User, error) {
	// TODO: Uncomment once a database is setup
	/*user := models.User{}
	stmt, err := repository.DB.Prepare("SELECT id, name FROM user WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(userID).Scan(&user.ID, &user.Name)
	if err != nil {
		if err.Error() == sqlNotFound {
			return nil, errors.NotFound{Message: fmt.Sprintf("User with ID %d not found", userID)}
		}
		return nil, err
	}
	return &user, nil*/
	return &models.User{Name: "hardcoded-name", ID: 1}, nil
}

// CreateUser in repository and return repository
func (repository *UsersRepository) CreateUser(user *models.User) (*models.User, error) {
	// TODO: Uncomment once a database is setup
	/*resultUser := models.User{}
	stmtInsert, err := repository.DB.Prepare("INSERT INTO user (name) values(?)")
	if err != nil {
		return nil, err
	}
	defer stmtInsert.Close()
	result, err := stmtInsert.Exec(user.Name)
	if err != nil {
		return nil, err
	}
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	stmtSelect, err := repository.DB.Prepare("SELECT id, name FROM user WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmtSelect.Close()
	err = stmtSelect.QueryRow(lastInsertedID).Scan(&resultUser.ID, &resultUser.Name)
	if err != nil {
		return nil, err
	}
	return &resultUser, nil*/
	return &models.User{Name: "hardcoded-name", ID: 1}, nil
}
