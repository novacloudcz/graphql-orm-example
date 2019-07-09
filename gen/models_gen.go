// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gen

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CompanyFilterType struct {
	And          []*CompanyFilterType `json:"AND"`
	Or           []*CompanyFilterType `json:"OR"`
	ID           *string              `json:"id"`
	IDNe         *string              `json:"id_ne"`
	IDGt         *string              `json:"id_gt"`
	IDLt         *string              `json:"id_lt"`
	IDGte        *string              `json:"id_gte"`
	IDLte        *string              `json:"id_lte"`
	IDIn         []string             `json:"id_in"`
	Name         *string              `json:"name"`
	NameNe       *string              `json:"name_ne"`
	NameGt       *string              `json:"name_gt"`
	NameLt       *string              `json:"name_lt"`
	NameGte      *string              `json:"name_gte"`
	NameLte      *string              `json:"name_lte"`
	NameIn       []string             `json:"name_in"`
	NameLike     *string              `json:"name_like"`
	NamePrefix   *string              `json:"name_prefix"`
	NameSuffix   *string              `json:"name_suffix"`
	UpdatedAt    *time.Time           `json:"updatedAt"`
	UpdatedAtNe  *time.Time           `json:"updatedAt_ne"`
	UpdatedAtGt  *time.Time           `json:"updatedAt_gt"`
	UpdatedAtLt  *time.Time           `json:"updatedAt_lt"`
	UpdatedAtGte *time.Time           `json:"updatedAt_gte"`
	UpdatedAtLte *time.Time           `json:"updatedAt_lte"`
	UpdatedAtIn  []*time.Time         `json:"updatedAt_in"`
	CreatedAt    *time.Time           `json:"createdAt"`
	CreatedAtNe  *time.Time           `json:"createdAt_ne"`
	CreatedAtGt  *time.Time           `json:"createdAt_gt"`
	CreatedAtLt  *time.Time           `json:"createdAt_lt"`
	CreatedAtGte *time.Time           `json:"createdAt_gte"`
	CreatedAtLte *time.Time           `json:"createdAt_lte"`
	CreatedAtIn  []*time.Time         `json:"createdAt_in"`
	UpdatedBy    *string              `json:"updatedBy"`
	UpdatedByNe  *string              `json:"updatedBy_ne"`
	UpdatedByGt  *string              `json:"updatedBy_gt"`
	UpdatedByLt  *string              `json:"updatedBy_lt"`
	UpdatedByGte *string              `json:"updatedBy_gte"`
	UpdatedByLte *string              `json:"updatedBy_lte"`
	UpdatedByIn  []string             `json:"updatedBy_in"`
	CreatedBy    *string              `json:"createdBy"`
	CreatedByNe  *string              `json:"createdBy_ne"`
	CreatedByGt  *string              `json:"createdBy_gt"`
	CreatedByLt  *string              `json:"createdBy_lt"`
	CreatedByGte *string              `json:"createdBy_gte"`
	CreatedByLte *string              `json:"createdBy_lte"`
	CreatedByIn  []string             `json:"createdBy_in"`
	Employees    *UserFilterType      `json:"employees"`
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
	Completed         *bool             `json:"completed"`
	CompletedNe       *bool             `json:"completed_ne"`
	CompletedGt       *bool             `json:"completed_gt"`
	CompletedLt       *bool             `json:"completed_lt"`
	CompletedGte      *bool             `json:"completed_gte"`
	CompletedLte      *bool             `json:"completed_lte"`
	CompletedIn       []bool            `json:"completed_in"`
	DueDate           *time.Time        `json:"dueDate"`
	DueDateNe         *time.Time        `json:"dueDate_ne"`
	DueDateGt         *time.Time        `json:"dueDate_gt"`
	DueDateLt         *time.Time        `json:"dueDate_lt"`
	DueDateGte        *time.Time        `json:"dueDate_gte"`
	DueDateLte        *time.Time        `json:"dueDate_lte"`
	DueDateIn         []*time.Time      `json:"dueDate_in"`
	Type              *TaskType         `json:"type"`
	TypeNe            *TaskType         `json:"type_ne"`
	TypeGt            *TaskType         `json:"type_gt"`
	TypeLt            *TaskType         `json:"type_lt"`
	TypeGte           *TaskType         `json:"type_gte"`
	TypeLte           *TaskType         `json:"type_lte"`
	TypeIn            []TaskType        `json:"type_in"`
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
	AssigneeID        *string           `json:"assigneeId"`
	AssigneeIDNe      *string           `json:"assigneeId_ne"`
	AssigneeIDGt      *string           `json:"assigneeId_gt"`
	AssigneeIDLt      *string           `json:"assigneeId_lt"`
	AssigneeIDGte     *string           `json:"assigneeId_gte"`
	AssigneeIDLte     *string           `json:"assigneeId_lte"`
	AssigneeIDIn      []string          `json:"assigneeId_in"`
	UpdatedAt         *time.Time        `json:"updatedAt"`
	UpdatedAtNe       *time.Time        `json:"updatedAt_ne"`
	UpdatedAtGt       *time.Time        `json:"updatedAt_gt"`
	UpdatedAtLt       *time.Time        `json:"updatedAt_lt"`
	UpdatedAtGte      *time.Time        `json:"updatedAt_gte"`
	UpdatedAtLte      *time.Time        `json:"updatedAt_lte"`
	UpdatedAtIn       []*time.Time      `json:"updatedAt_in"`
	CreatedAt         *time.Time        `json:"createdAt"`
	CreatedAtNe       *time.Time        `json:"createdAt_ne"`
	CreatedAtGt       *time.Time        `json:"createdAt_gt"`
	CreatedAtLt       *time.Time        `json:"createdAt_lt"`
	CreatedAtGte      *time.Time        `json:"createdAt_gte"`
	CreatedAtLte      *time.Time        `json:"createdAt_lte"`
	CreatedAtIn       []*time.Time      `json:"createdAt_in"`
	UpdatedBy         *string           `json:"updatedBy"`
	UpdatedByNe       *string           `json:"updatedBy_ne"`
	UpdatedByGt       *string           `json:"updatedBy_gt"`
	UpdatedByLt       *string           `json:"updatedBy_lt"`
	UpdatedByGte      *string           `json:"updatedBy_gte"`
	UpdatedByLte      *string           `json:"updatedBy_lte"`
	UpdatedByIn       []string          `json:"updatedBy_in"`
	CreatedBy         *string           `json:"createdBy"`
	CreatedByNe       *string           `json:"createdBy_ne"`
	CreatedByGt       *string           `json:"createdBy_gt"`
	CreatedByLt       *string           `json:"createdBy_lt"`
	CreatedByGte      *string           `json:"createdBy_gte"`
	CreatedByLte      *string           `json:"createdBy_lte"`
	CreatedByIn       []string          `json:"createdBy_in"`
	Assignee          *UserFilterType   `json:"assignee"`
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
	UpdatedAt       *time.Time         `json:"updatedAt"`
	UpdatedAtNe     *time.Time         `json:"updatedAt_ne"`
	UpdatedAtGt     *time.Time         `json:"updatedAt_gt"`
	UpdatedAtLt     *time.Time         `json:"updatedAt_lt"`
	UpdatedAtGte    *time.Time         `json:"updatedAt_gte"`
	UpdatedAtLte    *time.Time         `json:"updatedAt_lte"`
	UpdatedAtIn     []*time.Time       `json:"updatedAt_in"`
	CreatedAt       *time.Time         `json:"createdAt"`
	CreatedAtNe     *time.Time         `json:"createdAt_ne"`
	CreatedAtGt     *time.Time         `json:"createdAt_gt"`
	CreatedAtLt     *time.Time         `json:"createdAt_lt"`
	CreatedAtGte    *time.Time         `json:"createdAt_gte"`
	CreatedAtLte    *time.Time         `json:"createdAt_lte"`
	CreatedAtIn     []*time.Time       `json:"createdAt_in"`
	UpdatedBy       *string            `json:"updatedBy"`
	UpdatedByNe     *string            `json:"updatedBy_ne"`
	UpdatedByGt     *string            `json:"updatedBy_gt"`
	UpdatedByLt     *string            `json:"updatedBy_lt"`
	UpdatedByGte    *string            `json:"updatedBy_gte"`
	UpdatedByLte    *string            `json:"updatedBy_lte"`
	UpdatedByIn     []string           `json:"updatedBy_in"`
	CreatedBy       *string            `json:"createdBy"`
	CreatedByNe     *string            `json:"createdBy_ne"`
	CreatedByGt     *string            `json:"createdBy_gt"`
	CreatedByLt     *string            `json:"createdBy_lt"`
	CreatedByGte    *string            `json:"createdBy_gte"`
	CreatedByLte    *string            `json:"createdBy_lte"`
	CreatedByIn     []string           `json:"createdBy_in"`
	Tasks           *TaskFilterType    `json:"tasks"`
	Companies       *CompanyFilterType `json:"companies"`
	Friends         *UserFilterType    `json:"friends"`
}

