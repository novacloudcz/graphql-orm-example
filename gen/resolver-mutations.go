package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/novacloudcz/graphql-orm/events"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

func (r *GeneratedMutationResolver) CreateCompany(ctx context.Context, input map[string]interface{}) (item *Company, err error) {
	return r.Handlers.CreateCompany(ctx, r.GeneratedResolver, input)
}
func CreateCompanyHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Company, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
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

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, ok := input["employeesIds"].([]interface{}); ok {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employees")
		association.Replace(items)
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
	return r.Handlers.UpdateCompany(ctx, r.GeneratedResolver, id, input)
}
func UpdateCompanyHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Company, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
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

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, ok := input["employeesIds"].([]interface{}); ok {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employees")
		association.Replace(items)
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
	return r.Handlers.DeleteCompany(ctx, r.GeneratedResolver, id)
}
func DeleteCompanyHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Company, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Company{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = GetItem(ctx, tx, item, &id)
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

	err = tx.Delete(item, TableName("companies")+".id = ?", id).Error
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
func (r *GeneratedMutationResolver) DeleteAllCompanies(ctx context.Context) (bool, error) {
	return r.Handlers.DeleteAllCompanies(ctx, r.GeneratedResolver)
}
func DeleteAllCompaniesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	err := r.DB.db.Delete(&Company{}).Error
	return err == nil, err
}

func (r *GeneratedMutationResolver) CreateUser(ctx context.Context, input map[string]interface{}) (item *User, err error) {
	return r.Handlers.CreateUser(ctx, r.GeneratedResolver, input)
}
func CreateUserHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *User, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
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

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
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
	return r.Handlers.UpdateUser(ctx, r.GeneratedResolver, id, input)
}
func UpdateUserHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *User, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
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

	err = GetItem(ctx, tx, item, &id)
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

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
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
	return r.Handlers.DeleteUser(ctx, r.GeneratedResolver, id)
}
func DeleteUserHandler(ctx context.Context, r *GeneratedResolver, id string) (item *User, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &User{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = GetItem(ctx, tx, item, &id)
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

	err = tx.Delete(item, TableName("users")+".id = ?", id).Error
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
func (r *GeneratedMutationResolver) DeleteAllUsers(ctx context.Context) (bool, error) {
	return r.Handlers.DeleteAllUsers(ctx, r.GeneratedResolver)
}
func DeleteAllUsersHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	err := r.DB.db.Delete(&User{}).Error
	return err == nil, err
}

func (r *GeneratedMutationResolver) CreateTask(ctx context.Context, input map[string]interface{}) (item *Task, err error) {
	return r.Handlers.CreateTask(ctx, r.GeneratedResolver, input)
}
func CreateTaskHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Task, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
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
	return r.Handlers.UpdateTask(ctx, r.GeneratedResolver, id, input)
}
func UpdateTaskHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Task, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
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

	err = GetItem(ctx, tx, item, &id)
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
	return r.Handlers.DeleteTask(ctx, r.GeneratedResolver, id)
}
func DeleteTaskHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Task, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Task{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = GetItem(ctx, tx, item, &id)
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

	err = tx.Delete(item, TableName("tasks")+".id = ?", id).Error
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
func (r *GeneratedMutationResolver) DeleteAllTasks(ctx context.Context) (bool, error) {
	return r.Handlers.DeleteAllTasks(ctx, r.GeneratedResolver)
}
func DeleteAllTasksHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	err := r.DB.db.Delete(&Task{}).Error
	return err == nil, err
}
