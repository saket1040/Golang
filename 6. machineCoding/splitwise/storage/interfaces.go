package storage

import "splitwise/model"

type Storage interface {
    AddUser(user model.User)
    GetUser(id string) (model.User, bool)
    AddGroup(group model.Group)
    GetGroup(id string) (model.Group, bool)
    AddExpense(expense model.Expense)
    GetExpense(id string) (model.Expense, bool)
    UpdateGroup(group model.Group)
}
