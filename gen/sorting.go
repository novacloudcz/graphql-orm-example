package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

func (s CompanySortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("companies"), sorts, joins)
}
func (s CompanySortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+"id "+s.ID.String())
	}

	if s.Name != nil {
		*sorts = append(*sorts, aliasPrefix+"name "+s.Name.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedAt "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"createdAt "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedBy "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"createdBy "+s.CreatedBy.String())
	}

	if s.Employees != nil {
		_alias := alias + "_employees"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("company_employees"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("company_id")+" LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employee_id")+" = "+dialect.Quote(_alias)+".id")
		err := s.Employees.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s UserSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("users"), sorts, joins)
}
func (s UserSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+"id "+s.ID.String())
	}

	if s.Email != nil {
		*sorts = append(*sorts, aliasPrefix+"email "+s.Email.String())
	}

	if s.FirstName != nil {
		*sorts = append(*sorts, aliasPrefix+"firstName "+s.FirstName.String())
	}

	if s.LastName != nil {
		*sorts = append(*sorts, aliasPrefix+"lastName "+s.LastName.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedAt "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"createdAt "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedBy "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"createdBy "+s.CreatedBy.String())
	}

	if s.Tasks != nil {
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("tasks"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("assigneeId")+" = "+dialect.Quote(alias)+".id")
		err := s.Tasks.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Companies != nil {
		_alias := alias + "_companies"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("company_employees"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employee_id")+" LEFT JOIN "+dialect.Quote(TableName("companies"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("company_id")+" = "+dialect.Quote(_alias)+".id")
		err := s.Companies.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Friends != nil {
		_alias := alias + "_friends"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("user_friends"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("friend_id")+" LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("friend_id")+" = "+dialect.Quote(_alias)+".id")
		err := s.Friends.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s TaskSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("tasks"), sorts, joins)
}
func (s TaskSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+"id "+s.ID.String())
	}

	if s.Title != nil {
		*sorts = append(*sorts, aliasPrefix+"title "+s.Title.String())
	}

	if s.Completed != nil {
		*sorts = append(*sorts, aliasPrefix+"completed "+s.Completed.String())
	}

	if s.DueDate != nil {
		*sorts = append(*sorts, aliasPrefix+"dueDate "+s.DueDate.String())
	}

	if s.Description != nil {
		*sorts = append(*sorts, aliasPrefix+"description "+s.Description.String())
	}

	if s.AssigneeID != nil {
		*sorts = append(*sorts, aliasPrefix+"assigneeId "+s.AssigneeID.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedAt "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"createdAt "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedBy "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"createdBy "+s.CreatedBy.String())
	}

	if s.Assignee != nil {
		_alias := alias + "_assignee"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("assigneeId"))
		err := s.Assignee.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
