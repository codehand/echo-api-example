package db

import (
	"fmt"
	"time"

	"github.com/codehand/echo-restful-crud-api-example/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getNextID(collection *mgo.Collection, name string) int {
	result := bson.M{}
	if _, err := collection.Find(bson.M{"_id": name}).Apply(mgo.Change{
		Update:    bson.M{"$set": bson.M{"_id": name}, "$inc": bson.M{"seq": 1}},
		Upsert:    true,
		ReturnNew: true,
	}, &result); err != nil {
		fmt.Println("Autoincrement error(1):", err.Error())
	}
	sec, _ := result["seq"].(int)
	return sec
}

// CreateNewProduct is func adapter save record under database
func CreateNewProduct(in *types.Product) (*types.Product, error) {
	sessionCopy := pullSession()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(databaseName).C("product")
	nextID := getNextID(collection, "product")
	in.ID.Hex()
	in.ProductID = fmt.Sprintf("%d", nextID)
	in.CreatedAt = time.Now()
	in.UpdatedAt = in.CreatedAt
	err := collection.Insert(in)
	if err != nil {
		return nil, err
	}
	return in, nil
}

// GetAllProducts is func get all products
func GetAllProducts() ([]*types.Product, error) {
	sessionCopy := pullSession()
	defer sessionCopy.Close()
	var data []*types.Product
	err := sessionCopy.DB(databaseName).C("product").Find(bson.M{"deleted_at": nil, "_id": bson.M{"$ne": "product"}}).All(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetProduct is func get one product
func GetProduct(id int) (*types.Product, error) {
	sessionCopy := pullSession()
	defer sessionCopy.Close()
	var data *types.Product
	err := sessionCopy.DB(databaseName).C("product").Find(bson.M{"deleted_at": nil, "product_id": fmt.Sprintf("%d", id), "_id": bson.M{"$ne": "product"}}).One(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateProduct is func update one product
func UpdateProduct(id int, in *types.ProductUpdate) (*types.Product, error) {
	sessionCopy := pullSession()
	defer sessionCopy.Close()
	_, err := GetProduct(id)
	if err != nil {
		return nil, err
	}
	update := buildConditions(*in)
	if len(update) > 0 {
		condition := bson.M{"deleted_at": nil, "product_id": fmt.Sprintf("%d", id), "_id": bson.M{"$ne": "product"}}
		err = sessionCopy.DB(databaseName).C("product").Update(condition, bson.M{"$set": update})
		if err != nil {
			return nil, err
		}
		data, err := GetProduct(id)
		if err != nil {
			return nil, err
		}

		return data, nil
	}
	return nil, fmt.Errorf("%v", "Not update any fields")
}

// DeleteProduct is func delete one product
func DeleteProduct(id int) (*types.Product, error) {
	sessionCopy := pullSession()
	defer sessionCopy.Close()
	data, err := GetProduct(id)
	if err != nil {
		return nil, err
	}
	condition := bson.M{"deleted_at": nil, "product_id": fmt.Sprintf("%d", id), "_id": bson.M{"$ne": "product"}}

	err = sessionCopy.DB(databaseName).C("product").Remove(condition)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// DeleteAtProduct is func delete one product
func DeleteAtProduct(id int) (*types.Product, error) {
	sessionCopy := pullSession()
	defer sessionCopy.Close()
	data, err := GetProduct(id)
	if err != nil {
		return nil, err
	}
	t := time.Now()
	condition := bson.M{"deleted_at": nil, "product_id": fmt.Sprintf("%d", id), "_id": bson.M{"$ne": "product"}}

	err = sessionCopy.DB(databaseName).C("product").Update(condition, bson.M{"$set": bson.M{"deleted_at": t}})
	if err != nil {
		return nil, err
	}
	data.DeletedAt = &t
	return data, nil
}
