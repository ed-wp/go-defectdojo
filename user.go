package defectdojo

import (
	"context"
	"fmt"
)

const (
	//lint:ignore U1000 todo
	usersAPIBase = "/api/v2/users/"
)

type User struct {
	UserName    string `json:"username" url:"username"`
	FirstName   string `json:"first_name,omitempty" url:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty" url:"last_name,omitempty"`
	Email       string `json:"email,omitempty" url:"email,omitempty"`
	IsActive    bool   `json:"is_active,omitempty" url:"is_active,omitempty"`
	IsStaff     bool   `json:"is_staff,omitempty" url:"is_staff,omitempty"`
	IsSuperuser bool   `json:"is_superuser,omitempty" url:"is_superuser,omitempty"`
}

func (d *DefectDojoAPI) GetUsers(ctx context.Context, user *User) ([]*User, error) {

	return make([]*User, 0), fmt.Errorf("not implemented")
}

func (d *DefectDojoAPI) AddUser(ctx context.Context, user *User) error {

	return fmt.Errorf("not implemented")
}

func (d *DefectDojoAPI) UpdateUser(ctx context.Context, user *User) error {

	return fmt.Errorf("not implemented")
}

func (d *DefectDojoAPI) RemoveUser(ctx context.Context, user *User) error {

	return fmt.Errorf("not implemented")
}
