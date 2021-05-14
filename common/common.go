package common

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	Host     string `json:"mysql-host"`
	User     string `json:"mysql-user"`
	Pass     string `json:"mysql-password"`
	Db       string `json:"mysql-db"`
	Token    string `json:"token"`
	HostPort string `json:"hostport"`
}

func LoadConfiguration() config {

	var config config

	configFile, err := ioutil.ReadFile("/etc/statuzpage-api/config.json")
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(configFile, &config)
	return config
}

func DBConnection() (*sql.DB, error) {

	config := LoadConfiguration()
	db, err := sql.Open("mysql", ""+config.User+":"+config.Pass+"@tcp("+config.Host+")/"+config.Db)

	return db, err
}

func CheckToken(token string) bool {
	config := LoadConfiguration()
	if token == config.Token {
		return true
	}
	return false
}

func Message(w http.ResponseWriter, message string) {
	log.Println(message)
	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
}
