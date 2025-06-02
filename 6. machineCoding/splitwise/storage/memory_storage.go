package storage

import "splitwise/model"

type MemoryStorage struct {
    Users    map[string]model.User
    Groups   map[string]model.Group
    Expenses map[string]model.Expense
}

func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{
        Users:    make(map[string]model.User),
        Groups:   make(map[string]model.Group),
        Expenses: make(map[string]model.Expense),
    }
}

func (m *MemoryStorage) AddUser(user model.User) { m.Users[user.ID] = user }
func (m *MemoryStorage) GetUser(id string) (model.User, bool) { u, ok := m.Users[id]; return u, ok }
func (m *MemoryStorage) AddGroup(group model.Group) { m.Groups[group.ID] = group }
func (m *MemoryStorage) GetGroup(id string) (model.Group, bool) { g, ok := m.Groups[id]; return g, ok }
func (m *MemoryStorage) AddExpense(expense model.Expense) { m.Expenses[expense.ID] = expense }
func (m *MemoryStorage) GetExpense(id string) (model.Expense, bool) { e, ok := m.Expenses[id]; return e, ok }
func (m *MemoryStorage) UpdateGroup(group model.Group) { m.Groups[group.ID] = group }
