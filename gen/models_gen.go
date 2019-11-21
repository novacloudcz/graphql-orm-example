// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gen

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CompanyFilterType struct {
	And           []*CompanyFilterType `json:"AND"`
	Or            []*CompanyFilterType `json:"OR"`
	ID            *string              `json:"id"`
	IDNe          *string              `json:"id_ne"`
	IDGt          *string              `json:"id_gt"`
	IDLt          *string              `json:"id_lt"`
	IDGte         *string              `json:"id_gte"`
	IDLte         *string              `json:"id_lte"`
	IDIn          []string             `json:"id_in"`
	IDNull        *bool                `json:"id_null"`
	Name          *string              `json:"name"`
	NameNe        *string              `json:"name_ne"`
	NameGt        *string              `json:"name_gt"`
	NameLt        *string              `json:"name_lt"`
	NameGte       *string              `json:"name_gte"`
	NameLte       *string              `json:"name_lte"`
	NameIn        []string             `json:"name_in"`
	NameLike      *string              `json:"name_like"`
	NamePrefix    *string              `json:"name_prefix"`
	NameSuffix    *string              `json:"name_suffix"`
	NameNull      *bool                `json:"name_null"`
	UpdatedAt     *time.Time           `json:"updatedAt"`
	UpdatedAtNe   *time.Time           `json:"updatedAt_ne"`
	UpdatedAtGt   *time.Time           `json:"updatedAt_gt"`
	UpdatedAtLt   *time.Time           `json:"updatedAt_lt"`
	UpdatedAtGte  *time.Time           `json:"updatedAt_gte"`
	UpdatedAtLte  *time.Time           `json:"updatedAt_lte"`
	UpdatedAtIn   []*time.Time         `json:"updatedAt_in"`
	UpdatedAtNull *bool                `json:"updatedAt_null"`
	CreatedAt     *time.Time           `json:"createdAt"`
	CreatedAtNe   *time.Time           `json:"createdAt_ne"`
	CreatedAtGt   *time.Time           `json:"createdAt_gt"`
	CreatedAtLt   *time.Time           `json:"createdAt_lt"`
	CreatedAtGte  *time.Time           `json:"createdAt_gte"`
	CreatedAtLte  *time.Time           `json:"createdAt_lte"`
	CreatedAtIn   []*time.Time         `json:"createdAt_in"`
	CreatedAtNull *bool                `json:"createdAt_null"`
	UpdatedBy     *string              `json:"updatedBy"`
	UpdatedByNe   *string              `json:"updatedBy_ne"`
	UpdatedByGt   *string              `json:"updatedBy_gt"`
	UpdatedByLt   *string              `json:"updatedBy_lt"`
	UpdatedByGte  *string              `json:"updatedBy_gte"`
	UpdatedByLte  *string              `json:"updatedBy_lte"`
	UpdatedByIn   []string             `json:"updatedBy_in"`
	UpdatedByNull *bool                `json:"updatedBy_null"`
	CreatedBy     *string              `json:"createdBy"`
	CreatedByNe   *string              `json:"createdBy_ne"`
	CreatedByGt   *string              `json:"createdBy_gt"`
	CreatedByLt   *string              `json:"createdBy_lt"`
	CreatedByGte  *string              `json:"createdBy_gte"`
	CreatedByLte  *string              `json:"createdBy_lte"`
	CreatedByIn   []string             `json:"createdBy_in"`
	CreatedByNull *bool                `json:"createdBy_null"`
	Employees     *UserFilterType      `json:"employees"`
}

type CompanySortType struct {
	ID           *ObjectSortType `json:"id"`
	Name         *ObjectSortType `json:"name"`
	UpdatedAt    *ObjectSortType `json:"updatedAt"`
	CreatedAt    *ObjectSortType `json:"createdAt"`
	UpdatedBy    *ObjectSortType `json:"updatedBy"`
	CreatedBy    *ObjectSortType `json:"createdBy"`
	EmployeesIds *ObjectSortType `json:"employeesIds"`
	Employees    *UserSortType   `json:"employees"`
}

