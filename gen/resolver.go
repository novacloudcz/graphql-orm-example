package gen

import (
	"context"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/novacloudcz/graphql-orm/resolvers"
	uuid "github.com/satori/go.uuid"
)

func ToTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
		// Convert it by parsing
	}
}

func getPrincipalID(ctx context.Context) string {
	v, _ := ctx.Value(KeyPrincipalID).(string)
	return v
}

type Resolver struct {
	DB *DB
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) CompanyResultType() CompanyResultTypeResolver {
	return &companyResultTypeResolver{r}
}

func (r *Resolver) Company() CompanyResolver {
	return &companyResolver{r}
}

func (r *Resolver) UserResultType() UserResultTypeResolver {
	return &userResultTypeResolver{r}
}

func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

func (r *Resolver) TaskResultType() TaskResultTypeResolver {
	return &taskResultTypeResolver{r}
}

func (r *Resolver) Task() TaskResolver {
	return &taskResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateCompany(ctx context.Context, input map[string]interface{}) (item *Company, err error) {
	ID, ok := input["id"].(string)
	if !ok || ID == "" {
		ID = uuid.Must(uuid.NewV4()).String()
	}
	principalID := getPrincipalID(ctx)
	item = &Company{ID: ID, CreatedBy: principalID}
	tx := r.DB.db.Begin()

	if ids, ok := input["employeesIds"].([]interface{}); ok {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employees")
		association.Replace(items)
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc()),
		Result: item,
	})
	if err != nil {
		tx.Rollback()
		return
	}

	err = decoder.Decode(input)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *mutationResolver) UpdateCompany(ctx context.Context, id string, input map[string]interface{}) (item *Company, err error) {
	item = &Company{}
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	principalID := getPrincipalID(ctx)
	item.UpdatedBy = &principalID

	if ids, ok := input["employeesIds"].([]interface{}); ok {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employees")
		association.Replace(items)
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc()),
		Result: item,
	})
	if err != nil {
		tx.Rollback()
		return
	}

	err = decoder.Decode(input)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *mutationResolver) DeleteCompany(ctx context.Context, id string) (item *Company, err error) {
	item = &Company{}
	err = resolvers.GetItem(ctx, r.DB.Query(), item, &id)
	if err != nil {
		return
	}

	err = r.DB.Query().Delete(item, "id = ?", id).Error

	return
}

func (r *mutationResolver) CreateUser(ctx context.Context, input map[string]interface{}) (item *User, err error) {
	ID, ok := input["id"].(string)
	if !ok || ID == "" {
		ID = uuid.Must(uuid.NewV4()).String()
	}
	principalID := getPrincipalID(ctx)
	item = &User{ID: ID, CreatedBy: principalID}
	tx := r.DB.db.Begin()

	if ids, ok := input["tasksIds"].([]interface{}); ok {
		items := []Task{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Tasks")
		association.Replace(items)
	}

	if ids, ok := input["companiesIds"].([]interface{}); ok {
		items := []Company{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Companies")
		association.Replace(items)
	}

	if ids, ok := input["friendsIds"].([]interface{}); ok {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Friends")
		association.Replace(items)
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc()),
		Result: item,
	})
	if err != nil {
		tx.Rollback()
		return
	}

	err = decoder.Decode(input)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input map[string]interface{}) (item *User, err error) {
	item = &User{}
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	principalID := getPrincipalID(ctx)
	item.UpdatedBy = &principalID

	if ids, ok := input["tasksIds"].([]interface{}); ok {
		items := []Task{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Tasks")
		association.Replace(items)
	}

	if ids, ok := input["companiesIds"].([]interface{}); ok {
		items := []Company{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Companies")
		association.Replace(items)
	}

	if ids, ok := input["friendsIds"].([]interface{}); ok {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Friends")
		association.Replace(items)
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc()),
		Result: item,
	})
	if err != nil {
		tx.Rollback()
		return
	}

	err = decoder.Decode(input)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (item *User, err error) {
	item = &User{}
	err = resolvers.GetItem(ctx, r.DB.Query(), item, &id)
	if err != nil {
		return
	}

	err = r.DB.Query().Delete(item, "id = ?", id).Error

	return
}

func (r *mutationResolver) CreateTask(ctx context.Context, input map[string]interface{}) (item *Task, err error) {
	ID, ok := input["id"].(string)
	if !ok || ID == "" {
		ID = uuid.Must(uuid.NewV4()).String()
	}
	principalID := getPrincipalID(ctx)
	item = &Task{ID: ID, CreatedBy: principalID}
	tx := r.DB.db.Begin()

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc()),
		Result: item,
	})
	if err != nil {
		tx.Rollback()
		return
	}

	err = decoder.Decode(input)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *mutationResolver) UpdateTask(ctx context.Context, id string, input map[string]interface{}) (item *Task, err error) {
	item = &Task{}
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	principalID := getPrincipalID(ctx)
	item.UpdatedBy = &principalID

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc()),
		Result: item,
	})
	if err != nil {
		tx.Rollback()
		return
	}

	err = decoder.Decode(input)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (item *Task, err error) {
	item = &Task{}
	err = resolvers.GetItem(ctx, r.DB.Query(), item, &id)
	if err != nil {
		return
	}

	err = r.DB.Query().Delete(item, "id = ?", id).Error

	return
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Company(ctx context.Context, id *string, q *string) (*Company, error) {
	t := Company{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *queryResolver) Companies(ctx context.Context, offset *int, limit *int, q *string, sort []CompanySortType, filter *CompanyFilterType) (*CompanyResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range sort {
		_sort = append(_sort, s)
	}
	query := CompanyQueryFilter{q}
	return &CompanyResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: offset,
			Limit:  limit,
			Query:  &query,
			Sort:   _sort,
			Filter: filter,
		},
	}, nil
}

