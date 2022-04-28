package ecs

import "reflect"

type Manager struct {
	entities []*Entity
}

// New ecs manager.
func New() *Manager {
	return &Manager{
		entities: make([]*Entity, 0),
	}
}

// AddEntity to the ecs.
func (m *Manager) Add(e *Entity) {
	m.entities = append(m.entities, e)
}

// Query the ecs.
func (m *Manager) Query(components ...reflect.Type) []*Entity {
	var res []*Entity

	for _, component := range components {
		for _, e := range m.entities {
			if e.Has(component) {
				res = append(res, e)
			}
		}
	}

	return res
}

// Remove an entity.
func (m *Manager) Remove(entity *Entity) {
	for i, e := range m.entities {
		if e == entity {
			m.entities = append(m.entities[:i], m.entities[i+1:]...)
			return
		}
	}
}