type TaskFilterType struct {
	And               []*TaskFilterType `json:"AND"`
	Or                []*TaskFilterType `json:"OR"`
	ID                *string           `json:"id"`
	IDNe              *string           `json:"id_ne"`
	IDGt              *string           `json:"id_gt"`
	IDLt              *string           `json:"id_lt"`
	IDGte             *string           `json:"id_gte"`
	IDLte             *string           `json:"id_lte"`
	IDIn              []string          `json:"id_in"`
	IDNull            *bool             `json:"id_null"`
	Title             *string           `json:"title"`
	TitleNe           *string           `json:"title_ne"`
	TitleGt           *string           `json:"title_gt"`
	TitleLt           *string           `json:"title_lt"`
	TitleGte          *string           `json:"title_gte"`
	TitleLte          *string           `json:"title_lte"`
	TitleIn           []string          `json:"title_in"`
	TitleLike         *string           `json:"title_like"`
	TitlePrefix       *string           `json:"title_prefix"`
	TitleSuffix       *string           `json:"title_suffix"`
	TitleNull         *bool             `json:"title_null"`
	Completed         *bool             `json:"completed"`
	CompletedNe       *bool             `json:"completed_ne"`
	CompletedGt       *bool             `json:"completed_gt"`
	CompletedLt       *bool             `json:"completed_lt"`
	CompletedGte      *bool             `json:"completed_gte"`
	CompletedLte      *bool             `json:"completed_lte"`
	CompletedIn       []bool            `json:"completed_in"`
	CompletedNull     *bool             `json:"completed_null"`
	DueDate           *time.Time        `json:"dueDate"`
	DueDateNe         *time.Time        `json:"dueDate_ne"`
	DueDateGt         *time.Time        `json:"dueDate_gt"`
	DueDateLt         *time.Time        `json:"dueDate_lt"`
	DueDateGte        *time.Time        `json:"dueDate_gte"`
	DueDateLte        *time.Time        `json:"dueDate_lte"`
	DueDateIn         []*time.Time      `json:"dueDate_in"`
	DueDateNull       *bool             `json:"dueDate_null"`
	Type              *TaskType         `json:"type"`
	TypeNe            *TaskType         `json:"type_ne"`
	TypeGt            *TaskType         `json:"type_gt"`
	TypeLt            *TaskType         `json:"type_lt"`
	TypeGte           *TaskType         `json:"type_gte"`
	TypeLte           *TaskType         `json:"type_lte"`
	TypeIn            []TaskType        `json:"type_in"`
	TypeNull          *bool             `json:"type_null"`
	Description       *string           `json:"description"`
	DescriptionNe     *string           `json:"description_ne"`
	DescriptionGt     *string           `json:"description_gt"`
	DescriptionLt     *string           `json:"description_lt"`
	DescriptionGte    *string           `json:"description_gte"`
	DescriptionLte    *string           `json:"description_lte"`
	DescriptionIn     []string          `json:"description_in"`
	DescriptionLike   *string           `json:"description_like"`
	DescriptionPrefix *string           `json:"description_prefix"`
	DescriptionSuffix *string           `json:"description_suffix"`
	DescriptionNull   *bool             `json:"description_null"`
	AssigneeID        *string           `json:"assigneeId"`
	AssigneeIDNe      *string           `json:"assigneeId_ne"`
	AssigneeIDGt      *string           `json:"assigneeId_gt"`
	AssigneeIDLt      *string           `json:"assigneeId_lt"`
	AssigneeIDGte     *string           `json:"assigneeId_gte"`
	AssigneeIDLte     *string           `json:"assigneeId_lte"`
	AssigneeIDIn      []string          `json:"assigneeId_in"`
	AssigneeIDNull    *bool             `json:"assigneeId_null"`
	UpdatedAt         *time.Time        `json:"updatedAt"`
	UpdatedAtNe       *time.Time        `json:"updatedAt_ne"`
	UpdatedAtGt       *time.Time        `json:"updatedAt_gt"`
	UpdatedAtLt       *time.Time        `json:"updatedAt_lt"`
	UpdatedAtGte      *time.Time        `json:"updatedAt_gte"`
	UpdatedAtLte      *time.Time        `json:"updatedAt_lte"`
	UpdatedAtIn       []*time.Time      `json:"updatedAt_in"`
	UpdatedAtNull     *bool             `json:"updatedAt_null"`
	CreatedAt         *time.Time        `json:"createdAt"`
	CreatedAtNe       *time.Time        `json:"createdAt_ne"`
	CreatedAtGt       *time.Time        `json:"createdAt_gt"`
	CreatedAtLt       *time.Time        `json:"createdAt_lt"`
	CreatedAtGte      *time.Time        `json:"createdAt_gte"`
	CreatedAtLte      *time.Time        `json:"createdAt_lte"`
	CreatedAtIn       []*time.Time      `json:"createdAt_in"`
	CreatedAtNull     *bool             `json:"createdAt_null"`
	UpdatedBy         *string           `json:"updatedBy"`
	UpdatedByNe       *string           `json:"updatedBy_ne"`
	UpdatedByGt       *string           `json:"updatedBy_gt"`
	UpdatedByLt       *string           `json:"updatedBy_lt"`
	UpdatedByGte      *string           `json:"updatedBy_gte"`
	UpdatedByLte      *string           `json:"updatedBy_lte"`
	UpdatedByIn       []string          `json:"updatedBy_in"`
	UpdatedByNull     *bool             `json:"updatedBy_null"`
	CreatedBy         *string           `json:"createdBy"`
	CreatedByNe       *string           `json:"createdBy_ne"`
	CreatedByGt       *string           `json:"createdBy_gt"`
	CreatedByLt       *string           `json:"createdBy_lt"`
	CreatedByGte      *string           `json:"createdBy_gte"`
	CreatedByLte      *string           `json:"createdBy_lte"`
	CreatedByIn       []string          `json:"createdBy_in"`
	CreatedByNull     *bool             `json:"createdBy_null"`
	Assignee          *UserFilterType   `json:"assignee"`
}

