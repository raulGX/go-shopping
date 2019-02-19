package cart

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PRODUCTCOLLECTION is the name of the collection in DB
const PRODUCTCOLLECTION = "products"

type Product struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}
type Products []Product

type MongoProductRepository struct {
	Collection *mgo.Collection
}

func NewMongoProductRepository(session *mgo.Session, c *DBConfig) *MongoProductRepository {
	collection := session.DB(c.DBName).C(PRODUCTCOLLECTION)
	// add indexes here
	return &MongoProductRepository{collection}
}

func (r *MongoProductRepository) GetProducts() (Products, error) {
	results := Products{}
	if err := r.Collection.Find(nil).All(&results); err != nil {
		return nil, errors.New("Couldn't connect to the database")
	}

	return results, nil
}

func (r *MongoProductRepository) AddProduct(p Product) error {
	err := r.Collection.Insert(p)
	if err != nil {
		return errors.New("Failed to insert product")
	}
	return nil
}
