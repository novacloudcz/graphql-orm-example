package gen

import (
	"time"
)

type Todo struct {
	ID string `json:"id" gorm:"primary_key"`

	State *TodoState `json:"state"`
	Text  *string    `json:"text"`
	Blah  *float64   `json:"blah"`
	Done  bool       `json:"done"`

	User   *User `json:"user"`
	UserID string

	UpdatedAt time.Time  `json:"updatedAt"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type User struct {
	ID string `json:"id" gorm:"primary_key"`

	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`

	Todos []*Todo `json:"todos" gorm:"foreignkey:UserID"`

	Friends []*User `json:"friends" gorm:"many2many:user_friends;association_jointable_foreignkey:friend_id"`

	Employers []*Company `json:"employers" gorm:"many2many:user_employers;"`

	UpdatedAt time.Time  `json:"updatedAt"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type Company struct {
	ID string `json:"id" gorm:"primary_key"`

	Name *string `json:"name"`

	Employees []*User `json:"employees" gorm:"many2many:user_employers;"`

	UpdatedAt time.Time  `json:"updatedAt"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
