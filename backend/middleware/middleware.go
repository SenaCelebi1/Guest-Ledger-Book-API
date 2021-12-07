package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	helper "bookapi/helper"

	models "bookapi/model"
)

//Connection mongoDB with helper class

func GuestLedger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := helper.ConnectDB()
	// create mails array
	var mails []models.GuestLedger
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var mail models.GuestLedger
		// & character returns the memory address of the following variable.
		err := cur.Decode(&mail) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		mails = append(mails, mail)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(mails) // encode similar to serialize process.
}

func CreateMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := helper.ConnectDB()
	var mail models.GuestLedger
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&mail)

	// insert our mail model.
	result, err := collection.InsertOne(context.TODO(), mail)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}
