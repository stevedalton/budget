// Package models provides the models
// for the data stored in the datastore.
package models

// User is the owner of accounts.
type User struct {
	Name string
}

// Type is the type of an account in a transaction.
type Type struct {
	Name string
}

// TableName sets the table name for the Type model.
func (Type) TableName() string {
	return "account_types"
}

// Category is the category of an account.
// It is independent of the account type.
type Category struct {
	Name string
}

// TableName sets the table name for the Category model.
func (Category) TableName() string {
	return "categories"
}

// Account is a main model of the budget.
// Every transaction will be assigned to exactly one account.
type Account struct {
	User     User
	Name     string
	Type     Type
	Category Category
	Debit    bool
	Closed   bool
}

// TransactionType is just the transaction type.
type TransactionType struct {
	Name string
}

// TableName sets the table name for the TransactionType model.
func (TransactionType) TableName() string {
	return "transaction_types"
}

// Transaction is what forms the meat of the budget.
// Every transaction is associated to an account.
type Transaction struct {
	Account Account
	Value   currency.Decimal
	Group   int64
	Note    string
}
