package gen

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/novacloudcz/graphql-orm/events"
	"github.com/novacloudcz/graphql-orm/resolvers"
	uuid "github.com/satori/go.uuid"
	"github.com/vektah/gqlparser/ast"
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

	var changes CompanyChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name
		event.AddNewValue("name", changes.Name)
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

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateCompany(ctx context.Context, id string, input map[string]interface{}) (item *Company, err error) {
	principalID := getPrincipalID(ctx)
	item = &Company{}
	now := time.Now()
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Company",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes CompanyChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
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
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
		// data, _ := json.Marshal(event)
		// fmt.Println("?",string(data))
	}

	return
}
func (r *GeneratedMutationResolver) DeleteCompany(ctx context.Context, id string) (item *Company, err error) {
	principalID := getPrincipalID(ctx)
	item = &Company{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Company",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, "companies.id = ?", id).Error
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

	var changes UserChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["email"]; ok && (item.Email != changes.Email) && (item.Email == nil || changes.Email == nil || *item.Email != *changes.Email) {
		item.Email = changes.Email
		event.AddNewValue("email", changes.Email)
	}

	if _, ok := input["firstName"]; ok && (item.FirstName != changes.FirstName) && (item.FirstName == nil || changes.FirstName == nil || *item.FirstName != *changes.FirstName) {
		item.FirstName = changes.FirstName
		event.AddNewValue("firstName", changes.FirstName)
	}

	if _, ok := input["lastName"]; ok && (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {
		item.LastName = changes.LastName
		event.AddNewValue("lastName", changes.LastName)
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

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateUser(ctx context.Context, id string, input map[string]interface{}) (item *User, err error) {
	principalID := getPrincipalID(ctx)
	item = &User{}
	now := time.Now()
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "User",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes UserChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["email"]; ok && (item.Email != changes.Email) && (item.Email == nil || changes.Email == nil || *item.Email != *changes.Email) {
		event.AddOldValue("email", item.Email)
		event.AddNewValue("email", changes.Email)
		item.Email = changes.Email
	}

	if _, ok := input["firstName"]; ok && (item.FirstName != changes.FirstName) && (item.FirstName == nil || changes.FirstName == nil || *item.FirstName != *changes.FirstName) {
		event.AddOldValue("firstName", item.FirstName)
		event.AddNewValue("firstName", changes.FirstName)
		item.FirstName = changes.FirstName
	}

	if _, ok := input["lastName"]; ok && (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {
		event.AddOldValue("lastName", item.LastName)
		event.AddNewValue("lastName", changes.LastName)
		item.LastName = changes.LastName
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
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
		// data, _ := json.Marshal(event)
		// fmt.Println("?",string(data))
	}

	return
}
func (r *GeneratedMutationResolver) DeleteUser(ctx context.Context, id string) (item *User, err error) {
	principalID := getPrincipalID(ctx)
	item = &User{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "User",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, "users.id = ?", id).Error
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

	var changes TaskChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["title"]; ok && (item.Title != changes.Title) && (item.Title == nil || changes.Title == nil || *item.Title != *changes.Title) {
		item.Title = changes.Title
		event.AddNewValue("title", changes.Title)
	}

	if _, ok := input["completed"]; ok && (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {
		item.Completed = changes.Completed
		event.AddNewValue("completed", changes.Completed)
	}

	if _, ok := input["dueDate"]; ok && (item.DueDate != changes.DueDate) && (item.DueDate == nil || changes.DueDate == nil || *item.DueDate != *changes.DueDate) {
		item.DueDate = changes.DueDate
		event.AddNewValue("dueDate", changes.DueDate)
	}

	if _, ok := input["type"]; ok && (item.Type != changes.Type) && (item.Type == nil || changes.Type == nil || *item.Type != *changes.Type) {
		item.Type = changes.Type
		event.AddNewValue("type", changes.Type)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description
		event.AddNewValue("description", changes.Description)
	}

	if _, ok := input["assigneeId"]; ok && (item.AssigneeID != changes.AssigneeID) && (item.AssigneeID == nil || changes.AssigneeID == nil || *item.AssigneeID != *changes.AssigneeID) {
		item.AssigneeID = changes.AssigneeID
		event.AddNewValue("assigneeId", changes.AssigneeID)
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

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateTask(ctx context.Context, id string, input map[string]interface{}) (item *Task, err error) {
	principalID := getPrincipalID(ctx)
	item = &Task{}
	now := time.Now()
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Task",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes TaskChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["title"]; ok && (item.Title != changes.Title) && (item.Title == nil || changes.Title == nil || *item.Title != *changes.Title) {
		event.AddOldValue("title", item.Title)
		event.AddNewValue("title", changes.Title)
		item.Title = changes.Title
	}

	if _, ok := input["completed"]; ok && (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {
		event.AddOldValue("completed", item.Completed)
		event.AddNewValue("completed", changes.Completed)
		item.Completed = changes.Completed
	}

	if _, ok := input["dueDate"]; ok && (item.DueDate != changes.DueDate) && (item.DueDate == nil || changes.DueDate == nil || *item.DueDate != *changes.DueDate) {
		event.AddOldValue("dueDate", item.DueDate)
		event.AddNewValue("dueDate", changes.DueDate)
		item.DueDate = changes.DueDate
	}

	if _, ok := input["type"]; ok && (item.Type != changes.Type) && (item.Type == nil || changes.Type == nil || *item.Type != *changes.Type) {
		event.AddOldValue("type", item.Type)
		event.AddNewValue("type", changes.Type)
		item.Type = changes.Type
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	if _, ok := input["assigneeId"]; ok && (item.AssigneeID != changes.AssigneeID) && (item.AssigneeID == nil || changes.AssigneeID == nil || *item.AssigneeID != *changes.AssigneeID) {
		event.AddOldValue("assigneeId", item.AssigneeID)
		event.AddNewValue("assigneeId", changes.AssigneeID)
		item.AssigneeID = changes.AssigneeID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
		// data, _ := json.Marshal(event)
		// fmt.Println("?",string(data))
	}

	return
}
func (r *GeneratedMutationResolver) DeleteTask(ctx context.Context, id string) (item *Task, err error) {
	principalID := getPrincipalID(ctx)
	item = &Task{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Task",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, "tasks.id = ?", id).Error
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

type GeneratedQueryResolver struct{ *GeneratedResolver }

func (r *GeneratedQueryResolver) Company(ctx context.Context, id *string, q *string, filter *CompanyFilterType) (*Company, error) {
	query := CompanyQueryFilter{q}
	offset := 0
	limit := 1
	rt := &CompanyResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: filter,
		},
	}
	qb := r.DB.Query()
	if id != nil {
		qb = qb.Where("companies.id = ?", *id)
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
func (r *GeneratedQueryResolver) Companies(ctx context.Context, offset *int, limit *int, q *string, sort []CompanySortType, filter *CompanyFilterType) (*CompanyResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range sort {
		_sort = append(_sort, s)
	}
	query := CompanyQueryFilter{q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	return &CompanyResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset:       offset,
			Limit:        limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       filter,
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

func (r *GeneratedQueryResolver) User(ctx context.Context, id *string, q *string, filter *UserFilterType) (*User, error) {
	query := UserQueryFilter{q}
	offset := 0
	limit := 1
	rt := &UserResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: filter,
		},
	}
	qb := r.DB.Query()
	if id != nil {
		qb = qb.Where("users.id = ?", *id)
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
func (r *GeneratedQueryResolver) Users(ctx context.Context, offset *int, limit *int, q *string, sort []UserSortType, filter *UserFilterType) (*UserResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range sort {
		_sort = append(_sort, s)
	}
	query := UserQueryFilter{q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	return &UserResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset:       offset,
			Limit:        limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       filter,
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

func (r *GeneratedQueryResolver) Task(ctx context.Context, id *string, q *string, filter *TaskFilterType) (*Task, error) {
	query := TaskQueryFilter{q}
	offset := 0
	limit := 1
	rt := &TaskResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: filter,
		},
	}
	qb := r.DB.Query()
	if id != nil {
		qb = qb.Where("tasks.id = ?", *id)
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
func (r *GeneratedQueryResolver) Tasks(ctx context.Context, offset *int, limit *int, q *string, sort []TaskSortType, filter *TaskFilterType) (*TaskResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range sort {
		_sort = append(_sort, s)
	}
	query := TaskQueryFilter{q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	return &TaskResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset:       offset,
			Limit:        limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       filter,
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
