package gen

import (
	"context"

	"github.com/novacloudcz/graphql-orm/resolvers"
	uuid "github.com/satori/go.uuid"
)

type Resolver struct {
	DB *DB
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) TodoResultType() TodoResultTypeResolver {
	return &todoResultTypeResolver{r}
}

func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

func (r *Resolver) UserResultType() UserResultTypeResolver {
	return &userResultTypeResolver{r}
}

func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

func (r *Resolver) CompanyResultType() CompanyResultTypeResolver {
	return &companyResultTypeResolver{r}
}

func (r *Resolver) Company() CompanyResolver {
	return &companyResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input map[string]interface{}) (item *Todo, err error) {
	item = &Todo{ID: uuid.Must(uuid.NewV4()).String()}
	tx := r.DB.db.Begin()

	err = resolvers.CreateItem(ctx, tx, item, input)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, input map[string]interface{}) (item *Todo, err error) {
	item = &Todo{}
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	err = resolvers.UpdateItem(ctx, tx, item, input)

	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (item *Todo, err error) {
	item = &Todo{}
	err = resolvers.GetItem(ctx, r.DB.Query(), item, &id)
	if err != nil {
		return
	}

	err = resolvers.DeleteItem(ctx, r.DB.Query(), item, id)

	return
}

func (r *mutationResolver) CreateUser(ctx context.Context, input map[string]interface{}) (item *User, err error) {
	item = &User{ID: uuid.Must(uuid.NewV4()).String()}
	tx := r.DB.db.Begin()

	if ids, ok := input["todosIds"].([]interface{}); ok {
		items := []Todo{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Todos")
		association.Replace(items)
	}

	if ids, ok := input["friendsIds"].([]interface{}); ok {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Friends")
		association.Replace(items)
	}

	if ids, ok := input["employersIds"].([]interface{}); ok {
		items := []Company{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employers")
		association.Replace(items)
	}

	err = resolvers.CreateItem(ctx, tx, item, input)
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

	if ids, ok := input["todosIds"].([]interface{}); ok {
		items := []Todo{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Todos")
		association.Replace(items)
	}

	if ids, ok := input["friendsIds"].([]interface{}); ok {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Friends")
		association.Replace(items)
	}

	if ids, ok := input["employersIds"].([]interface{}); ok {
		items := []Company{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employers")
		association.Replace(items)
	}

	err = resolvers.UpdateItem(ctx, tx, item, input)

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

	err = resolvers.DeleteItem(ctx, r.DB.Query(), item, id)

	return
}

func (r *mutationResolver) CreateCompany(ctx context.Context, input map[string]interface{}) (item *Company, err error) {
	item = &Company{ID: uuid.Must(uuid.NewV4()).String()}
	tx := r.DB.db.Begin()

	if ids, ok := input["employeesIds"].([]interface{}); ok {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employees")
		association.Replace(items)
	}

	err = resolvers.CreateItem(ctx, tx, item, input)
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

	if ids, ok := input["employeesIds"].([]interface{}); ok {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employees")
		association.Replace(items)
	}

	err = resolvers.UpdateItem(ctx, tx, item, input)

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

	err = resolvers.DeleteItem(ctx, r.DB.Query(), item, id)

	return
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todo(ctx context.Context, id *string, q *string) (*Todo, error) {
	t := Todo{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *queryResolver) Todos(ctx context.Context, offset *int, limit *int, q *string) (*TodoResultType, error) {
	return &TodoResultType{}, nil
}

type todoResultTypeResolver struct{ *Resolver }

func (r *todoResultTypeResolver) Items(ctx context.Context, obj *TodoResultType) (items []*Todo, err error) {
	err = resolvers.GetResultTypeItems(ctx, r.DB.db, &items)
	return
}

func (r *todoResultTypeResolver) Count(ctx context.Context, obj *TodoResultType) (count int, err error) {
	return resolvers.GetResultTypeCount(ctx, r.DB.db, &Todo{})
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *Todo) (res *User, err error) {

	item := User{}
	_res := r.DB.Query().Model(obj).Related(&item, "User")
	if _res.RecordNotFound() {
		return
	} else {
		err = _res.Error
	}
	res = &item

	return
}

func (r *queryResolver) User(ctx context.Context, id *string, q *string) (*User, error) {
	t := User{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *queryResolver) Users(ctx context.Context, offset *int, limit *int, q *string) (*UserResultType, error) {
	return &UserResultType{}, nil
}

type userResultTypeResolver struct{ *Resolver }

func (r *userResultTypeResolver) Items(ctx context.Context, obj *UserResultType) (items []*User, err error) {
	err = resolvers.GetResultTypeItems(ctx, r.DB.db, &items)
	return
}

func (r *userResultTypeResolver) Count(ctx context.Context, obj *UserResultType) (count int, err error) {
	return resolvers.GetResultTypeCount(ctx, r.DB.db, &User{})
}

type userResolver struct{ *Resolver }

func (r *userResolver) Todos(ctx context.Context, obj *User) (res []*Todo, err error) {

	items := []*Todo{}
	err = r.DB.Query().Model(obj).Related(&items, "Todos").Error
	res = items

	return
}

func (r *userResolver) Friends(ctx context.Context, obj *User) (res []*User, err error) {

	items := []*User{}
	err = r.DB.Query().Model(obj).Related(&items, "Friends").Error
	res = items

	return
}

func (r *userResolver) Employers(ctx context.Context, obj *User) (res []*Company, err error) {

	items := []*Company{}
	err = r.DB.Query().Model(obj).Related(&items, "Employers").Error
	res = items

	return
}

func (r *queryResolver) Company(ctx context.Context, id *string, q *string) (*Company, error) {
	t := Company{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *queryResolver) Companies(ctx context.Context, offset *int, limit *int, q *string) (*CompanyResultType, error) {
	return &CompanyResultType{}, nil
}

type companyResultTypeResolver struct{ *Resolver }

func (r *companyResultTypeResolver) Items(ctx context.Context, obj *CompanyResultType) (items []*Company, err error) {
	err = resolvers.GetResultTypeItems(ctx, r.DB.db, &items)
	return
}

func (r *companyResultTypeResolver) Count(ctx context.Context, obj *CompanyResultType) (count int, err error) {
	return resolvers.GetResultTypeCount(ctx, r.DB.db, &Company{})
}

type companyResolver struct{ *Resolver }

func (r *companyResolver) Employees(ctx context.Context, obj *Company) (res []*User, err error) {

	items := []*User{}
	err = r.DB.Query().Model(obj).Related(&items, "Employees").Error
	res = items

	return
}
