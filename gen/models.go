package gen

import (
	"github.com/novacloudcz/graphql-orm/resolvers"
	"time"
)

type CompanyResultType struct {
	resolvers.EntityResultType
}

type Company struct {
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	Name      *string    `json:"name" gorm:"column:name"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy string     `json:"createdBy" gorm:"column:createdBy"`

	Employees []*User `json:"employees" gorm:"many2many:company_employees;jointable_foreignkey:employee_id;association_jointable_foreignkey:company_id"`
}

type UserResultType struct {
	resolvers.EntityResultType
}

type User struct {
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	Email     *string    `json:"email" gorm:"column:email"`
	FirstName *string    `json:"firstName" gorm:"column:firstName"`
	LastName  *string    `json:"lastName" gorm:"column:lastName"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy string     `json:"createdBy" gorm:"column:createdBy"`

	Tasks []*Task `json:"tasks" gorm:"foreignkey:AssigneeID"`

	Companies []*Company `json:"companies" gorm:"many2many:company_employees;jointable_foreignkey:employee_id;association_jointable_foreignkey:company_id"`

	Friends []*User `json:"friends" gorm:"many2many:user_friends;jointable_foreignkey:user_id;association_jointable_foreignkey:friend_id"`
}

type TaskResultType struct {
	resolvers.EntityResultType
}

type Task struct {
	ID         string     `json:"id" gorm:"column:id;primary_key"`
	Title      *string    `json:"title" gorm:"column:title"`
	Completed  *bool      `json:"completed" gorm:"column:completed"`
	DueDate    *time.Time `json:"dueDate" gorm:"column:dueDate"`
	Type       *TaskType  `json:"type" gorm:"column:type"`
	AssigneeID *string    `json:"assigneeId" gorm:"column:assigneeId"`
	UpdatedAt  *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt  time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy  *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy  string     `json:"createdBy" gorm:"column:createdBy"`

	Assignee *User `json:"assignee"`
}
