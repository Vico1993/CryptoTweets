package database

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "portfolio"

// Portfolio all item saved
type Portfolio struct {
	collection *mongo.Collection
	Items      []PortfolioItem
}

// NewPorfolio Instantiate a new portfolio with default db connection
func NewPorfolio() (Portfolio, error) {
	database, err := getDatabase()
	if err != nil {
		return Portfolio{}, err
	}

	return Portfolio{collection: database.Collection(collectionName)}, err
}

func (p *Portfolio) get(tag string) (PortfolioItem, error) {
	var item PortfolioItem

	err := p.collection.FindOne(context.TODO(), bson.M{"tag": tag}).Decode(&item)
	if err != nil {
		return PortfolioItem{}, err
	}

	return item, nil
}

func (p *Portfolio) update(tag string, item PortfolioItem) (PortfolioItem, error) {
	result, err := p.collection.UpdateOne(context.TODO(), bson.M{"tag": tag}, item)
	if err != nil {
		return PortfolioItem{}, err
	}

	if result.ModifiedCount > 0 {
		return item, nil
	}

	return PortfolioItem{}, errors.New("No item found for tag " + tag)
}

func (p *Portfolio) add(items []PortfolioItem) (bool, error) {
	for _, item := range items {
		_, err := p.collection.InsertOne(context.TODO(), item)
		if err != nil {
			log.Fatal(err)
		}
	}

	return true, nil
}

func (p Portfolio) delete(tag string) (bool, error) {
	result, err := p.collection.DeleteOne(context.TODO(), bson.M{"tag": tag})
	if err != nil {
		return false, err
	}

	if result.DeletedCount > 0 {
		return true, nil
	}

	return false, errors.New("No item found for tag " + tag)
}