type CompanySortType string

const (
	CompanySortTypeIDAsc         CompanySortType = "ID_ASC"
	CompanySortTypeIDDesc        CompanySortType = "ID_DESC"
	CompanySortTypeNameAsc       CompanySortType = "NAME_ASC"
	CompanySortTypeNameDesc      CompanySortType = "NAME_DESC"
	CompanySortTypeUpdatedAtAsc  CompanySortType = "UPDATED_AT_ASC"
	CompanySortTypeUpdatedAtDesc CompanySortType = "UPDATED_AT_DESC"
	CompanySortTypeCreatedAtAsc  CompanySortType = "CREATED_AT_ASC"
	CompanySortTypeCreatedAtDesc CompanySortType = "CREATED_AT_DESC"
	CompanySortTypeUpdatedByAsc  CompanySortType = "UPDATED_BY_ASC"
	CompanySortTypeUpdatedByDesc CompanySortType = "UPDATED_BY_DESC"
	CompanySortTypeCreatedByAsc  CompanySortType = "CREATED_BY_ASC"
	CompanySortTypeCreatedByDesc CompanySortType = "CREATED_BY_DESC"
)

var AllCompanySortType = []CompanySortType{
	CompanySortTypeIDAsc,
	CompanySortTypeIDDesc,
	CompanySortTypeNameAsc,
	CompanySortTypeNameDesc,
	CompanySortTypeUpdatedAtAsc,
	CompanySortTypeUpdatedAtDesc,
	CompanySortTypeCreatedAtAsc,
	CompanySortTypeCreatedAtDesc,
	CompanySortTypeUpdatedByAsc,
	CompanySortTypeUpdatedByDesc,
	CompanySortTypeCreatedByAsc,
	CompanySortTypeCreatedByDesc,
}

