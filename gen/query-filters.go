package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/vektah/gqlparser/ast"
)

type CompanyQueryFilter struct {
	Query *string
}

func (qf *CompanyQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field.")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("companies"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *CompanyQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["name"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("name")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if fs, ok := fieldsMap["employees"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_employees"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("company_employees"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("company_id")+" LEFT JOIN "+dialect.Quote("users")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employee_id")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := UserQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

type UserQueryFilter struct {
	Query *string
}

func (qf *UserQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field.")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("users"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *UserQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["email"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("email")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["firstName"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("firstName")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["lastName"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("lastName")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if fs, ok := fieldsMap["tasks"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("tasks"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("assigneeId")+" = "+dialect.Quote(alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := TaskQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["companies"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_companies"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("company_employees"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employee_id")+" LEFT JOIN "+dialect.Quote("companies")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("company_id")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := CompanyQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["friends"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_friends"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("user_friends"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("friend_id")+" LEFT JOIN "+dialect.Quote("users")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("friend_id")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := UserQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

type TaskQueryFilter struct {
	Query *string
}

func (qf *TaskQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field.")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("tasks"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *TaskQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["title"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("title")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["description"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("description")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if fs, ok := fieldsMap["assignee"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_assignee"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("assigneeId"))

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := UserQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