type TaskSortType struct {
	ID          *ObjectSortType `json:"id"`
	Title       *ObjectSortType `json:"title"`
	Completed   *ObjectSortType `json:"completed"`
	DueDate     *ObjectSortType `json:"dueDate"`
	Type        *ObjectSortType `json:"type"`
	Description *ObjectSortType `json:"description"`
	AssigneeID  *ObjectSortType `json:"assigneeId"`
	UpdatedAt   *ObjectSortType `json:"updatedAt"`
	CreatedAt   *ObjectSortType `json:"createdAt"`
	UpdatedBy   *ObjectSortType `json:"updatedBy"`
	CreatedBy   *ObjectSortType `json:"createdBy"`
	Assignee    *UserSortType   `json:"assignee"`
}

type UserFilterType struct {
	And             []*UserFilterType  `json:"AND"`
	Or              []*UserFilterType  `json:"OR"`
	ID              *string            `json:"id"`
	IDNe            *string            `json:"id_ne"`
	IDGt            *string            `json:"id_gt"`
	IDLt            *string            `json:"id_lt"`
	IDGte           *string            `json:"id_gte"`
	IDLte           *string            `json:"id_lte"`
	IDIn            []string           `json:"id_in"`
	IDNull          *bool              `json:"id_null"`
	Email           *string            `json:"email"`
	EmailNe         *string            `json:"email_ne"`
	EmailGt         *string            `json:"email_gt"`
	EmailLt         *string            `json:"email_lt"`
	EmailGte        *string            `json:"email_gte"`
	EmailLte        *string            `json:"email_lte"`
	EmailIn         []string           `json:"email_in"`
	EmailLike       *string            `json:"email_like"`
	EmailPrefix     *string            `json:"email_prefix"`
	EmailSuffix     *string            `json:"email_suffix"`
	EmailNull       *bool              `json:"email_null"`
	FirstName       *string            `json:"firstName"`
	FirstNameNe     *string            `json:"firstName_ne"`
	FirstNameGt     *string            `json:"firstName_gt"`
	FirstNameLt     *string            `json:"firstName_lt"`
	FirstNameGte    *string            `json:"firstName_gte"`
	FirstNameLte    *string            `json:"firstName_lte"`
	FirstNameIn     []string           `json:"firstName_in"`
	FirstNameLike   *string            `json:"firstName_like"`
	FirstNamePrefix *string            `json:"firstName_prefix"`
	FirstNameSuffix *string            `json:"firstName_suffix"`
	FirstNameNull   *bool              `json:"firstName_null"`
	LastName        *string            `json:"lastName"`
	LastNameNe      *string            `json:"lastName_ne"`
	LastNameGt      *string            `json:"lastName_gt"`
	LastNameLt      *string            `json:"lastName_lt"`
	LastNameGte     *string            `json:"lastName_gte"`
	LastNameLte     *string            `json:"lastName_lte"`
	LastNameIn      []string           `json:"lastName_in"`
	LastNameLike    *string            `json:"lastName_like"`
	LastNamePrefix  *string            `json:"lastName_prefix"`
	LastNameSuffix  *string            `json:"lastName_suffix"`
	LastNameNull    *bool              `json:"lastName_null"`
	UpdatedAt       *time.Time         `json:"updatedAt"`
	UpdatedAtNe     *time.Time         `json:"updatedAt_ne"`
	UpdatedAtGt     *time.Time         `json:"updatedAt_gt"`
	UpdatedAtLt     *time.Time         `json:"updatedAt_lt"`
	UpdatedAtGte    *time.Time         `json:"updatedAt_gte"`
	UpdatedAtLte    *time.Time         `json:"updatedAt_lte"`
	UpdatedAtIn     []*time.Time       `json:"updatedAt_in"`
	UpdatedAtNull   *bool              `json:"updatedAt_null"`
	CreatedAt       *time.Time         `json:"createdAt"`
	CreatedAtNe     *time.Time         `json:"createdAt_ne"`
	CreatedAtGt     *time.Time         `json:"createdAt_gt"`
	CreatedAtLt     *time.Time         `json:"createdAt_lt"`
	CreatedAtGte    *time.Time         `json:"createdAt_gte"`
	CreatedAtLte    *time.Time         `json:"createdAt_lte"`
	CreatedAtIn     []*time.Time       `json:"createdAt_in"`
	CreatedAtNull   *bool              `json:"createdAt_null"`
	UpdatedBy       *string            `json:"updatedBy"`
	UpdatedByNe     *string            `json:"updatedBy_ne"`
	UpdatedByGt     *string            `json:"updatedBy_gt"`
	UpdatedByLt     *string            `json:"updatedBy_lt"`
	UpdatedByGte    *string            `json:"updatedBy_gte"`
	UpdatedByLte    *string            `json:"updatedBy_lte"`
	UpdatedByIn     []string           `json:"updatedBy_in"`
	UpdatedByNull   *bool              `json:"updatedBy_null"`
	CreatedBy       *string            `json:"createdBy"`
	CreatedByNe     *string            `json:"createdBy_ne"`
	CreatedByGt     *string            `json:"createdBy_gt"`
	CreatedByLt     *string            `json:"createdBy_lt"`
	CreatedByGte    *string            `json:"createdBy_gte"`
	CreatedByLte    *string            `json:"createdBy_lte"`
	CreatedByIn     []string           `json:"createdBy_in"`
	CreatedByNull   *bool              `json:"createdBy_null"`
	Tasks           *TaskFilterType    `json:"tasks"`
	Companies       *CompanyFilterType `json:"companies"`
	Friends         *UserFilterType    `json:"friends"`
}

