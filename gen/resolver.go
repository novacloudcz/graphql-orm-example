package gen

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/novacloudcz/graphql-orm/events"
)

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *events.Event) error

	CreateCompany      func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Company, err error)
	UpdateCompany      func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Company, err error)
	DeleteCompany      func(ctx context.Context, r *GeneratedResolver, id string) (item *Company, err error)
	DeleteAllCompanies func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryCompany       func(ctx context.Context, r *GeneratedResolver, opts QueryCompanyHandlerOptions) (*Company, error)
	QueryCompanies     func(ctx context.Context, r *GeneratedResolver, opts QueryCompaniesHandlerOptions) (*CompanyResultType, error)

	CompanyEmployees func(ctx context.Context, r *GeneratedResolver, obj *Company) (res []*User, err error)

	CreateUser     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *User, err error)
	UpdateUser     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *User, err error)
	DeleteUser     func(ctx context.Context, r *GeneratedResolver, id string) (item *User, err error)
	DeleteAllUsers func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryUser      func(ctx context.Context, r *GeneratedResolver, opts QueryUserHandlerOptions) (*User, error)
	QueryUsers     func(ctx context.Context, r *GeneratedResolver, opts QueryUsersHandlerOptions) (*UserResultType, error)

	UserTasks func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Task, err error)

	UserCompanies func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Company, err error)

	CreateTask     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Task, err error)
	UpdateTask     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Task, err error)
	DeleteTask     func(ctx context.Context, r *GeneratedResolver, id string) (item *Task, err error)
	DeleteAllTasks func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryTask      func(ctx context.Context, r *GeneratedResolver, opts QueryTaskHandlerOptions) (*Task, error)
	QueryTasks     func(ctx context.Context, r *GeneratedResolver, opts QueryTasksHandlerOptions) (*TaskResultType, error)

	TaskAssignee func(ctx context.Context, r *GeneratedResolver, obj *Task) (res *User, err error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *events.Event) error { return nil },

		CreateCompany:      CreateCompanyHandler,
		UpdateCompany:      UpdateCompanyHandler,
		DeleteCompany:      DeleteCompanyHandler,
		DeleteAllCompanies: DeleteAllCompaniesHandler,
		QueryCompany:       QueryCompanyHandler,
		QueryCompanies:     QueryCompaniesHandler,

		CompanyEmployees: CompanyEmployeesHandler,

		CreateUser:     CreateUserHandler,
		UpdateUser:     UpdateUserHandler,
		DeleteUser:     DeleteUserHandler,
		DeleteAllUsers: DeleteAllUsersHandler,
		QueryUser:      QueryUserHandler,
		QueryUsers:     QueryUsersHandler,

		UserTasks: UserTasksHandler,

		UserCompanies: UserCompaniesHandler,

		CreateTask:     CreateTaskHandler,
		UpdateTask:     UpdateTaskHandler,
		DeleteTask:     DeleteTaskHandler,
		DeleteAllTasks: DeleteAllTasksHandler,
		QueryTask:      QueryTaskHandler,
		QueryTasks:     QueryTasksHandler,

		TaskAssignee: TaskAssigneeHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *EventController
}

// GetDB returns database connection or transaction for given context (if exists)
func (r *GeneratedResolver) GetDB(ctx context.Context) *gorm.DB {
	db, _ := ctx.Value(KeyMutationTransaction).(*gorm.DB)
	if db == nil {
		db = r.DB.Query()
	}
	return db
}
