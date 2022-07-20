package misc

import (
	"sync"
)

type MutexRegistry struct {
	registeredMapLock sync.Mutex
	registeredMap     map[string]string
}

func NewMutexRegistry() *MutexRegistry {
	return &MutexRegistry{
		registeredMapLock: sync.Mutex{},
		registeredMap:     map[string]string{},
	}
}

func (mutexRegistry *MutexRegistry) Register(key string) (string, bool) {
	mutexRegistry.registeredMapLock.Lock()
	defer mutexRegistry.registeredMapLock.Unlock()
	id, found := mutexRegistry.registeredMap[key]
	if !found {
		id = RandomString(8)
		mutexRegistry.registeredMap[key] = id
	}
	return id, !found
}

func (mutexRegistry *MutexRegistry) Unregister(key string, registeredId string) bool {
	mutexRegistry.registeredMapLock.Lock()
	defer mutexRegistry.registeredMapLock.Unlock()
	id, found := mutexRegistry.registeredMap[key]
	if found && id == registeredId {
		delete(mutexRegistry.registeredMap, key)
		return true
	}
	return false
}

func (mutexRegistry *MutexRegistry) IsRegistered(key string) bool {
	mutexRegistry.registeredMapLock.Lock()
	defer mutexRegistry.registeredMapLock.Unlock()
	_, found := mutexRegistry.registeredMap[key]
	return found
}

func (mutexRegistry *MutexRegistry) AllRegistered() map[string]string {
	mutexRegistry.registeredMapLock.Lock()
	defer mutexRegistry.registeredMapLock.Unlock()
	newMap := map[string]string{}
	for key, value := range mutexRegistry.registeredMap {
		newMap[key] = value
	}
	return newMap
}
