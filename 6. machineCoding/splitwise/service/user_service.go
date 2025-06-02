package service

import (
    "splitwise/model"
    "splitwise/storage"
)

type UserService struct {
    store storage.Storage
}

func NewUserService(store storage.Storage) *UserService {
    return &UserService{store: store}
}

func (u *UserService) CreateUser(id, name, email string) {
    user := model.User{ID: id, Name: name, Email: email}
    u.store.AddUser(user)
}