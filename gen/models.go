package gen

import (
	"fmt"
	"reflect"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
)

type CompanyResultType struct {
	EntityResultType
}

type Company struct {
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	Name      *string    `json:"name" gorm:"column:name"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy *string    `json:"createdBy" gorm:"column:createdBy"`

	Employees []*User `json:"employees" gorm:"many2many:company_employees;jointable_foreignkey:company_id;association_jointable_foreignkey:employee_id"`
}

func (m *Company) Is_Entity() {}

type CompanyChanges struct {
	ID        string
	Name      *string
	UpdatedAt *time.Time
	CreatedAt time.Time
	UpdatedBy *string
	CreatedBy *string

	EmployeesIDs []*string
}

type Company_employees struct {
	company_id  string
	employee_id string
}

func (Company_employees) TableName() string {
	return TableName("company_employees")
}

type UserResultType struct {
	EntityResultType
}

type User struct {
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	Email     *string    `json:"email" gorm:"column:email"`
	FirstName *string    `json:"firstName" gorm:"column:firstName"`
	LastName  *string    `json:"lastName" gorm:"column:lastName"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy *string    `json:"createdBy" gorm:"column:createdBy"`

	Tasks []*Task `json:"tasks" gorm:"foreignkey:AssigneeID"`

	Companies []*Company `json:"companies" gorm:"many2many:company_employees;jointable_foreignkey:employee_id;association_jointable_foreignkey:company_id"`
}

func (m *User) Is_Entity() {}

type UserChanges struct {
	ID        string
	Email     *string
	FirstName *string
	LastName  *string
	UpdatedAt *time.Time
	CreatedAt time.Time
	UpdatedBy *string
	CreatedBy *string

	TasksIDs     []*string
	CompaniesIDs []*string
}

type TaskResultType struct {
	EntityResultType
}

type Task struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Title       *string    `json:"title" gorm:"column:title"`
	Completed   *bool      `json:"completed" gorm:"column:completed"`
	DueDate     *time.Time `json:"dueDate" gorm:"column:dueDate"`
	Type        *TaskType  `json:"type" gorm:"column:type"`
	Description *string    `json:"description" gorm:"column:description;type:text"`
	AssigneeID  *string    `json:"assigneeId" gorm:"column:assigneeId"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`

	Assignee *User `json:"assignee"`
}

func (m *Task) Is_Entity() {}

type TaskChanges struct {
	ID          string
	Title       *string
	Completed   *bool
	DueDate     *time.Time
	Type        *TaskType
	Description *string
	AssigneeID  *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string
}

// used to convert map[string]interface{} to EntityChanges struct
func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v interface{}) (interface{}, error) {

			if b == reflect.TypeOf(time.Time{}) {
				switch a.Kind() {
				case reflect.String:
					return time.Parse(time.RFC3339, v.(string))
				case reflect.Float64:
					return time.Unix(0, int64(v.(float64))*int64(time.Millisecond)), nil
				case reflect.Int64:
					return time.Unix(0, v.(int64)*int64(time.Millisecond)), nil
				default:
					return v, fmt.Errorf("Unable to parse date from %v", v)
				}
			}

			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}
