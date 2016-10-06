package migrationtest

import (
	"google.golang.org/appengine/datastore"
)

type (
	EntityV1 struct {
		_kind string `goon:"kind,Entity"`
		ID    int64  `datastore:"-" goon:"id"`
	}
	EntityV2 struct {
		_kind string `goon:"kind,Entity"`
		ID    int64  `datastore:"-" goon:"id"`
		Name  string `datastore:"name"`
	}
)

func (e *EntityV1) Load(p []datastore.Property) error {
	err := datastore.LoadStruct(e, p)
	if fmerr, ok := err.(*datastore.ErrFieldMismatch); ok && fmerr != nil && fmerr.Reason == "no such struct field" {
		// ignore
	} else if err != nil {
		return err
	}

	return nil
}

func (e *EntityV1) Save() ([]datastore.Property, error) {
	p, err := datastore.SaveStruct(e)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (e *EntityV2) Load(p []datastore.Property) error {
	err := datastore.LoadStruct(e, p)
	if fmerr, ok := err.(*datastore.ErrFieldMismatch); ok && fmerr != nil && fmerr.Reason == "no such struct field" {
		// ignore
	} else if err != nil {
		return err
	}

	return nil
}

func (e *EntityV2) Save() ([]datastore.Property, error) {
	p, err := datastore.SaveStruct(e)
	if err != nil {
		return nil, err
	}
	return p, nil
}
