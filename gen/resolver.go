package gen

import (
	"context"
	"time"

	"github.com/novacloudcz/graphql-orm/events"
	"github.com/novacloudcz/graphql-orm/resolvers"
	uuid "github.com/satori/go.uuid"
)

func getPrincipalID(ctx context.Context) *string {
	v, _ := ctx.Value(KeyPrincipalID).(*string)
	return v
}

type GeneratedResolver struct {
	DB              *DB
	EventController *events.EventController
}

func (r *GeneratedResolver) Mutation() MutationResolver {
	return &GeneratedMutationResolver{r}
}
func (r *GeneratedResolver) Query() QueryResolver {
	return &GeneratedQueryResolver{r}
}

func (r *GeneratedResolver) CompanyResultType() CompanyResultTypeResolver {
	return &GeneratedCompanyResultTypeResolver{r}
}

func (r *GeneratedResolver) Company() CompanyResolver {
	return &GeneratedCompanyResolver{r}
}

func (r *GeneratedResolver) UserResultType() UserResultTypeResolver {
	return &GeneratedUserResultTypeResolver{r}
}

func (r *GeneratedResolver) User() UserResolver {
	return &GeneratedUserResolver{r}
}

func (r *GeneratedResolver) TaskResultType() TaskResultTypeResolver {
	return &GeneratedTaskResultTypeResolver{r}
}

func (r *GeneratedResolver) Task() TaskResolver {
	return &GeneratedTaskResolver{r}
}

type GeneratedMutationResolver struct{ *GeneratedResolver }

