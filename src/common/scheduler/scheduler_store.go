package scheduler

import "strings"

//Store define the basic operations for storing and managing policy watcher.
type Store interface {
	//Put a new policy in.
	Put(key string, value *Watcher)

	//Get the corresponding policy with the key.
	Get(key string) *Watcher

	//Exists is to check if the key existing in the store.
	Exists(key string) bool

	//Remove the specified policy and return its reference.
	Remove(key string) *Watcher

	//Size return the total count of items in store.
	Size() uint32

	//GetAll is to get all the items in the store.
	GetAll() []*Watcher

	//Clear store.
	Clear()
}

//DefaultStore implements Store interface to keep the scheduled policies.
//Not support concurrent sync.
type DefaultStore struct {
	//Map used to keep the policy list.
	data map[string]*Watcher
}

//NewDefaultStore is used to create a new store and return the pointer reference.
func NewDefaultStore() *DefaultStore {
	return &DefaultStore{make(map[string]*Watcher)}
}

//Put a policy into store.
func (cs *DefaultStore) Put(key string, value *Watcher) {
	if strings.TrimSpace(key) == "" || value == nil {
		return
	}

	cs.data[key] = value
}

//Get policy via key.
func (cs *DefaultStore) Get(key string) *Watcher {
	if strings.TrimSpace(key) == "" {
		return nil
	}

	return cs.data[key]
}

//Exists is used to check whether or not the key exists in store.
func (cs *DefaultStore) Exists(key string) bool {
	if strings.TrimSpace(key) == "" {
		return false
	}

	_, ok := cs.data[key]

	return ok
}

//Remove is to delete the specified policy.
func (cs *DefaultStore) Remove(key string) *Watcher {
	if !cs.Exists(key) {
		return nil
	}

	if wt, ok := cs.data[key]; ok {
		delete(cs.data, key)
		return wt
	}

	return nil
}

//Size return the total count of items in store.
func (cs *DefaultStore) Size() uint32 {
	return (uint32)(len(cs.data))
}

//GetAll to get all the items of store.
func (cs *DefaultStore) GetAll() []*Watcher {
	all := []*Watcher{}

	for _, v := range cs.data {
		all = append(all, v)
	}

	return all
}

//Clear all the items in store.
func (cs *DefaultStore) Clear() {
	if cs.Size() == 0 {
		return
	}

	for k := range cs.data {
		delete(cs.data, k)
	}
}
