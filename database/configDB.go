package database

import (
	"database/sql"
	"fmt"
	"test_yo/logs"

	"log"

	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type DatosDB struct {
	ip   string `json:"ip"`
	port string `json:"port"`
	usr  string `json:"usr"`
	pwd  string `json:"pwd"`
	db   string `json:"db"`
}

func CnnDB() (db *sql.DB, e error) {
	viper.SetConfigFile("configDB.json")
	if err := viper.ReadInConfig(); err != nil {
		logs.EscribirLineaLog("CnnDB_Err: " + err.Error())
		panic(fmt.Errorf("Fatal error in config file: %s ", err))
		return
	}
	sIP := fmt.Sprintf(":%s", viper.GetString("DB.IP"))
	sIP = strings.Replace(fmt.Sprint(sIP), ":", "", 1)

	sPort := fmt.Sprintf(":%s", viper.GetString("DB.port"))
	sPort = strings.Replace(fmt.Sprint(sPort), ":", "", 1)

	sUsr := fmt.Sprintf(":%s", viper.GetString("DB.usr"))
	sUsr = strings.Replace(fmt.Sprint(sUsr), ":", "", 1)

	sPwd := fmt.Sprintf(":%s", viper.GetString("DB.pwd"))
	sPwd = strings.Replace(fmt.Sprint(sPwd), ":", "", 1)

	sDB := fmt.Sprintf(":%s", viper.GetString("DB.basedatos"))
	sDB = strings.Replace(fmt.Sprint(sDB), ":", "", 1)

	db, err := sql.Open("mysql", sUsr+":"+sPwd+"@tcp("+sIP+":"+sPort+")/"+sDB+"?parseTime=true")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return db, nil
}

