package db

import "sync"

type Datasource struct {
	Values      *sync.Map
	Expirations *sync.Map
}

func New() *Datasource {
	return &Datasource{
		Values:      &sync.Map{},
		Expirations: &sync.Map{},
	}
}
