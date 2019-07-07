package db

import (
	"fmt"
	"reflect"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func userModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// buildConditions ..
func buildConditions(rs interface{}) bson.M {
	rt := reflect.TypeOf(rs)
	rv := reflect.ValueOf(rs)
	nv := rv.NumField()

	nosql := make(map[string]interface{})
	for i := 0; i < nv; i++ {
		if rv.Field(i) != reflect.Zero(reflect.TypeOf(rv.Field(i))).Interface() {
			if rt.Field(i).Tag.Get("nosql") == "" || fmt.Sprintf("%v", rv.Field(i)) == "" {
				continue
			}
			nosql[rt.Field(i).Tag.Get("nosql")] = fmt.Sprintf("%v", rv.Field(i))
		}
	}
	nosql["updated_at"] = time.Now()
	return nosql
}