func (r *GeneratedMutationResolver) CreateCompany(ctx context.Context, input map[string]interface{}) (item *Company, err error) {
	principalID := getPrincipalID(ctx)
	now := time.Now()
	item = &Company{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Company",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	if val, ok := input["id"].(string); ok && (item.ID != val) {
		item.ID = val
		event.AddNewValue("id", &val)
	}

	if val, ok := input["name"].(string); ok && (item.Name == nil || *item.Name != val) {
		item.Name = &val
		event.AddNewValue("name", &val)
	}

	if ids, ok := input["employeesIds"].([]interface{}); ok {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employees")
		association.Replace(items)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = r.EventController.SendEvent(ctx, &event)

	return
}
func (r *GeneratedMutationResolver) UpdateCompany(ctx context.Context, id string, input map[string]interface{}) (item *Company, err error) {
	item = &Company{}
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	principalID := getPrincipalID(ctx)
	item.UpdatedBy = principalID

	if val, ok := input["name"].(string); ok && (item.Name == nil || *item.Name != val) {
		item.Name = &val
	}

	if ids, ok := input["employeesIds"].([]interface{}); ok {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employees")
		association.Replace(items)
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *GeneratedMutationResolver) DeleteCompany(ctx context.Context, id string) (item *Company, err error) {
	item = &Company{}
	err = resolvers.GetItem(ctx, r.DB.Query(), item, &id)
	if err != nil {
		return
	}

	err = r.DB.Query().Delete(item, "id = ?", id).Error

	return
}

func (r *GeneratedMutationResolver) CreateUser(ctx context.Context, input map[string]interface{}) (item *User, err error) {
	principalID := getPrincipalID(ctx)
	now := time.Now()
	item = &User{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "User",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	if val, ok := input["id"].(string); ok && (item.ID != val) {
		item.ID = val
		event.AddNewValue("id", &val)
	}

	if val, ok := input["email"].(string); ok && (item.Email == nil || *item.Email != val) {
		item.Email = &val
		event.AddNewValue("email", &val)
	}

	if val, ok := input["firstName"].(string); ok && (item.FirstName == nil || *item.FirstName != val) {
		item.FirstName = &val
		event.AddNewValue("firstName", &val)
	}

	if val, ok := input["lastName"].(string); ok && (item.LastName == nil || *item.LastName != val) {
		item.LastName = &val
		event.AddNewValue("lastName", &val)
	}

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

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = r.EventController.SendEvent(ctx, &event)

	return
}
func (r *GeneratedMutationResolver) UpdateUser(ctx context.Context, id string, input map[string]interface{}) (item *User, err error) {
	item = &User{}
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	principalID := getPrincipalID(ctx)
	item.UpdatedBy = principalID

	if val, ok := input["email"].(string); ok && (item.Email == nil || *item.Email != val) {
		item.Email = &val
	}

	if val, ok := input["firstName"].(string); ok && (item.FirstName == nil || *item.FirstName != val) {
		item.FirstName = &val
	}

	if val, ok := input["lastName"].(string); ok && (item.LastName == nil || *item.LastName != val) {
		item.LastName = &val
	}

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

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *GeneratedMutationResolver) DeleteUser(ctx context.Context, id string) (item *User, err error) {
	item = &User{}
	err = resolvers.GetItem(ctx, r.DB.Query(), item, &id)
	if err != nil {
		return
	}

	err = r.DB.Query().Delete(item, "id = ?", id).Error

	return
}

func (r *GeneratedMutationResolver) CreateTask(ctx context.Context, input map[string]interface{}) (item *Task, err error) {
	principalID := getPrincipalID(ctx)
	now := time.Now()
	item = &Task{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Task",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	if val, ok := input["id"].(string); ok && (item.ID != val) {
		item.ID = val
		event.AddNewValue("id", &val)
	}

	if val, ok := input["title"].(string); ok && (item.Title == nil || *item.Title != val) {
		item.Title = &val
		event.AddNewValue("title", &val)
	}

	if val, ok := input["completed"].(bool); ok && (item.Completed == nil || *item.Completed != val) {
		item.Completed = &val
		event.AddNewValue("completed", &val)
	}

	if val, ok := input["dueDate"].(time.Time); ok && (item.DueDate == nil || *item.DueDate != val) {
		item.DueDate = &val
		event.AddNewValue("dueDate", &val)
	}

	if val, ok := input["type"].(TaskType); ok && (item.Type == nil || *item.Type != val) {
		item.Type = &val
		event.AddNewValue("type", &val)
	}

	if val, ok := input["assigneeId"].(string); ok && (item.AssigneeID == nil || *item.AssigneeID != val) {
		item.AssigneeID = &val
		event.AddNewValue("assigneeId", &val)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = r.EventController.SendEvent(ctx, &event)

	return
}
func (r *GeneratedMutationResolver) UpdateTask(ctx context.Context, id string, input map[string]interface{}) (item *Task, err error) {
	item = &Task{}
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	principalID := getPrincipalID(ctx)
	item.UpdatedBy = principalID

	if val, ok := input["title"].(string); ok && (item.Title == nil || *item.Title != val) {
		item.Title = &val
	}

	if val, ok := input["completed"].(bool); ok && (item.Completed == nil || *item.Completed != val) {
		item.Completed = &val
	}

	if val, ok := input["dueDate"].(time.Time); ok && (item.DueDate == nil || *item.DueDate != val) {
		item.DueDate = &val
	}

	if val, ok := input["type"].(TaskType); ok && (item.Type == nil || *item.Type != val) {
		item.Type = &val
	}

	if val, ok := input["assigneeId"].(string); ok && (item.AssigneeID == nil || *item.AssigneeID != val) {
		item.AssigneeID = &val
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *GeneratedMutationResolver) DeleteTask(ctx context.Context, id string) (item *Task, err error) {
	item = &Task{}
	err = resolvers.GetItem(ctx, r.DB.Query(), item, &id)
	if err != nil {
		return
	}

	err = r.DB.Query().Delete(item, "id = ?", id).Error

	return
}

type GeneratedQueryResolver struct{ *GeneratedResolver }

func (r *GeneratedQueryResolver) Company(ctx context.Context, id *string, q *string) (*Company, error) {
	t := Company{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *GeneratedQueryResolver) Companies(ctx context.Context, offset *int, limit *int, q *string, sort []CompanySortType, filter *CompanyFilterType) (*CompanyResultType, error) {
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

	items := []*User{}
	err = r.DB.Query().Model(obj).Related(&items, "Employees").Error
	res = items

	return
}

func (r *GeneratedQueryResolver) User(ctx context.Context, id *string, q *string) (*User, error) {
	t := User{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *GeneratedQueryResolver) Users(ctx context.Context, offset *int, limit *int, q *string, sort []UserSortType, filter *UserFilterType) (*UserResultType, error) {
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

	items := []*Task{}
	err = r.DB.Query().Model(obj).Related(&items, "Tasks").Error
	res = items

	return
}

func (r *GeneratedUserResolver) Companies(ctx context.Context, obj *User) (res []*Company, err error) {

	items := []*Company{}
	err = r.DB.Query().Model(obj).Related(&items, "Companies").Error
	res = items

	return
}

func (r *GeneratedUserResolver) Friends(ctx context.Context, obj *User) (res []*User, err error) {

	items := []*User{}
	err = r.DB.Query().Model(obj).Related(&items, "Friends").Error
	res = items

	return
}

func (r *GeneratedQueryResolver) Task(ctx context.Context, id *string, q *string) (*Task, error) {
	t := Task{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *GeneratedQueryResolver) Tasks(ctx context.Context, offset *int, limit *int, q *string, sort []TaskSortType, filter *TaskFilterType) (*TaskResultType, error) {
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
