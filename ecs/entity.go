package ecs

import "reflect"

type Entity struct {
	id         int
	components map[reflect.Type]any
}

// NewEntity creates a new ECS entity that is assigned an ID.
func NewEntity() *Entity {
	return &Entity{}
}

// Add component to the entity.
func (e *Entity) Add(components ...any) {
	for _, component := range components {
		e.components[reflect.TypeOf(component)] = component
	}
}

// Get the component of the entity.
func (e *Entity) Get(component reflect.Type) any {
	return e.components[component]
}

// Has the entity got the component.
func (e *Entity) Has(component reflect.Type) bool {
	_, ok := e.components[component]
	return ok
}
