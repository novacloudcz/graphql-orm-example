package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func (f *CompanyFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *CompanyFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, "companies", wheres, values, joins)
}
func (f *CompanyFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			err := or.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Employees != nil {
		_alias := alias + "_employees"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("company_employees")+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("company_id")+" LEFT JOIN "+dialect.Quote("users")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employee_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.Employees.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *CompanyFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}
	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}
	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}
	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}
	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}
	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}
	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}
	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}
	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}
	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}
	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}
	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}
	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}
	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}
	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}
	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}
	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}
	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}
	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}
	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}
	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}
	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *CompanyFilterType) AndWith(f2 ...*CompanyFilterType) *CompanyFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &CompanyFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *CompanyFilterType) OrWith(f2 ...*CompanyFilterType) *CompanyFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &CompanyFilterType{
		Or: append(_f2, f),
	}
}

func (f *UserFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *UserFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, "users", wheres, values, joins)
}
func (f *UserFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			err := or.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Tasks != nil {
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("tasks")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("assigneeId")+" = "+dialect.Quote(alias)+".id")
		err := f.Tasks.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Companies != nil {
		_alias := alias + "_companies"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("company_employees")+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employee_id")+" LEFT JOIN "+dialect.Quote("companies")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("company_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.Companies.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Friends != nil {
		_alias := alias + "_friends"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("user_friends")+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("friend_id")+" LEFT JOIN "+dialect.Quote("users")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("friend_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.Friends.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *UserFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Email != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" = ?")
		values = append(values, f.Email)
	}
	if f.EmailNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" != ?")
		values = append(values, f.EmailNe)
	}
	if f.EmailGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" > ?")
		values = append(values, f.EmailGt)
	}
	if f.EmailLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" < ?")
		values = append(values, f.EmailLt)
	}
	if f.EmailGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" >= ?")
		values = append(values, f.EmailGte)
	}
	if f.EmailLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" <= ?")
		values = append(values, f.EmailLte)
	}
	if f.EmailIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IN (?)")
		values = append(values, f.EmailIn)
	}
	if f.EmailLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailLike, "?", "_", -1), "*", "%", -1))
	}
	if f.EmailPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailPrefix))
	}
	if f.EmailSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailSuffix))
	}

	if f.FirstName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" = ?")
		values = append(values, f.FirstName)
	}
	if f.FirstNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" != ?")
		values = append(values, f.FirstNameNe)
	}
	if f.FirstNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" > ?")
		values = append(values, f.FirstNameGt)
	}
	if f.FirstNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" < ?")
		values = append(values, f.FirstNameLt)
	}
	if f.FirstNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" >= ?")
		values = append(values, f.FirstNameGte)
	}
	if f.FirstNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" <= ?")
		values = append(values, f.FirstNameLte)
	}
	if f.FirstNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" IN (?)")
		values = append(values, f.FirstNameIn)
	}
	if f.FirstNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FirstNameLike, "?", "_", -1), "*", "%", -1))
	}
	if f.FirstNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNamePrefix))
	}
	if f.FirstNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameSuffix))
	}

	if f.LastName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" = ?")
		values = append(values, f.LastName)
	}
	if f.LastNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" != ?")
		values = append(values, f.LastNameNe)
	}
	if f.LastNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" > ?")
		values = append(values, f.LastNameGt)
	}
	if f.LastNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" < ?")
		values = append(values, f.LastNameLt)
	}
	if f.LastNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" >= ?")
		values = append(values, f.LastNameGte)
	}
	if f.LastNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" <= ?")
		values = append(values, f.LastNameLte)
	}
	if f.LastNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" IN (?)")
		values = append(values, f.LastNameIn)
	}
	if f.LastNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastNameLike, "?", "_", -1), "*", "%", -1))
	}
	if f.LastNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNamePrefix))
	}
	if f.LastNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameSuffix))
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}
	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}
	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}
	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}
	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}
	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}
	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}
	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}
	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}
	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}
	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}
	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}
	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *UserFilterType) AndWith(f2 ...*UserFilterType) *UserFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *UserFilterType) OrWith(f2 ...*UserFilterType) *UserFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserFilterType{
		Or: append(_f2, f),
	}
}

