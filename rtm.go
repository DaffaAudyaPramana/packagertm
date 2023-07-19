package packagertm

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataAccounts(db *mongo.Database, nama, email, sosial string, perusahaan string) (InsertedID interface{}) {
	var dataaccounts Accounts
	dataaccounts.Nama = nama
	dataaccounts.Email = email
	dataaccounts.Sosial = sosial
	dataaccounts.Perusahaan = perusahaan
	return InsertOneDoc(db, "data_accounts", dataaccounts)
}

func GetDataAccounts(perusahaan string, db *mongo.Database, col string) (data Accounts) {
	act := db.Collection(col)
	filter := bson.M{"perusahaan": perusahaan}
	err := act.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdataaccbyact: %v\n", err)
	}
	return data
}
func GetDataNama(nama string, db *mongo.Database, col string) (data Accounts) {
	accou := db.Collection(col)
	filter := bson.M{"nama": nama}
	err := accou.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdataaccount: %v\n", err)
	}
	return data
}
func DeleteDataAccounts(perusahaan string, db *mongo.Database, col string) (data Accounts) {
	dct := db.Collection(col)
	filter := bson.M{"perusahaan": perusahaan}
	err, _ := dct.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataAccounts : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}

func DeleteDataNama(nama string, db *mongo.Database, col string) (data Accounts) {
	dena := db.Collection(col)
	filter := bson.M{"nama": nama}
	err, _ := dena.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataNama : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}
