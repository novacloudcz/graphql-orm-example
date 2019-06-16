package gen

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func (f *CompanyFilterType) Apply(db *gorm.DB) (*gorm.DB, error) {
	return f.ApplyWithAlias(db, "companies")
}
func (f *CompanyFilterType) ApplyWithAlias(db *gorm.DB, alias string) (*gorm.DB, error) {
	if f == nil {
		return db, nil
	}
	aliasPrefix := alias
	if aliasPrefix != "" {
		aliasPrefix += "."
	}

	if f.ID != nil {
		db = db.Where(aliasPrefix+"id = ?", f.ID)
	}
	if f.IDNe != nil {
		db = db.Where(aliasPrefix+"id != ?", f.IDNe)
	}
	if f.IDGt != nil {
		db = db.Where(aliasPrefix+"id > ?", f.IDGt)
	}
	if f.IDLt != nil {
		db = db.Where(aliasPrefix+"id < ?", f.IDLt)
	}
	if f.IDGte != nil {
		db = db.Where(aliasPrefix+"id >= ?", f.IDGte)
	}
	if f.IDLte != nil {
		db = db.Where(aliasPrefix+"id <= ?", f.IDLte)
	}
	if f.IDIn != nil {
		db = db.Where(aliasPrefix+"id IN (?)", f.IDIn)
	}

	if f.Name != nil {
		db = db.Where(aliasPrefix+"name = ?", f.Name)
	}
	if f.NameNe != nil {
		db = db.Where(aliasPrefix+"name != ?", f.NameNe)
	}
	if f.NameGt != nil {
		db = db.Where(aliasPrefix+"name > ?", f.NameGt)
	}
	if f.NameLt != nil {
		db = db.Where(aliasPrefix+"name < ?", f.NameLt)
	}
	if f.NameGte != nil {
		db = db.Where(aliasPrefix+"name >= ?", f.NameGte)
	}
	if f.NameLte != nil {
		db = db.Where(aliasPrefix+"name <= ?", f.NameLte)
	}
	if f.NameIn != nil {
		db = db.Where(aliasPrefix+"name IN (?)", f.NameIn)
	}
	if f.NameLike != nil {
		db = db.Where(aliasPrefix+"name LIKE ?", strings.ReplaceAll(strings.ReplaceAll(*f.NameLike, "?", "_"), "*", "%"))
	}
	if f.NamePrefix != nil {
		db = db.Where(aliasPrefix+"name LIKE ?", fmt.Sprintf("%s%%", *f.NamePrefix))
	}
	if f.NameSuffix != nil {
		db = db.Where(aliasPrefix+"name LIKE ?", fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.UpdatedAt != nil {
		db = db.Where(aliasPrefix+"updatedAt = ?", f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		db = db.Where(aliasPrefix+"updatedAt != ?", f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		db = db.Where(aliasPrefix+"updatedAt > ?", f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		db = db.Where(aliasPrefix+"updatedAt < ?", f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		db = db.Where(aliasPrefix+"updatedAt >= ?", f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		db = db.Where(aliasPrefix+"updatedAt <= ?", f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		db = db.Where(aliasPrefix+"updatedAt IN (?)", f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		db = db.Where(aliasPrefix+"createdAt = ?", f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		db = db.Where(aliasPrefix+"createdAt != ?", f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		db = db.Where(aliasPrefix+"createdAt > ?", f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		db = db.Where(aliasPrefix+"createdAt < ?", f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		db = db.Where(aliasPrefix+"createdAt >= ?", f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		db = db.Where(aliasPrefix+"createdAt <= ?", f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		db = db.Where(aliasPrefix+"createdAt IN (?)", f.CreatedAtIn)
	}

	if f.Employees != nil {
		_alias := alias + "_employees"
		db = db.Joins("LEFT JOIN company_employees " + _alias + "_jointable ON " + alias + ".id = " + _alias + "_jointable.company_id LEFT JOIN users " + _alias + " ON " + _alias + "_jointable.employee_id = " + _alias + ".id")
		_db, err := f.Employees.ApplyWithAlias(db, _alias)
		if err != nil {
			return db, err
		}
		db = _db
	}

	return db, nil
}

func (f *UserFilterType) Apply(db *gorm.DB) (*gorm.DB, error) {
	return f.ApplyWithAlias(db, "users")
}
func (f *UserFilterType) ApplyWithAlias(db *gorm.DB, alias string) (*gorm.DB, error) {
	if f == nil {
		return db, nil
	}
	aliasPrefix := alias
	if aliasPrefix != "" {
		aliasPrefix += "."
	}

	if f.ID != nil {
		db = db.Where(aliasPrefix+"id = ?", f.ID)
	}
	if f.IDNe != nil {
		db = db.Where(aliasPrefix+"id != ?", f.IDNe)
	}
	if f.IDGt != nil {
		db = db.Where(aliasPrefix+"id > ?", f.IDGt)
	}
	if f.IDLt != nil {
		db = db.Where(aliasPrefix+"id < ?", f.IDLt)
	}
	if f.IDGte != nil {
		db = db.Where(aliasPrefix+"id >= ?", f.IDGte)
	}
	if f.IDLte != nil {
		db = db.Where(aliasPrefix+"id <= ?", f.IDLte)
	}
	if f.IDIn != nil {
		db = db.Where(aliasPrefix+"id IN (?)", f.IDIn)
	}

	if f.Email != nil {
		db = db.Where(aliasPrefix+"email = ?", f.Email)
	}
	if f.EmailNe != nil {
		db = db.Where(aliasPrefix+"email != ?", f.EmailNe)
	}
	if f.EmailGt != nil {
		db = db.Where(aliasPrefix+"email > ?", f.EmailGt)
	}
	if f.EmailLt != nil {
		db = db.Where(aliasPrefix+"email < ?", f.EmailLt)
	}
	if f.EmailGte != nil {
		db = db.Where(aliasPrefix+"email >= ?", f.EmailGte)
	}
	if f.EmailLte != nil {
		db = db.Where(aliasPrefix+"email <= ?", f.EmailLte)
	}
	if f.EmailIn != nil {
		db = db.Where(aliasPrefix+"email IN (?)", f.EmailIn)
	}
	if f.EmailLike != nil {
		db = db.Where(aliasPrefix+"email LIKE ?", strings.ReplaceAll(strings.ReplaceAll(*f.EmailLike, "?", "_"), "*", "%"))
	}
	if f.EmailPrefix != nil {
		db = db.Where(aliasPrefix+"email LIKE ?", fmt.Sprintf("%s%%", *f.EmailPrefix))
	}
	if f.EmailSuffix != nil {
		db = db.Where(aliasPrefix+"email LIKE ?", fmt.Sprintf("%%%s", *f.EmailSuffix))
	}

	if f.FirstName != nil {
		db = db.Where(aliasPrefix+"firstName = ?", f.FirstName)
	}
	if f.FirstNameNe != nil {
		db = db.Where(aliasPrefix+"firstName != ?", f.FirstNameNe)
	}
	if f.FirstNameGt != nil {
		db = db.Where(aliasPrefix+"firstName > ?", f.FirstNameGt)
	}
	if f.FirstNameLt != nil {
		db = db.Where(aliasPrefix+"firstName < ?", f.FirstNameLt)
	}
	if f.FirstNameGte != nil {
		db = db.Where(aliasPrefix+"firstName >= ?", f.FirstNameGte)
	}
	if f.FirstNameLte != nil {
		db = db.Where(aliasPrefix+"firstName <= ?", f.FirstNameLte)
	}
	if f.FirstNameIn != nil {
		db = db.Where(aliasPrefix+"firstName IN (?)", f.FirstNameIn)
	}
	if f.FirstNameLike != nil {
		db = db.Where(aliasPrefix+"firstName LIKE ?", strings.ReplaceAll(strings.ReplaceAll(*f.FirstNameLike, "?", "_"), "*", "%"))
	}
	if f.FirstNamePrefix != nil {
		db = db.Where(aliasPrefix+"firstName LIKE ?", fmt.Sprintf("%s%%", *f.FirstNamePrefix))
	}
	if f.FirstNameSuffix != nil {
		db = db.Where(aliasPrefix+"firstName LIKE ?", fmt.Sprintf("%%%s", *f.FirstNameSuffix))
	}

	if f.LastName != nil {
		db = db.Where(aliasPrefix+"lastName = ?", f.LastName)
	}
	if f.LastNameNe != nil {
		db = db.Where(aliasPrefix+"lastName != ?", f.LastNameNe)
	}
	if f.LastNameGt != nil {
		db = db.Where(aliasPrefix+"lastName > ?", f.LastNameGt)
	}
	if f.LastNameLt != nil {
		db = db.Where(aliasPrefix+"lastName < ?", f.LastNameLt)
	}
	if f.LastNameGte != nil {
		db = db.Where(aliasPrefix+"lastName >= ?", f.LastNameGte)
	}
	if f.LastNameLte != nil {
		db = db.Where(aliasPrefix+"lastName <= ?", f.LastNameLte)
	}
	if f.LastNameIn != nil {
		db = db.Where(aliasPrefix+"lastName IN (?)", f.LastNameIn)
	}
	if f.LastNameLike != nil {
		db = db.Where(aliasPrefix+"lastName LIKE ?", strings.ReplaceAll(strings.ReplaceAll(*f.LastNameLike, "?", "_"), "*", "%"))
	}
	if f.LastNamePrefix != nil {
		db = db.Where(aliasPrefix+"lastName LIKE ?", fmt.Sprintf("%s%%", *f.LastNamePrefix))
	}
	if f.LastNameSuffix != nil {
		db = db.Where(aliasPrefix+"lastName LIKE ?", fmt.Sprintf("%%%s", *f.LastNameSuffix))
	}

	if f.UpdatedAt != nil {
		db = db.Where(aliasPrefix+"updatedAt = ?", f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		db = db.Where(aliasPrefix+"updatedAt != ?", f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		db = db.Where(aliasPrefix+"updatedAt > ?", f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		db = db.Where(aliasPrefix+"updatedAt < ?", f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		db = db.Where(aliasPrefix+"updatedAt >= ?", f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		db = db.Where(aliasPrefix+"updatedAt <= ?", f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		db = db.Where(aliasPrefix+"updatedAt IN (?)", f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		db = db.Where(aliasPrefix+"createdAt = ?", f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		db = db.Where(aliasPrefix+"createdAt != ?", f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		db = db.Where(aliasPrefix+"createdAt > ?", f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		db = db.Where(aliasPrefix+"createdAt < ?", f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		db = db.Where(aliasPrefix+"createdAt >= ?", f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		db = db.Where(aliasPrefix+"createdAt <= ?", f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		db = db.Where(aliasPrefix+"createdAt IN (?)", f.CreatedAtIn)
	}

	if f.Tasks != nil {
		_alias := alias + "_tasks"
		db = db.Joins("LEFT JOIN tasks " + _alias + " ON " + _alias + ".assigneeId = " + alias + ".id")
		_db, err := f.Tasks.ApplyWithAlias(db, _alias)
		if err != nil {
			return db, err
		}
		db = _db
	}

	if f.Companies != nil {
		_alias := alias + "_companies"
		db = db.Joins("LEFT JOIN company_employees " + _alias + "_jointable ON " + alias + ".id = " + _alias + "_jointable.employee_id LEFT JOIN companies " + _alias + " ON " + _alias + "_jointable.company_id = " + _alias + ".id")
		_db, err := f.Companies.ApplyWithAlias(db, _alias)
		if err != nil {
			return db, err
		}
		db = _db
	}

	if f.Friends != nil {
		_alias := alias + "_friends"
		db = db.Joins("LEFT JOIN user_friends " + _alias + "_jointable ON " + alias + ".id = " + _alias + "_jointable.friend_id LEFT JOIN users " + _alias + " ON " + _alias + "_jointable.friend_id = " + _alias + ".id")
		_db, err := f.Friends.ApplyWithAlias(db, _alias)
		if err != nil {
			return db, err
		}
		db = _db
	}

	return db, nil
}

func (f *TaskFilterType) Apply(db *gorm.DB) (*gorm.DB, error) {
	return f.ApplyWithAlias(db, "tasks")
}
func (f *TaskFilterType) ApplyWithAlias(db *gorm.DB, alias string) (*gorm.DB, error) {
	if f == nil {
		return db, nil
	}
	aliasPrefix := alias
	if aliasPrefix != "" {
		aliasPrefix += "."
	}

	if f.ID != nil {
		db = db.Where(aliasPrefix+"id = ?", f.ID)
	}
	if f.IDNe != nil {
		db = db.Where(aliasPrefix+"id != ?", f.IDNe)
	}
	if f.IDGt != nil {
		db = db.Where(aliasPrefix+"id > ?", f.IDGt)
	}
	if f.IDLt != nil {
		db = db.Where(aliasPrefix+"id < ?", f.IDLt)
	}
	if f.IDGte != nil {
		db = db.Where(aliasPrefix+"id >= ?", f.IDGte)
	}
	if f.IDLte != nil {
		db = db.Where(aliasPrefix+"id <= ?", f.IDLte)
	}
	if f.IDIn != nil {
		db = db.Where(aliasPrefix+"id IN (?)", f.IDIn)
	}

	if f.Title != nil {
		db = db.Where(aliasPrefix+"title = ?", f.Title)
	}
	if f.TitleNe != nil {
		db = db.Where(aliasPrefix+"title != ?", f.TitleNe)
	}
	if f.TitleGt != nil {
		db = db.Where(aliasPrefix+"title > ?", f.TitleGt)
	}
	if f.TitleLt != nil {
		db = db.Where(aliasPrefix+"title < ?", f.TitleLt)
	}
	if f.TitleGte != nil {
		db = db.Where(aliasPrefix+"title >= ?", f.TitleGte)
	}
	if f.TitleLte != nil {
		db = db.Where(aliasPrefix+"title <= ?", f.TitleLte)
	}
	if f.TitleIn != nil {
		db = db.Where(aliasPrefix+"title IN (?)", f.TitleIn)
	}
	if f.TitleLike != nil {
		db = db.Where(aliasPrefix+"title LIKE ?", strings.ReplaceAll(strings.ReplaceAll(*f.TitleLike, "?", "_"), "*", "%"))
	}
	if f.TitlePrefix != nil {
		db = db.Where(aliasPrefix+"title LIKE ?", fmt.Sprintf("%s%%", *f.TitlePrefix))
	}
	if f.TitleSuffix != nil {
		db = db.Where(aliasPrefix+"title LIKE ?", fmt.Sprintf("%%%s", *f.TitleSuffix))
	}

	if f.Completed != nil {
		db = db.Where(aliasPrefix+"completed = ?", f.Completed)
	}
	if f.CompletedNe != nil {
		db = db.Where(aliasPrefix+"completed != ?", f.CompletedNe)
	}
	if f.CompletedGt != nil {
		db = db.Where(aliasPrefix+"completed > ?", f.CompletedGt)
	}
	if f.CompletedLt != nil {
		db = db.Where(aliasPrefix+"completed < ?", f.CompletedLt)
	}
	if f.CompletedGte != nil {
		db = db.Where(aliasPrefix+"completed >= ?", f.CompletedGte)
	}
	if f.CompletedLte != nil {
		db = db.Where(aliasPrefix+"completed <= ?", f.CompletedLte)
	}
	if f.CompletedIn != nil {
		db = db.Where(aliasPrefix+"completed IN (?)", f.CompletedIn)
	}

	if f.DueDate != nil {
		db = db.Where(aliasPrefix+"dueDate = ?", f.DueDate)
	}
	if f.DueDateNe != nil {
		db = db.Where(aliasPrefix+"dueDate != ?", f.DueDateNe)
	}
	if f.DueDateGt != nil {
		db = db.Where(aliasPrefix+"dueDate > ?", f.DueDateGt)
	}
	if f.DueDateLt != nil {
		db = db.Where(aliasPrefix+"dueDate < ?", f.DueDateLt)
	}
	if f.DueDateGte != nil {
		db = db.Where(aliasPrefix+"dueDate >= ?", f.DueDateGte)
	}
	if f.DueDateLte != nil {
		db = db.Where(aliasPrefix+"dueDate <= ?", f.DueDateLte)
	}
	if f.DueDateIn != nil {
		db = db.Where(aliasPrefix+"dueDate IN (?)", f.DueDateIn)
	}

	if f.Type != nil {
		db = db.Where(aliasPrefix+"type = ?", f.Type)
	}
	if f.TypeNe != nil {
		db = db.Where(aliasPrefix+"type != ?", f.TypeNe)
	}
	if f.TypeGt != nil {
		db = db.Where(aliasPrefix+"type > ?", f.TypeGt)
	}
	if f.TypeLt != nil {
		db = db.Where(aliasPrefix+"type < ?", f.TypeLt)
	}
	if f.TypeGte != nil {
		db = db.Where(aliasPrefix+"type >= ?", f.TypeGte)
	}
	if f.TypeLte != nil {
		db = db.Where(aliasPrefix+"type <= ?", f.TypeLte)
	}
	if f.TypeIn != nil {
		db = db.Where(aliasPrefix+"type IN (?)", f.TypeIn)
	}

	if f.UpdatedAt != nil {
		db = db.Where(aliasPrefix+"updatedAt = ?", f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		db = db.Where(aliasPrefix+"updatedAt != ?", f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		db = db.Where(aliasPrefix+"updatedAt > ?", f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		db = db.Where(aliasPrefix+"updatedAt < ?", f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		db = db.Where(aliasPrefix+"updatedAt >= ?", f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		db = db.Where(aliasPrefix+"updatedAt <= ?", f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		db = db.Where(aliasPrefix+"updatedAt IN (?)", f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		db = db.Where(aliasPrefix+"createdAt = ?", f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		db = db.Where(aliasPrefix+"createdAt != ?", f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		db = db.Where(aliasPrefix+"createdAt > ?", f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		db = db.Where(aliasPrefix+"createdAt < ?", f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		db = db.Where(aliasPrefix+"createdAt >= ?", f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		db = db.Where(aliasPrefix+"createdAt <= ?", f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		db = db.Where(aliasPrefix+"createdAt IN (?)", f.CreatedAtIn)
	}

	if f.Assignee != nil {
		_alias := alias + "_assignee"
		db = db.Joins("LEFT JOIN users " + _alias + " ON " + _alias + ".id = " + alias + ".assigneeId")
		_db, err := f.Assignee.ApplyWithAlias(db, _alias)
		if err != nil {
			return db, err
		}
		db = _db
	}

	return db, nil
}
