package packagertm

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "Accounts",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertData(t *testing.T) {
	Nama := "Daffa"
	Email := "Daffaganteng@gmail.com"
	Sosial := "Daffanigaes"
	Perusahaan := "Daffakomputer"
	hasil := InsertDataAccounts(MongoConn, Nama, Email, Sosial, Perusahaan)
	fmt.Println(hasil)

}

func TestGetDataAccounts(t *testing.T) {
	Nama := "Daffa"
	hasil := GetDataAccounts(Nama, MongoConn, "data_accounts")
	fmt.Println(hasil)

}

func TestDeleteData(t *testing.T) {
	Perusahaan := "Daffakomputer"
	hasil := DeleteDataAccounts(Perusahaan, MongoConn, "data_accounts")
	fmt.Println(hasil)

}
