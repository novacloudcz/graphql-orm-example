package gen

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/novacloudcz/graphql-orm/resolvers"
	"github.com/vektah/gqlparser/ast"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

type QueryCompanyHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *CompanyFilterType
}

func (r *GeneratedQueryResolver) Company(ctx context.Context, id *string, q *string, filter *CompanyFilterType) (*Company, error) {
	opts := QueryCompanyHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryCompany(ctx, r.GeneratedResolver, opts)
}
func QueryCompanyHandler(ctx context.Context, r *GeneratedResolver, opts QueryCompanyHandlerOptions) (*Company, error) {
	query := CompanyQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &CompanyResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where("companies.id = ?", *opts.ID)
	}

	var items []*Company
	err := rt.GetItems(ctx, qb, "companies", &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("Company not found")
	}
	return items[0], err
}

type QueryCompaniesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []CompanySortType
	Filter *CompanyFilterType
}

func (r *GeneratedQueryResolver) Companies(ctx context.Context, offset *int, limit *int, q *string, sort []CompanySortType, filter *CompanyFilterType) (*CompanyResultType, error) {
	opts := QueryCompaniesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryCompanies(ctx, r.GeneratedResolver, opts)
}
func QueryCompaniesHandler(ctx context.Context, r *GeneratedResolver, opts QueryCompaniesHandlerOptions) (*CompanyResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range opts.Sort {
		_sort = append(_sort, s)
	}
	query := CompanyQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	return &CompanyResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedCompanyResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedCompanyResultTypeResolver) Items(ctx context.Context, obj *CompanyResultType) (items []*Company, err error) {
	err = obj.GetItems(ctx, r.DB.db, "companies", &items)
	return
}

func (r *GeneratedCompanyResultTypeResolver) Count(ctx context.Context, obj *CompanyResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Company{})
}

type GeneratedCompanyResolver struct{ *GeneratedResolver }

func (r *GeneratedCompanyResolver) Employees(ctx context.Context, obj *Company) (res []*User, err error) {
	return r.Handlers.CompanyEmployees(ctx, r, obj)
}
func CompanyEmployeesHandler(ctx context.Context, r *GeneratedCompanyResolver, obj *Company) (res []*User, err error) {

	items := []*User{}
	err = r.DB.Query().Model(obj).Related(&items, "Employees").Error
	res = items

	return
}

func (r *GeneratedCompanyResolver) EmployeesIds(ctx context.Context, obj *Company) (ids []string, err error) {
	ids = []string{}

	items := []*User{}
	err = r.DB.Query().Model(obj).Select("users.id").Related(&items, "Employees").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

type QueryUserHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *UserFilterType
}

func (r *GeneratedQueryResolver) User(ctx context.Context, id *string, q *string, filter *UserFilterType) (*User, error) {
	opts := QueryUserHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryUser(ctx, r.GeneratedResolver, opts)
}
func QueryUserHandler(ctx context.Context, r *GeneratedResolver, opts QueryUserHandlerOptions) (*User, error) {
	query := UserQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &UserResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where("users.id = ?", *opts.ID)
	}

	var items []*User
	err := rt.GetItems(ctx, qb, "users", &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("User not found")
	}
	return items[0], err
}

type QueryUsersHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []UserSortType
	Filter *UserFilterType
}

func (r *GeneratedQueryResolver) Users(ctx context.Context, offset *int, limit *int, q *string, sort []UserSortType, filter *UserFilterType) (*UserResultType, error) {
	opts := QueryUsersHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryUsers(ctx, r.GeneratedResolver, opts)
}
func QueryUsersHandler(ctx context.Context, r *GeneratedResolver, opts QueryUsersHandlerOptions) (*UserResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range opts.Sort {
		_sort = append(_sort, s)
	}
	query := UserQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	return &UserResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedUserResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedUserResultTypeResolver) Items(ctx context.Context, obj *UserResultType) (items []*User, err error) {
	err = obj.GetItems(ctx, r.DB.db, "users", &items)
	return
}

func (r *GeneratedUserResultTypeResolver) Count(ctx context.Context, obj *UserResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &User{})
}

