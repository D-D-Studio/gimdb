package domain

import (
	"errors"
	"sync"
)

type Workspace struct {
	storages *sync.Map
}

func NewWorkspace() *Workspace {
	workspace := new(Workspace)

	workspace.storages = new(sync.Map)

	return workspace
}

func (w Workspace) GetStorage(name string) (*Storage, error) {
	storage, ok := w.storages.Load(name)

	if !ok {
		return nil, errors.New("cannot find storage")
	}

	return storage.(*Storage), nil
}

func (w *Workspace) CreateStorage(name string) error {
	_, err := w.GetStorage(name)

	if err == nil {
		return errors.New("storage already exists")
	}

	w.storages.Store(name, NewStorage())

	return nil
}

func (w *Workspace) DeleteStorage(name string) error {
	_, loaded := w.storages.LoadAndDelete(name)

	if !loaded {
		return errors.New("cannot find storage")
	}

	return nil
}

func (w *Workspace) IsStorageExists(name string) bool {
	_, ok := w.storages.Load(name)

	return ok
}