type companyResultTypeResolver struct{ *Resolver }

func (r *companyResultTypeResolver) Items(ctx context.Context, obj *CompanyResultType) (items []*Company, err error) {
	err = obj.GetItems(ctx, r.DB.db, "companies", &items)
	return
}

func (r *companyResultTypeResolver) Count(ctx context.Context, obj *CompanyResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Company{})
}

type companyResolver struct{ *Resolver }

func (r *companyResolver) Employees(ctx context.Context, obj *Company) (res []*User, err error) {

	items := []*User{}
	err = r.DB.Query().Model(obj).Related(&items, "Employees").Error
	res = items

	return
}

func (r *queryResolver) User(ctx context.Context, id *string, q *string) (*User, error) {
	t := User{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *queryResolver) Users(ctx context.Context, offset *int, limit *int, q *string, sort []UserSortType, filter *UserFilterType) (*UserResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range sort {
		_sort = append(_sort, s)
	}
	query := UserQueryFilter{q}
	return &UserResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: offset,
			Limit:  limit,
			Query:  &query,
			Sort:   _sort,
			Filter: filter,
		},
	}, nil
}

type userResultTypeResolver struct{ *Resolver }

func (r *userResultTypeResolver) Items(ctx context.Context, obj *UserResultType) (items []*User, err error) {
	err = obj.GetItems(ctx, r.DB.db, "users", &items)
	return
}

func (r *userResultTypeResolver) Count(ctx context.Context, obj *UserResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &User{})
}

type userResolver struct{ *Resolver }

func (r *userResolver) Tasks(ctx context.Context, obj *User) (res []*Task, err error) {

	items := []*Task{}
	err = r.DB.Query().Model(obj).Related(&items, "Tasks").Error
	res = items

	return
}

func (r *userResolver) Companies(ctx context.Context, obj *User) (res []*Company, err error) {

	items := []*Company{}
	err = r.DB.Query().Model(obj).Related(&items, "Companies").Error
	res = items

	return
}

func (r *userResolver) Friends(ctx context.Context, obj *User) (res []*User, err error) {

	items := []*User{}
	err = r.DB.Query().Model(obj).Related(&items, "Friends").Error
	res = items

	return
}

func (r *queryResolver) Task(ctx context.Context, id *string, q *string) (*Task, error) {
	t := Task{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *queryResolver) Tasks(ctx context.Context, offset *int, limit *int, q *string, sort []TaskSortType, filter *TaskFilterType) (*TaskResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range sort {
		_sort = append(_sort, s)
	}
	query := TaskQueryFilter{q}
	return &TaskResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: offset,
			Limit:  limit,
			Query:  &query,
			Sort:   _sort,
			Filter: filter,
		},
	}, nil
}

type taskResultTypeResolver struct{ *Resolver }

func (r *taskResultTypeResolver) Items(ctx context.Context, obj *TaskResultType) (items []*Task, err error) {
	err = obj.GetItems(ctx, r.DB.db, "tasks", &items)
	return
}

func (r *taskResultTypeResolver) Count(ctx context.Context, obj *TaskResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Task{})
}

type taskResolver struct{ *Resolver }

func (r *taskResolver) Assignee(ctx context.Context, obj *Task) (res *User, err error) {

	item := User{}
	_res := r.DB.Query().Model(obj).Related(&item, "Assignee")
	if _res.RecordNotFound() {
		return
	} else {
		err = _res.Error
	}
	res = &item

	return
}
