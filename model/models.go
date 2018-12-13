// Package models provides the models
// for the data stored in the datastore.
package models

import "github.com/jinzhu/gorm"

// User is the owner of accounts.
type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);"`
}

// Type is the type of an account in a transaction.
type Type struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)"`
}

// TableName sets the table name for the Type model.
func (Type) TableName() string {
	return "account_types"
}

// Category is the category of an account.
// It is independent of the account type.
type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)"`
}

// TableName sets the table name for the Category model.
func (Category) TableName() string {
	return "categories"
}

// Account is a main model of the budget.
// Every transaction will be assigned to exactly one account.
type Account struct {
	gorm.Model
	User       User
	UserID     int
	Name       string `gorm:"type:varchar(255)"`
	Type       Type
	TypeID     int
	Category   Category
	CategoryID int
	Debit      bool
	Closed     bool
}

// TransactionType is just the transaction type.
type TransactionType struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)"`
}

// TableName sets the table name for the TransactionType model.
func (TransactionType) TableName() string {
	return "transaction_types"
}

// Transaction is what forms the meat of the budget.
// Every transaction is associated to an account.
type Transaction struct {
	gorm.Model
	Account   Account
	AccountID int
	Value     currency.Decimal
	Group     int64
	Note      string `gorm:"type:varchar(255)"`
}