type GeneratedUserResolver struct{ *GeneratedResolver }

func (r *GeneratedUserResolver) Tasks(ctx context.Context, obj *User) (res []*Task, err error) {
	return r.Handlers.UserTasks(ctx, r, obj)
}
func UserTasksHandler(ctx context.Context, r *GeneratedUserResolver, obj *User) (res []*Task, err error) {

	items := []*Task{}
	err = r.DB.Query().Model(obj).Related(&items, "Tasks").Error
	res = items

	return
}

func (r *GeneratedUserResolver) TasksIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}

	items := []*Task{}
	err = r.DB.Query().Model(obj).Select("tasks.id").Related(&items, "Tasks").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

func (r *GeneratedUserResolver) Companies(ctx context.Context, obj *User) (res []*Company, err error) {
	return r.Handlers.UserCompanies(ctx, r, obj)
}
func UserCompaniesHandler(ctx context.Context, r *GeneratedUserResolver, obj *User) (res []*Company, err error) {

	items := []*Company{}
	err = r.DB.Query().Model(obj).Related(&items, "Companies").Error
	res = items

	return
}

func (r *GeneratedUserResolver) CompaniesIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}

	items := []*Company{}
	err = r.DB.Query().Model(obj).Select("companies.id").Related(&items, "Companies").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

func (r *GeneratedUserResolver) Friends(ctx context.Context, obj *User) (res []*User, err error) {
	return r.Handlers.UserFriends(ctx, r, obj)
}
func UserFriendsHandler(ctx context.Context, r *GeneratedUserResolver, obj *User) (res []*User, err error) {

	items := []*User{}
	err = r.DB.Query().Model(obj).Related(&items, "Friends").Error
	res = items

	return
}

func (r *GeneratedUserResolver) FriendsIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}

	items := []*User{}
	err = r.DB.Query().Model(obj).Select("users.id").Related(&items, "Friends").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

type QueryTaskHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *TaskFilterType
}

func (r *GeneratedQueryResolver) Task(ctx context.Context, id *string, q *string, filter *TaskFilterType) (*Task, error) {
	opts := QueryTaskHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryTask(ctx, r.GeneratedResolver, opts)
}
func QueryTaskHandler(ctx context.Context, r *GeneratedResolver, opts QueryTaskHandlerOptions) (*Task, error) {
	query := TaskQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &TaskResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where("tasks.id = ?", *opts.ID)
	}

	var items []*Task
	err := rt.GetItems(ctx, qb, "tasks", &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("Task not found")
	}
	return items[0], err
}

type QueryTasksHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []TaskSortType
	Filter *TaskFilterType
}

func (r *GeneratedQueryResolver) Tasks(ctx context.Context, offset *int, limit *int, q *string, sort []TaskSortType, filter *TaskFilterType) (*TaskResultType, error) {
	opts := QueryTasksHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryTasks(ctx, r.GeneratedResolver, opts)
}
func QueryTasksHandler(ctx context.Context, r *GeneratedResolver, opts QueryTasksHandlerOptions) (*TaskResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range opts.Sort {
		_sort = append(_sort, s)
	}
	query := TaskQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	return &TaskResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedTaskResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedTaskResultTypeResolver) Items(ctx context.Context, obj *TaskResultType) (items []*Task, err error) {
	err = obj.GetItems(ctx, r.DB.db, "tasks", &items)
	return
}

func (r *GeneratedTaskResultTypeResolver) Count(ctx context.Context, obj *TaskResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Task{})
}

type GeneratedTaskResolver struct{ *GeneratedResolver }

func (r *GeneratedTaskResolver) Assignee(ctx context.Context, obj *Task) (res *User, err error) {
	return r.Handlers.TaskAssignee(ctx, r, obj)
}
func TaskAssigneeHandler(ctx context.Context, r *GeneratedTaskResolver, obj *Task) (res *User, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.AssigneeID != nil {
		item, _err := loaders["User"].Load(ctx, dataloader.StringKey(*obj.AssigneeID))()
		res, _ = item.(*User)
		err = _err
	}

	return
}
