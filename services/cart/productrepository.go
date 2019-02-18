package cart

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SERVER the DB server
// TODO inject this in the controller
// Maybe mock for in memory
const SERVER = "mongodb://localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "store"

// COLLECTION is the name of the collection in DB
const PRODUCTCOLLECTION = "products"

type Product struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}
type Products []Product

type ProductRepository struct{}

func (r ProductRepository) GetProducts() (Products, error) {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		return nil, errors.New("Couldn't connect to the database")
	}

	defer session.Close()
	c := session.DB(DBNAME).C(PRODUCTCOLLECTION)

	results := Products{}
	if err := c.Find(nil).All(&results); err != nil {
		return nil, errors.New("Couldn't connect to the database")
	}

	return results, nil
}

func (r ProductRepository) AddProduct(p Product) error {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		return errors.New("Couldn't connect to the database")
	}

	defer session.Close()
	c := session.DB(DBNAME).C(PRODUCTCOLLECTION)
	// if p.ID == "" {
	// 	p.ID = bson.NewObjectId()
	// }
	err = c.Insert(p)
	if err != nil {
		return errors.New("Failed to insert product")
	}
	return nil
}
