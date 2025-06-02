package service

import (
	"splitwise/model"
	"splitwise/storage"
)

type GroupService struct {
	store storage.Storage
}

func NewGroupService(store storage.Storage) *GroupService {
	return &GroupService{store: store}
}

func (g *GroupService) CreateGroup(id, name string, memberIDs []string) {
	group := model.Group{ID: id, Name: name, Members: memberIDs}
	g.store.AddGroup(group)
}
