package gorest

import (
	// "errors"
)

const (
	PropertyInt = "int"
	PropertyString = "string"
	PropertyBool = "bool"
	PropertyFloat = "float"
)

type Property struct {
	Name string
	Type string
	Required bool
	MinLength uint
	MaxLength uint
}

type Model interface {
	HasName
	HasLogger

	GetProperties() (map[string]Property)
	SetProperties(properties map[string]Property) (error)

	IsValidCreate(values map[string]interface{}) (bool, map[string]string)
	IsValidRead(id interface{}) (bool, string)
	IsValidFind(query map[string]interface{}) (bool, map[string]string)
	IsValidUpdate(id interface{}, values map[string]interface{}) (bool, map[string]string)
	IsValidDelete(id interface{}) (bool, string)
}

type HasModel interface {
	GetModel() (Model)
	SetModel(model Model)
}