func (e CompanySortType) IsValid() bool {
	switch e {
	case CompanySortTypeIDAsc, CompanySortTypeIDDesc, CompanySortTypeNameAsc, CompanySortTypeNameDesc, CompanySortTypeUpdatedAtAsc, CompanySortTypeUpdatedAtDesc, CompanySortTypeCreatedAtAsc, CompanySortTypeCreatedAtDesc, CompanySortTypeUpdatedByAsc, CompanySortTypeUpdatedByDesc, CompanySortTypeCreatedByAsc, CompanySortTypeCreatedByDesc:
		return true
	}
	return false
}

func (e CompanySortType) String() string {
	return string(e)
}

func (e *CompanySortType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CompanySortType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CompanySortType", str)
	}
	return nil
}

func (e CompanySortType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskSortType string

const (
	TaskSortTypeIDAsc           TaskSortType = "ID_ASC"
	TaskSortTypeIDDesc          TaskSortType = "ID_DESC"
	TaskSortTypeTitleAsc        TaskSortType = "TITLE_ASC"
	TaskSortTypeTitleDesc       TaskSortType = "TITLE_DESC"
	TaskSortTypeCompletedAsc    TaskSortType = "COMPLETED_ASC"
	TaskSortTypeCompletedDesc   TaskSortType = "COMPLETED_DESC"
	TaskSortTypeDueDateAsc      TaskSortType = "DUE_DATE_ASC"
	TaskSortTypeDueDateDesc     TaskSortType = "DUE_DATE_DESC"
	TaskSortTypeTypeAsc         TaskSortType = "TYPE_ASC"
	TaskSortTypeTypeDesc        TaskSortType = "TYPE_DESC"
	TaskSortTypeDescriptionAsc  TaskSortType = "DESCRIPTION_ASC"
	TaskSortTypeDescriptionDesc TaskSortType = "DESCRIPTION_DESC"
	TaskSortTypeAssigneeIDAsc   TaskSortType = "ASSIGNEE_ID_ASC"
	TaskSortTypeAssigneeIDDesc  TaskSortType = "ASSIGNEE_ID_DESC"
	TaskSortTypeUpdatedAtAsc    TaskSortType = "UPDATED_AT_ASC"
	TaskSortTypeUpdatedAtDesc   TaskSortType = "UPDATED_AT_DESC"
	TaskSortTypeCreatedAtAsc    TaskSortType = "CREATED_AT_ASC"
	TaskSortTypeCreatedAtDesc   TaskSortType = "CREATED_AT_DESC"
	TaskSortTypeUpdatedByAsc    TaskSortType = "UPDATED_BY_ASC"
	TaskSortTypeUpdatedByDesc   TaskSortType = "UPDATED_BY_DESC"
	TaskSortTypeCreatedByAsc    TaskSortType = "CREATED_BY_ASC"
	TaskSortTypeCreatedByDesc   TaskSortType = "CREATED_BY_DESC"
)

var AllTaskSortType = []TaskSortType{
	TaskSortTypeIDAsc,
	TaskSortTypeIDDesc,
	TaskSortTypeTitleAsc,
	TaskSortTypeTitleDesc,
	TaskSortTypeCompletedAsc,
	TaskSortTypeCompletedDesc,
	TaskSortTypeDueDateAsc,
	TaskSortTypeDueDateDesc,
	TaskSortTypeTypeAsc,
	TaskSortTypeTypeDesc,
	TaskSortTypeDescriptionAsc,
	TaskSortTypeDescriptionDesc,
	TaskSortTypeAssigneeIDAsc,
	TaskSortTypeAssigneeIDDesc,
	TaskSortTypeUpdatedAtAsc,
	TaskSortTypeUpdatedAtDesc,
	TaskSortTypeCreatedAtAsc,
	TaskSortTypeCreatedAtDesc,
	TaskSortTypeUpdatedByAsc,
	TaskSortTypeUpdatedByDesc,
	TaskSortTypeCreatedByAsc,
	TaskSortTypeCreatedByDesc,
}

func (e TaskSortType) IsValid() bool {
	switch e {
	case TaskSortTypeIDAsc, TaskSortTypeIDDesc, TaskSortTypeTitleAsc, TaskSortTypeTitleDesc, TaskSortTypeCompletedAsc, TaskSortTypeCompletedDesc, TaskSortTypeDueDateAsc, TaskSortTypeDueDateDesc, TaskSortTypeTypeAsc, TaskSortTypeTypeDesc, TaskSortTypeDescriptionAsc, TaskSortTypeDescriptionDesc, TaskSortTypeAssigneeIDAsc, TaskSortTypeAssigneeIDDesc, TaskSortTypeUpdatedAtAsc, TaskSortTypeUpdatedAtDesc, TaskSortTypeCreatedAtAsc, TaskSortTypeCreatedAtDesc, TaskSortTypeUpdatedByAsc, TaskSortTypeUpdatedByDesc, TaskSortTypeCreatedByAsc, TaskSortTypeCreatedByDesc:
		return true
	}
	return false
}

func (e TaskSortType) String() string {
	return string(e)
}

func (e *TaskSortType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskSortType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskSortType", str)
	}
	return nil
}