func (f *TaskFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *TaskFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, "tasks", wheres, values, joins)
}
func (f *TaskFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			err := or.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Assignee != nil {
		_alias := alias + "_assignee"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("users")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("assigneeId"))
		err := f.Assignee.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *TaskFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Title != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" = ?")
		values = append(values, f.Title)
	}
	if f.TitleNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" != ?")
		values = append(values, f.TitleNe)
	}
	if f.TitleGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" > ?")
		values = append(values, f.TitleGt)
	}
	if f.TitleLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" < ?")
		values = append(values, f.TitleLt)
	}
	if f.TitleGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" >= ?")
		values = append(values, f.TitleGte)
	}
	if f.TitleLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" <= ?")
		values = append(values, f.TitleLte)
	}
	if f.TitleIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" IN (?)")
		values = append(values, f.TitleIn)
	}
	if f.TitleLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TitleLike, "?", "_", -1), "*", "%", -1))
	}
	if f.TitlePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TitlePrefix))
	}
	if f.TitleSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TitleSuffix))
	}

	if f.Completed != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" = ?")
		values = append(values, f.Completed)
	}
	if f.CompletedNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" != ?")
		values = append(values, f.CompletedNe)
	}
	if f.CompletedGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" > ?")
		values = append(values, f.CompletedGt)
	}
	if f.CompletedLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" < ?")
		values = append(values, f.CompletedLt)
	}
	if f.CompletedGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" >= ?")
		values = append(values, f.CompletedGte)
	}
	if f.CompletedLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" <= ?")
		values = append(values, f.CompletedLte)
	}
	if f.CompletedIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" IN (?)")
		values = append(values, f.CompletedIn)
	}

	if f.DueDate != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" = ?")
		values = append(values, f.DueDate)
	}
	if f.DueDateNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" != ?")
		values = append(values, f.DueDateNe)
	}
	if f.DueDateGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" > ?")
		values = append(values, f.DueDateGt)
	}
	if f.DueDateLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" < ?")
		values = append(values, f.DueDateLt)
	}
	if f.DueDateGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" >= ?")
		values = append(values, f.DueDateGte)
	}
	if f.DueDateLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" <= ?")
		values = append(values, f.DueDateLte)
	}
	if f.DueDateIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" IN (?)")
		values = append(values, f.DueDateIn)
	}

	if f.Type != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" = ?")
		values = append(values, f.Type)
	}
	if f.TypeNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" != ?")
		values = append(values, f.TypeNe)
	}
	if f.TypeGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" > ?")
		values = append(values, f.TypeGt)
	}
	if f.TypeLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" < ?")
		values = append(values, f.TypeLt)
	}
	if f.TypeGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" >= ?")
		values = append(values, f.TypeGte)
	}
	if f.TypeLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" <= ?")
		values = append(values, f.TypeLte)
	}
	if f.TypeIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" IN (?)")
		values = append(values, f.TypeIn)
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}
	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}
	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}
	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}
	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}
	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}
	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}
	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}
	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}
	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.AssigneeID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" = ?")
		values = append(values, f.AssigneeID)
	}
	if f.AssigneeIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" != ?")
		values = append(values, f.AssigneeIDNe)
	}
	if f.AssigneeIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" > ?")
		values = append(values, f.AssigneeIDGt)
	}
	if f.AssigneeIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" < ?")
		values = append(values, f.AssigneeIDLt)
	}
	if f.AssigneeIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" >= ?")
		values = append(values, f.AssigneeIDGte)
	}
	if f.AssigneeIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" <= ?")
		values = append(values, f.AssigneeIDLte)
	}
	if f.AssigneeIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" IN (?)")
		values = append(values, f.AssigneeIDIn)
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}
	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}
	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}
	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}
	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}
	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}
	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}
	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}
	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}
	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}
	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}
	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}
	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *TaskFilterType) AndWith(f2 ...*TaskFilterType) *TaskFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TaskFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *TaskFilterType) OrWith(f2 ...*TaskFilterType) *TaskFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TaskFilterType{
		Or: append(_f2, f),
	}
}
