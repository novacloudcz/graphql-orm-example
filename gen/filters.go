package gen

import (
	"context"
	"fmt"
	"strings"
)

func (f *CompanyFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, "companies", wheres, values, joins)
}
func (f *CompanyFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := alias + "."

	_where, _values := f.WhereContent(aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Employees != nil {
		_alias := alias + "_employees"
		*joins = append(*joins, "LEFT JOIN company_employees "+_alias+"_jointable ON "+alias+".id = "+_alias+"_jointable.company_id LEFT JOIN users "+_alias+" ON "+_alias+"_jointable.employee_id = "+_alias+".id")
		err := f.Employees.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *CompanyFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		for _, or := range f.Or {
			_cond, _values := or.WhereContent(aliasPrefix)
			cs = append(cs, _cond...)
			vs = append(vs, _values...)
		}
		conditions = append(conditions, "("+strings.Join(cs, " OR ")+")")
		values = append(values, vs...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		for _, or := range f.Or {
			_cond, _values := or.WhereContent(aliasPrefix)
			cs = append(cs, _cond...)
			vs = append(vs, _values...)
		}
		conditions = append(conditions, strings.Join(cs, " AND "))
		values = append(values, vs...)
	}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+"id = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+"id != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+"id > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+"id < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+"id >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+"id <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+"id IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+"name = ?")
		values = append(values, f.Name)
	}
	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+"name != ?")
		values = append(values, f.NameNe)
	}
	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+"name > ?")
		values = append(values, f.NameGt)
	}
	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+"name < ?")
		values = append(values, f.NameLt)
	}
	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+"name >= ?")
		values = append(values, f.NameGte)
	}
	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+"name <= ?")
		values = append(values, f.NameLte)
	}
	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+"name IN (?)")
		values = append(values, f.NameIn)
	}
	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+"name LIKE ?")
		values = append(values, strings.ReplaceAll(strings.ReplaceAll(*f.NameLike, "?", "_"), "*", "%"))
	}
	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+"name LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}
	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+"name LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+"createdAt != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+"createdAt >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+"createdAt <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+"createdAt IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	return
}

