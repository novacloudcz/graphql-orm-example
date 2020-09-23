package gen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
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
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := CompanyQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &CompanyResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("companies")+".id = ?", *opts.ID)
	}

	var items []*Company
	giOpts := GetItemsOptions{
		Alias:      TableName("companies"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

type QueryCompaniesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*CompanySortType
	Filter *CompanyFilterType
}

func (r *GeneratedQueryResolver) Companies(ctx context.Context, offset *int, limit *int, q *string, sort []*CompanySortType, filter *CompanyFilterType) (*CompanyResultType, error) {
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
	query := CompanyQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &CompanyResultType{
		EntityResultType: EntityResultType{
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
	otps := GetItemsOptions{
		Alias:      TableName("companies"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*Company{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

func (r *GeneratedCompanyResultTypeResolver) Count(ctx context.Context, obj *CompanyResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("companies"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &Company{})
}

type GeneratedCompanyResolver struct{ *GeneratedResolver }

func (r *GeneratedCompanyResolver) Employees(ctx context.Context, obj *Company) (res []*User, err error) {
	return r.Handlers.CompanyEmployees(ctx, r.GeneratedResolver, obj)
}
func CompanyEmployeesHandler(ctx context.Context, r *GeneratedResolver, obj *Company) (res []*User, err error) {

	items := []*User{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "Employees").Error
	res = items

	return
}

func (r *GeneratedCompanyResolver) EmployeesIds(ctx context.Context, obj *Company) (ids []string, err error) {
	ids = []string{}

	items := []*User{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Select(TableName("users")+".id").Related(&items, "Employees").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}
func (r *GeneratedCompanyResolver) EmployeesConnection(ctx context.Context, obj *Company, offset *int, limit *int, q *string, sort []*UserSortType, filter *UserFilterType) (res *UserResultType, err error) {
	f := &UserFilterType{
		Companies: &CompanyFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &UserFilterType{
			And: []*UserFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryUsersHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryUsers(ctx, r.GeneratedResolver, opts)
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
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := UserQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &UserResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("users")+".id = ?", *opts.ID)
	}

	var items []*User
	giOpts := GetItemsOptions{
		Alias:      TableName("users"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

type QueryUsersHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*UserSortType
	Filter *UserFilterType
}

func (r *GeneratedQueryResolver) Users(ctx context.Context, offset *int, limit *int, q *string, sort []*UserSortType, filter *UserFilterType) (*UserResultType, error) {
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
	query := UserQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &UserResultType{
		EntityResultType: EntityResultType{
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
	otps := GetItemsOptions{
		Alias:      TableName("users"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*User{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

func (r *GeneratedUserResultTypeResolver) Count(ctx context.Context, obj *UserResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("users"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &User{})
}

type GeneratedUserResolver struct{ *GeneratedResolver }

func (r *GeneratedUserResolver) Tasks(ctx context.Context, obj *User) (res []*Task, err error) {
	return r.Handlers.UserTasks(ctx, r.GeneratedResolver, obj)
}
func UserTasksHandler(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Task, err error) {

	items := []*Task{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "Tasks").Error
	res = items

	return
}

func (r *GeneratedUserResolver) TasksIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}

	items := []*Task{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Select(TableName("tasks")+".id").Related(&items, "Tasks").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}
func (r *GeneratedUserResolver) TasksConnection(ctx context.Context, obj *User, offset *int, limit *int, q *string, sort []*TaskSortType, filter *TaskFilterType) (res *TaskResultType, err error) {
	f := &TaskFilterType{
		Assignee: &UserFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &TaskFilterType{
			And: []*TaskFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryTasksHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryTasks(ctx, r.GeneratedResolver, opts)
}

func (r *GeneratedUserResolver) Companies(ctx context.Context, obj *User) (res []*Company, err error) {
	return r.Handlers.UserCompanies(ctx, r.GeneratedResolver, obj)
}
func UserCompaniesHandler(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Company, err error) {

	items := []*Company{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "Companies").Error
	res = items

	return
}

func (r *GeneratedUserResolver) CompaniesIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}

	items := []*Company{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Select(TableName("companies")+".id").Related(&items, "Companies").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}
func (r *GeneratedUserResolver) CompaniesConnection(ctx context.Context, obj *User, offset *int, limit *int, q *string, sort []*CompanySortType, filter *CompanyFilterType) (res *CompanyResultType, err error) {
	f := &CompanyFilterType{
		Employees: &UserFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &CompanyFilterType{
			And: []*CompanyFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryCompaniesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryCompanies(ctx, r.GeneratedResolver, opts)
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
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := TaskQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &TaskResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("tasks")+".id = ?", *opts.ID)
	}

	var items []*Task
	giOpts := GetItemsOptions{
		Alias:      TableName("tasks"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

type QueryTasksHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*TaskSortType
	Filter *TaskFilterType
}

func (r *GeneratedQueryResolver) Tasks(ctx context.Context, offset *int, limit *int, q *string, sort []*TaskSortType, filter *TaskFilterType) (*TaskResultType, error) {
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
	query := TaskQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &TaskResultType{
		EntityResultType: EntityResultType{
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
	otps := GetItemsOptions{
		Alias:      TableName("tasks"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*Task{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

func (r *GeneratedTaskResultTypeResolver) Count(ctx context.Context, obj *TaskResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("tasks"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &Task{})
}

type GeneratedTaskResolver struct{ *GeneratedResolver }

func (r *GeneratedTaskResolver) Assignee(ctx context.Context, obj *Task) (res *User, err error) {
	return r.Handlers.TaskAssignee(ctx, r.GeneratedResolver, obj)
}
func TaskAssigneeHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (res *User, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.AssigneeID != nil {
		item, _err := loaders["User"].Load(ctx, dataloader.StringKey(*obj.AssigneeID))()
		res, _ = item.(*User)

		err = _err
	}

	return
}