func (e TaskSortType) MarshalGQL(w io.Writer) {
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

type UserSortType string

const (
	UserSortTypeIDAsc         UserSortType = "ID_ASC"
	UserSortTypeIDDesc        UserSortType = "ID_DESC"
	UserSortTypeEmailAsc      UserSortType = "EMAIL_ASC"
	UserSortTypeEmailDesc     UserSortType = "EMAIL_DESC"
	UserSortTypeFirstNameAsc  UserSortType = "FIRST_NAME_ASC"
	UserSortTypeFirstNameDesc UserSortType = "FIRST_NAME_DESC"
	UserSortTypeLastNameAsc   UserSortType = "LAST_NAME_ASC"
	UserSortTypeLastNameDesc  UserSortType = "LAST_NAME_DESC"
	UserSortTypeUpdatedAtAsc  UserSortType = "UPDATED_AT_ASC"
	UserSortTypeUpdatedAtDesc UserSortType = "UPDATED_AT_DESC"
	UserSortTypeCreatedAtAsc  UserSortType = "CREATED_AT_ASC"
	UserSortTypeCreatedAtDesc UserSortType = "CREATED_AT_DESC"
	UserSortTypeUpdatedByAsc  UserSortType = "UPDATED_BY_ASC"
	UserSortTypeUpdatedByDesc UserSortType = "UPDATED_BY_DESC"
	UserSortTypeCreatedByAsc  UserSortType = "CREATED_BY_ASC"
	UserSortTypeCreatedByDesc UserSortType = "CREATED_BY_DESC"
)

var AllUserSortType = []UserSortType{
	UserSortTypeIDAsc,
	UserSortTypeIDDesc,
	UserSortTypeEmailAsc,
	UserSortTypeEmailDesc,
	UserSortTypeFirstNameAsc,
	UserSortTypeFirstNameDesc,
	UserSortTypeLastNameAsc,
	UserSortTypeLastNameDesc,
	UserSortTypeUpdatedAtAsc,
	UserSortTypeUpdatedAtDesc,
	UserSortTypeCreatedAtAsc,
	UserSortTypeCreatedAtDesc,
	UserSortTypeUpdatedByAsc,
	UserSortTypeUpdatedByDesc,
	UserSortTypeCreatedByAsc,
	UserSortTypeCreatedByDesc,
}

func (e UserSortType) IsValid() bool {
	switch e {
	case UserSortTypeIDAsc, UserSortTypeIDDesc, UserSortTypeEmailAsc, UserSortTypeEmailDesc, UserSortTypeFirstNameAsc, UserSortTypeFirstNameDesc, UserSortTypeLastNameAsc, UserSortTypeLastNameDesc, UserSortTypeUpdatedAtAsc, UserSortTypeUpdatedAtDesc, UserSortTypeCreatedAtAsc, UserSortTypeCreatedAtDesc, UserSortTypeUpdatedByAsc, UserSortTypeUpdatedByDesc, UserSortTypeCreatedByAsc, UserSortTypeCreatedByDesc:
		return true
	}
	return false
}

func (e UserSortType) String() string {
	return string(e)
}

func (e *UserSortType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserSortType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserSortType", str)
	}
	return nil
}

func (e UserSortType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