func (f *UserFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, "users", wheres, values, joins)
}
func (f *UserFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := alias + "."

	_where, _values := f.WhereContent(aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Tasks != nil {
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN tasks "+_alias+" ON "+_alias+".assigneeId = "+alias+".id")
		err := f.Tasks.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Companies != nil {
		_alias := alias + "_companies"
		*joins = append(*joins, "LEFT JOIN company_employees "+_alias+"_jointable ON "+alias+".id = "+_alias+"_jointable.employee_id LEFT JOIN companies "+_alias+" ON "+_alias+"_jointable.company_id = "+_alias+".id")
		err := f.Companies.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Friends != nil {
		_alias := alias + "_friends"
		*joins = append(*joins, "LEFT JOIN user_friends "+_alias+"_jointable ON "+alias+".id = "+_alias+"_jointable.friend_id LEFT JOIN users "+_alias+" ON "+_alias+"_jointable.friend_id = "+_alias+".id")
		err := f.Friends.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *UserFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		for _, or := range f.Or {
			_cond, _values := or.WhereContent(aliasPrefix)
			cs = append(cs, _cond...)
			vs = append(vs, _values...)
		}
		conditions = append(conditions, "("+strings.Join(cs, " OR ")+")")
		values = append(values, vs...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		for _, or := range f.Or {
			_cond, _values := or.WhereContent(aliasPrefix)
			cs = append(cs, _cond...)
			vs = append(vs, _values...)
		}
		conditions = append(conditions, strings.Join(cs, " AND "))
		values = append(values, vs...)
	}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+"id = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+"id != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+"id > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+"id < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+"id >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+"id <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+"id IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Email != nil {
		conditions = append(conditions, aliasPrefix+"email = ?")
		values = append(values, f.Email)
	}
	if f.EmailNe != nil {
		conditions = append(conditions, aliasPrefix+"email != ?")
		values = append(values, f.EmailNe)
	}
	if f.EmailGt != nil {
		conditions = append(conditions, aliasPrefix+"email > ?")
		values = append(values, f.EmailGt)
	}
	if f.EmailLt != nil {
		conditions = append(conditions, aliasPrefix+"email < ?")
		values = append(values, f.EmailLt)
	}
	if f.EmailGte != nil {
		conditions = append(conditions, aliasPrefix+"email >= ?")
		values = append(values, f.EmailGte)
	}
	if f.EmailLte != nil {
		conditions = append(conditions, aliasPrefix+"email <= ?")
		values = append(values, f.EmailLte)
	}
	if f.EmailIn != nil {
		conditions = append(conditions, aliasPrefix+"email IN (?)")
		values = append(values, f.EmailIn)
	}
	if f.EmailLike != nil {
		conditions = append(conditions, aliasPrefix+"email LIKE ?")
		values = append(values, strings.ReplaceAll(strings.ReplaceAll(*f.EmailLike, "?", "_"), "*", "%"))
	}
	if f.EmailPrefix != nil {
		conditions = append(conditions, aliasPrefix+"email LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailPrefix))
	}
	if f.EmailSuffix != nil {
		conditions = append(conditions, aliasPrefix+"email LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailSuffix))
	}

	if f.FirstName != nil {
		conditions = append(conditions, aliasPrefix+"firstName = ?")
		values = append(values, f.FirstName)
	}
	if f.FirstNameNe != nil {
		conditions = append(conditions, aliasPrefix+"firstName != ?")
		values = append(values, f.FirstNameNe)
	}
	if f.FirstNameGt != nil {
		conditions = append(conditions, aliasPrefix+"firstName > ?")
		values = append(values, f.FirstNameGt)
	}
	if f.FirstNameLt != nil {
		conditions = append(conditions, aliasPrefix+"firstName < ?")
		values = append(values, f.FirstNameLt)
	}
	if f.FirstNameGte != nil {
		conditions = append(conditions, aliasPrefix+"firstName >= ?")
		values = append(values, f.FirstNameGte)
	}
	if f.FirstNameLte != nil {
		conditions = append(conditions, aliasPrefix+"firstName <= ?")
		values = append(values, f.FirstNameLte)
	}
	if f.FirstNameIn != nil {
		conditions = append(conditions, aliasPrefix+"firstName IN (?)")
		values = append(values, f.FirstNameIn)
	}
	if f.FirstNameLike != nil {
		conditions = append(conditions, aliasPrefix+"firstName LIKE ?")
		values = append(values, strings.ReplaceAll(strings.ReplaceAll(*f.FirstNameLike, "?", "_"), "*", "%"))
	}
	if f.FirstNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+"firstName LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNamePrefix))
	}
	if f.FirstNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+"firstName LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameSuffix))
	}

	if f.LastName != nil {
		conditions = append(conditions, aliasPrefix+"lastName = ?")
		values = append(values, f.LastName)
	}
	if f.LastNameNe != nil {
		conditions = append(conditions, aliasPrefix+"lastName != ?")
		values = append(values, f.LastNameNe)
	}
	if f.LastNameGt != nil {
		conditions = append(conditions, aliasPrefix+"lastName > ?")
		values = append(values, f.LastNameGt)
	}
	if f.LastNameLt != nil {
		conditions = append(conditions, aliasPrefix+"lastName < ?")
		values = append(values, f.LastNameLt)
	}
	if f.LastNameGte != nil {
		conditions = append(conditions, aliasPrefix+"lastName >= ?")
		values = append(values, f.LastNameGte)
	}
	if f.LastNameLte != nil {
		conditions = append(conditions, aliasPrefix+"lastName <= ?")
		values = append(values, f.LastNameLte)
	}
	if f.LastNameIn != nil {
		conditions = append(conditions, aliasPrefix+"lastName IN (?)")
		values = append(values, f.LastNameIn)
	}
	if f.LastNameLike != nil {
		conditions = append(conditions, aliasPrefix+"lastName LIKE ?")
		values = append(values, strings.ReplaceAll(strings.ReplaceAll(*f.LastNameLike, "?", "_"), "*", "%"))
	}
	if f.LastNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+"lastName LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNamePrefix))
	}
	if f.LastNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+"lastName LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameSuffix))
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+"createdAt != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+"createdAt >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+"createdAt <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+"createdAt IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	return
}

func (f *TaskFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, "tasks", wheres, values, joins)
}
func (f *TaskFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := alias + "."

	_where, _values := f.WhereContent(aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Assignee != nil {
		_alias := alias + "_assignee"
		*joins = append(*joins, "LEFT JOIN users "+_alias+" ON "+_alias+".id = "+alias+".assigneeId")
		err := f.Assignee.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *TaskFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		for _, or := range f.Or {
			_cond, _values := or.WhereContent(aliasPrefix)
			cs = append(cs, _cond...)
			vs = append(vs, _values...)
		}
		conditions = append(conditions, "("+strings.Join(cs, " OR ")+")")
		values = append(values, vs...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		for _, or := range f.Or {
			_cond, _values := or.WhereContent(aliasPrefix)
			cs = append(cs, _cond...)
			vs = append(vs, _values...)
		}
		conditions = append(conditions, strings.Join(cs, " AND "))
		values = append(values, vs...)
	}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+"id = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+"id != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+"id > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+"id < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+"id >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+"id <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+"id IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Title != nil {
		conditions = append(conditions, aliasPrefix+"title = ?")
		values = append(values, f.Title)
	}
	if f.TitleNe != nil {
		conditions = append(conditions, aliasPrefix+"title != ?")
		values = append(values, f.TitleNe)
	}
	if f.TitleGt != nil {
		conditions = append(conditions, aliasPrefix+"title > ?")
		values = append(values, f.TitleGt)
	}
	if f.TitleLt != nil {
		conditions = append(conditions, aliasPrefix+"title < ?")
		values = append(values, f.TitleLt)
	}
	if f.TitleGte != nil {
		conditions = append(conditions, aliasPrefix+"title >= ?")
		values = append(values, f.TitleGte)
	}
	if f.TitleLte != nil {
		conditions = append(conditions, aliasPrefix+"title <= ?")
		values = append(values, f.TitleLte)
	}
	if f.TitleIn != nil {
		conditions = append(conditions, aliasPrefix+"title IN (?)")
		values = append(values, f.TitleIn)
	}
	if f.TitleLike != nil {
		conditions = append(conditions, aliasPrefix+"title LIKE ?")
		values = append(values, strings.ReplaceAll(strings.ReplaceAll(*f.TitleLike, "?", "_"), "*", "%"))
	}
	if f.TitlePrefix != nil {
		conditions = append(conditions, aliasPrefix+"title LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TitlePrefix))
	}
	if f.TitleSuffix != nil {
		conditions = append(conditions, aliasPrefix+"title LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TitleSuffix))
	}

	if f.Completed != nil {
		conditions = append(conditions, aliasPrefix+"completed = ?")
		values = append(values, f.Completed)
	}
	if f.CompletedNe != nil {
		conditions = append(conditions, aliasPrefix+"completed != ?")
		values = append(values, f.CompletedNe)
	}
	if f.CompletedGt != nil {
		conditions = append(conditions, aliasPrefix+"completed > ?")
		values = append(values, f.CompletedGt)
	}
	if f.CompletedLt != nil {
		conditions = append(conditions, aliasPrefix+"completed < ?")
		values = append(values, f.CompletedLt)
	}
	if f.CompletedGte != nil {
		conditions = append(conditions, aliasPrefix+"completed >= ?")
		values = append(values, f.CompletedGte)
	}
	if f.CompletedLte != nil {
		conditions = append(conditions, aliasPrefix+"completed <= ?")
		values = append(values, f.CompletedLte)
	}
	if f.CompletedIn != nil {
		conditions = append(conditions, aliasPrefix+"completed IN (?)")
		values = append(values, f.CompletedIn)
	}

	if f.DueDate != nil {
		conditions = append(conditions, aliasPrefix+"dueDate = ?")
		values = append(values, f.DueDate)
	}
	if f.DueDateNe != nil {
		conditions = append(conditions, aliasPrefix+"dueDate != ?")
		values = append(values, f.DueDateNe)
	}
	if f.DueDateGt != nil {
		conditions = append(conditions, aliasPrefix+"dueDate > ?")
		values = append(values, f.DueDateGt)
	}
	if f.DueDateLt != nil {
		conditions = append(conditions, aliasPrefix+"dueDate < ?")
		values = append(values, f.DueDateLt)
	}
	if f.DueDateGte != nil {
		conditions = append(conditions, aliasPrefix+"dueDate >= ?")
		values = append(values, f.DueDateGte)
	}
	if f.DueDateLte != nil {
		conditions = append(conditions, aliasPrefix+"dueDate <= ?")
		values = append(values, f.DueDateLte)
	}
	if f.DueDateIn != nil {
		conditions = append(conditions, aliasPrefix+"dueDate IN (?)")
		values = append(values, f.DueDateIn)
	}

	if f.Type != nil {
		conditions = append(conditions, aliasPrefix+"type = ?")
		values = append(values, f.Type)
	}
	if f.TypeNe != nil {
		conditions = append(conditions, aliasPrefix+"type != ?")
		values = append(values, f.TypeNe)
	}
	if f.TypeGt != nil {
		conditions = append(conditions, aliasPrefix+"type > ?")
		values = append(values, f.TypeGt)
	}
	if f.TypeLt != nil {
		conditions = append(conditions, aliasPrefix+"type < ?")
		values = append(values, f.TypeLt)
	}
	if f.TypeGte != nil {
		conditions = append(conditions, aliasPrefix+"type >= ?")
		values = append(values, f.TypeGte)
	}
	if f.TypeLte != nil {
		conditions = append(conditions, aliasPrefix+"type <= ?")
		values = append(values, f.TypeLte)
	}
	if f.TypeIn != nil {
		conditions = append(conditions, aliasPrefix+"type IN (?)")
		values = append(values, f.TypeIn)
	}

	if f.AssigneeID != nil {
		conditions = append(conditions, aliasPrefix+"assigneeId = ?")
		values = append(values, f.AssigneeID)
	}
	if f.AssigneeIDNe != nil {
		conditions = append(conditions, aliasPrefix+"assigneeId != ?")
		values = append(values, f.AssigneeIDNe)
	}
	if f.AssigneeIDGt != nil {
		conditions = append(conditions, aliasPrefix+"assigneeId > ?")
		values = append(values, f.AssigneeIDGt)
	}
	if f.AssigneeIDLt != nil {
		conditions = append(conditions, aliasPrefix+"assigneeId < ?")
		values = append(values, f.AssigneeIDLt)
	}
	if f.AssigneeIDGte != nil {
		conditions = append(conditions, aliasPrefix+"assigneeId >= ?")
		values = append(values, f.AssigneeIDGte)
	}
	if f.AssigneeIDLte != nil {
		conditions = append(conditions, aliasPrefix+"assigneeId <= ?")
		values = append(values, f.AssigneeIDLte)
	}
	if f.AssigneeIDIn != nil {
		conditions = append(conditions, aliasPrefix+"assigneeId IN (?)")
		values = append(values, f.AssigneeIDIn)
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+"createdAt != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+"createdAt >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+"createdAt <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+"createdAt IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	return
}