type UserSortType struct {
	ID           *ObjectSortType  `json:"id"`
	Email        *ObjectSortType  `json:"email"`
	FirstName    *ObjectSortType  `json:"firstName"`
	LastName     *ObjectSortType  `json:"lastName"`
	UpdatedAt    *ObjectSortType  `json:"updatedAt"`
	CreatedAt    *ObjectSortType  `json:"createdAt"`
	UpdatedBy    *ObjectSortType  `json:"updatedBy"`
	CreatedBy    *ObjectSortType  `json:"createdBy"`
	TasksIds     *ObjectSortType  `json:"tasksIds"`
	CompaniesIds *ObjectSortType  `json:"companiesIds"`
	FriendsIds   *ObjectSortType  `json:"friendsIds"`
	Tasks        *TaskSortType    `json:"tasks"`
	Companies    *CompanySortType `json:"companies"`
	Friends      *UserSortType    `json:"friends"`
}

type _Service struct {
	Sdl *string `json:"sdl"`
}

type ObjectSortType string

const (
	ObjectSortTypeAsc  ObjectSortType = "ASC"
	ObjectSortTypeDesc ObjectSortType = "DESC"
)

var AllObjectSortType = []ObjectSortType{
	ObjectSortTypeAsc,
	ObjectSortTypeDesc,
}

func (e ObjectSortType) IsValid() bool {
	switch e {
	case ObjectSortTypeAsc, ObjectSortTypeDesc:
		return true
	}
	return false
}

func (e ObjectSortType) String() string {
	return string(e)
}

func (e *ObjectSortType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ObjectSortType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ObjectSortType", str)
	}
	return nil
}

func (e ObjectSortType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskType string

const (
	TaskTypeBug  TaskType = "BUG"
	TaskTypeTask TaskType = "TASK"
)

var AllTaskType = []TaskType{
	TaskTypeBug,
	TaskTypeTask,
}

func (e TaskType) IsValid() bool {
	switch e {
	case TaskTypeBug, TaskTypeTask:
		return true
	}
	return false
}

func (e TaskType) String() string {
	return string(e)
}

func (e *TaskType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskType", str)
	}
	return nil
}

func (e TaskType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
