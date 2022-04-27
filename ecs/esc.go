package esc

type manager struct {
	entities map[entity][]any
}

type entity struct {
	id int
}

// New ESC manager.
func New() *manager {
	return &manager{
		entities: make(map[entity][]any),
	}
}

// NewEntity creates a new ECS entity that is assigned an ID.
func (m *manager) NewEntity() *entity {
	return &entity{}
}

// AddComponent to a specfic entity within the ESC.
func (m *manager) AddComponent(entityName entity, component any) {
	m.entities[entityName] = append(m.entities[entityName], component)
}
