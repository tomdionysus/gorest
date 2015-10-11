package gorest

import (
	// "errors"
)

type Datastore interface {
	HasName
	HasLogger

	Connect(url string) (error)
	Disconnect() (error)
	Status() (error, string)

	Create(entity map[string]interface{}) (error, string)
	Read(id interface{}) (error, map[string]interface{})
	Find(query map[string]interface{}) (error, [](map[string]interface{}))
	Update(id interface{}, entity map[string]interface{}) (error)
	Delete(id interface{}) (error)
}

type HasDatastore interface {
	GetDatastore() (Datastore)
	SetDatastore(ds Datastore)
}