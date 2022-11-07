package data

import (
	"fmt"
	"strconv"
	"test_yo/database"
	"test_yo/logs"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func RegisterCredit(sIdUsr string, nInv string, c700 string, c500 string, c300 string, sSaldo string) bool {
	//logs.EscribirLineaLog("RegisterOrder ...")
	var nId int64 = 0
	fmt.Println(nId)

	//"paid_at": "2022-08-23T00:00:00" para obtener segundos
	dtFechaHoy := time.Now()
	sTmp := dtFechaHoy.String()
	sFA := sTmp[0:19]
	sUS := strconv.Itoa(int(time.Now().Unix()))

	//nDatePay := dTmp.Unix()
	db, err := database.CnnDB()
	if err != nil {
		//fmt.Printf("Error obteniendo base de datos: %v", err)
		logs.EscribirLineaLog("RegisterCredit_Error cnn_DB " + err.Error())
		return false
	}
	err = db.Ping()
	if err != nil {
		logs.EscribirLineaLog("RegisterCredit_Err: " + err.Error())
		return false
	} else {
		defer db.Close()
		//Registrar los creditos
		sSQL := "INSERT INTO credit(id_user, c700, c500, c300, fecha_alta, unix_stamp, ok, investment) "
		sSQL += "VALUES(" + sIdUsr + ", " + c700 + ", " + c500 + ", " + c300 + ", '"
		sSQL += sFA + "', " + sUS + ", "
		if sSaldo != "0" {
			sSQL += "1, " + nInv + ")"
		} else {
			sSQL += "0, " + nInv + ")"
		}

		/*sSQL := `INSERT INTO credit(id_user, c700, c500, c300, fecha_alta, unix_stamp, ok) `
		sSQL += `VALUES(?, ?, ?, ?, ?, ?, ?),`
		sSQL += sIdUsr + `,` + c700 + `,` + c500 + `,` + c300 + `, ` + sFA + `, ` + sUS + `, `
		sParam := sIdUsr + `,` + c700 + `,` + c500 + `,` + c300 + `, ` + sFA + `, ` + sUS + `, `
		if sSaldo != "0" {
			sSQL += `1)`
			sParam += `1`
		} else {
			sSQL += `0)`
			sParam += `0`
		}
		resDB, err := db.Query("INSERT INTO credit(id_user, c700, c500, c300, fecha_alta, unix_stamp, ok) VALUES(?, ?, ?, ?, ?, ?, ?) " + sParam + " ) ")
		*/
		resDB, err := db.Query(sSQL)
		fmt.Println(resDB)
		if err != nil {
			logs.EscribirLineaLog("SQL: " + sSQL)
			logs.EscribirLineaLog("RegisterCreditErr: No fue posible registrar order: " + err.Error())
			return false
		} else {
			return true
		}
	}
} //FIN de RegisterOrder
