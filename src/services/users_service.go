package services

import (
	"gitlab.com/chertokdmitry/surfapi/src/domain/users"
	"gitlab.com/chertokdmitry/surfapi/src/utils/crypto_utils"
	"gitlab.com/chertokdmitry/surfapi/src/utils/date_utils"
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
)

var (
	UserService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) ([]users.User, *errors.RestErr)
}

func (c *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
func (c *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}
func (c *usersService) UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := c.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (c *usersService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (c *usersService) Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}

	return dao.FindByStatus(status)
}
