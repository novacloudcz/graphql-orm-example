package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/ast"
)

type CompanyQueryFilter struct {
	Query *string
}

func (qf *CompanyQueryFilter) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}
	fields := []*ast.Field{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		fields = append(fields, f.Field)
	}

	ors := []string{}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		if err := qf.applyQueryWithFields(fields, part, "companies", &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *CompanyQueryFilter) applyQueryWithFields(fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}
	aliasPrefix := alias + "."

	fieldsMap := map[string]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = f
	}

	if _, ok := fieldsMap["name"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]sname LIKE ? OR %[1]sname LIKE ?", aliasPrefix))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if f, ok := fieldsMap["employees"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_employees"
		*joins = append(*joins, "LEFT JOIN company_employees "+_alias+"_jointable ON "+alias+".id = "+_alias+"_jointable.company_id LEFT JOIN users "+_alias+" ON "+_alias+"_jointable.employee_id = "+_alias+".id")

		for _, s := range f.SelectionSet {
			if f, ok := s.(*ast.Field); ok {
				_fields = append(_fields, f)
			}
		}
		q := UserQueryFilter{qf.Query}
		err := q.applyQueryWithFields(_fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

type UserQueryFilter struct {
	Query *string
}

func (qf *UserQueryFilter) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}
	fields := []*ast.Field{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		fields = append(fields, f.Field)
	}

	ors := []string{}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		if err := qf.applyQueryWithFields(fields, part, "users", &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *UserQueryFilter) applyQueryWithFields(fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}
	aliasPrefix := alias + "."

	fieldsMap := map[string]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = f
	}

	if _, ok := fieldsMap["email"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]semail LIKE ? OR %[1]semail LIKE ?", aliasPrefix))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if _, ok := fieldsMap["firstName"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]sfirstName LIKE ? OR %[1]sfirstName LIKE ?", aliasPrefix))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if _, ok := fieldsMap["lastName"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]slastName LIKE ? OR %[1]slastName LIKE ?", aliasPrefix))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if f, ok := fieldsMap["employees"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN tasks "+_alias+" ON "+_alias+".assigneeId = "+alias+".id")

		for _, s := range f.SelectionSet {
			if f, ok := s.(*ast.Field); ok {
				_fields = append(_fields, f)
			}
		}
		q := TaskQueryFilter{qf.Query}
		err := q.applyQueryWithFields(_fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if f, ok := fieldsMap["employees"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_companies"
		*joins = append(*joins, "LEFT JOIN company_employees "+_alias+"_jointable ON "+alias+".id = "+_alias+"_jointable.employee_id LEFT JOIN companies "+_alias+" ON "+_alias+"_jointable.company_id = "+_alias+".id")

		for _, s := range f.SelectionSet {
			if f, ok := s.(*ast.Field); ok {
				_fields = append(_fields, f)
			}
		}
		q := CompanyQueryFilter{qf.Query}
		err := q.applyQueryWithFields(_fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if f, ok := fieldsMap["employees"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_friends"
		*joins = append(*joins, "LEFT JOIN user_friends "+_alias+"_jointable ON "+alias+".id = "+_alias+"_jointable.friend_id LEFT JOIN users "+_alias+" ON "+_alias+"_jointable.friend_id = "+_alias+".id")

		for _, s := range f.SelectionSet {
			if f, ok := s.(*ast.Field); ok {
				_fields = append(_fields, f)
			}
		}
		q := UserQueryFilter{qf.Query}
		err := q.applyQueryWithFields(_fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

type TaskQueryFilter struct {
	Query *string
}

func (qf *TaskQueryFilter) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}
	fields := []*ast.Field{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		fields = append(fields, f.Field)
	}

	ors := []string{}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		if err := qf.applyQueryWithFields(fields, part, "tasks", &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *TaskQueryFilter) applyQueryWithFields(fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}
	aliasPrefix := alias + "."

	fieldsMap := map[string]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = f
	}

	if _, ok := fieldsMap["title"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]stitle LIKE ? OR %[1]stitle LIKE ?", aliasPrefix))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if f, ok := fieldsMap["employees"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_assignee"
		*joins = append(*joins, "LEFT JOIN users "+_alias+" ON "+_alias+".id = "+alias+".assigneeId")

		for _, s := range f.SelectionSet {
			if f, ok := s.(*ast.Field); ok {
				_fields = append(_fields, f)
			}
		}
		q := UserQueryFilter{qf.Query}
		err := q.applyQueryWithFields(_fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
