package gorest

import (
	// "errors"
)

type Presenter interface {
	HasName
	HasLogger

	Parse(data *[]byte) (int, []map[string]interface{})
	Serialize([]map[string]interface{}) (data *[]byte)
}

type HasPresenter interface {
	GetPresenter() (Presenter)
	SetPresenter(presenter Presenter)
}