package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
	"test_yo/data"
	"test_yo/logs"
	"test_yo/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

//=======================================================
type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type Monto struct {
	Invesment float64 `json:"investment" binding:"required,numeric"`
}
type Credit struct {
	n700 int32
	n500 int32
	n300 int32
	Err  error
}

//=======================================================
func CreditAssigment(c *gin.Context) {
	logs.EscribirLineaLog("CreditAssigment ...")
	sIdUsr := ""
	blnCall := false
	fmt.Println(blnCall)
	sToken := ""

	//Leer Json
	var monto Monto
	err := c.ShouldBind(&monto)

	if err != nil {
		c.JSON(400, gin.H{"code": "api_error", "message": "Formato de json no valido.", "status_code": 400})
		return
	}
	//Token con el formato XXXXX.XXXXX
	if len(c.Request.Header["Authorization"]) > 0 {
		sToken = c.Request.Header["Authorization"][0]
		if sToken == "Bearer" {
			c.JSON(400, gin.H{"code": "search_user_microservice", "message": "Se requiere de un token.", "status_code": 400})
			return
		}
		sToken = strings.Replace(sToken, "Bearer", "", -1)
		sToken = strings.Trim(sToken, " ")
		sIdUsr, blnCall = callMsSearchUser(sToken)
		fmt.Println(sIdUsr)
		if blnCall == false {
			c.JSON(400, gin.H{"code": "api_error", "message": "Token no valido.", "status_code": 400})
			return
		}
	} else {
		c.JSON(400, gin.H{"code": "api_error", "message": "Se requiere de un token.", "status_code": 400})
		return
	}
	//Validations
	if monto.Invesment < 300 {
		c.JSON(400, gin.H{"code": "api_error", "message": "La inversion minima es de $300, igual o mayor a $400 pesos.", "status_code": 400})
		return
	}
	//No aceptar decimales
	sTmp := fmt.Sprintf("%v", monto.Invesment)
	if strings.Contains(sTmp, ".") == true {
		c.JSON(400, gin.H{"code": "api_error", "message": "No se permiten montos con decimales.", "status_code": 400})
		return
	}

	//=======================================================
	//Genera creditos y guarda los procesados
	blnOK, err, sTmp := getCredit(sIdUsr, int32(monto.Invesment), Credit{})
	fmt.Println(err, blnOK)
	if err != nil {
		c.JSON(400, gin.H{"code": "api_error", "message": err.Error(), "status_code": 400})
		return
	} else {
		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(sTmp), &jsonMap)
		jsonData, _ := json.Marshal(jsonMap)
		c.Data(200, "application/json", jsonData)
		return
	}
	return
} //FIN de CreditAssigment

func Calculous7(investment int32) (int32, int32, int32, error) {
	z := math.Mod(float64(investment), 700.00)
	nCont := (investment / 700.00)
	if z == 0 {
		return int32(nCont), 0, 0, nil
	} else if z == 100 {
		nCont = (nCont - 1)
		return int32(nCont), 1, 1, nil
	} else if z == 200 {
		nCont = (nCont - 1)
		return int32(nCont), 0, 3, nil
	} else if z == 300 {
		nCont = (nCont - 1)
		return int32(nCont), 2, 0, nil
	} else if z == 400 {
		nCont = (nCont - 1)
		return int32(nCont), 5, 2, nil
	} else if z == 500 {
		nCont = (nCont - 1)
		return int32(nCont), 0, 4, nil
	} else if z == 600 {
		nCont = (nCont - 1)
		return int32(nCont), 5, 1, nil
	} else {
		err := errors.New("Monto no aceptado, deben ser multiplos de 100.")
		return 0, 0, 0, err
	}
	return 0, 0, 0, nil
}

func Calculous(investment int32) (int32, int32, int32, error) {
	z := math.Mod(float64(investment), 300.00)
	nCont := (investment / 300.00)
	if z == 0 {
		return int32(nCont), 0, 0, nil
	} else if z == 100 {
		if nCont >= 2 {
			nCont = (nCont - 2)
			return 1, 0, int32(nCont), nil
		} else {
			err := errors.New("Saldo de 100, no se puede procesar el credito.")
			return 1, 0, int32(nCont), err
		}
	} else if z == 200 {
		nCont = (nCont - 1)
		return 0, 1, int32(nCont), nil
	} else {
		err := errors.New("Monto no aceptado, deben ser multiplos de 100.")
		return 0, 0, int32(nCont), err
	}
	return 0, 0, 0, nil
}

func (c Credit) Assign(investment int32) (int32, int32, int32, error) {
	c.n700, c.n500, c.n300, c.Err = Calculous(investment)
	return c.n700, c.n500, c.n300, c.Err
}

func getCredit(IdUsr string, nInvesment int32, credit CreditAssigner) (bool, error, string) {
	sTmp := ""
	fmt.Println(sTmp)
	n700, n500, n300, err := credit.Assign(nInvesment)
	if err != nil {
		blnReg := data.RegisterCredit(IdUsr, fmt.Sprint(nInvesment), fmt.Sprint(n700), fmt.Sprint(n500), fmt.Sprint(n300), "0")
		if blnReg == true {
			sTmp = `{"code": "api_error", "message": ` + err.Error() + `, "status_code": 400}`
			return true, err, sTmp
		}
		return false, err, ""
	} else {
		//Guardar los datos
		//s700 := strconv.FormatInt(int64(n700), 10)
		blnData := data.RegisterCredit("1", fmt.Sprint(nInvesment), fmt.Sprint(n700), fmt.Sprint(n500), fmt.Sprint(n300), "1")
		if blnData == true {
			sTmp = `{"credit_type_300":` + fmt.Sprint(n300) + `, "credit_type_500":` + fmt.Sprint(n500) + `, "credit_type_700":` + fmt.Sprint(n700) + `}`
			return true, err, sTmp
		} else {
			sTmp = `{"code": "api_error", "message": "No fue posible registrar lel crédito.", "status_code": 400}`
			return false, err, sTmp
		}
	}
}

//=======================================================
type ResultMs struct {
	Code        string `json:"code"`
	Msg         string `json:"message"`
	Status_code int32  `json:"status_code"`
}

func callMsSearchUser(sToken string) (string, bool) {
	sUrl := utils.ObtenerUrlMsSearchUsr()

	var resMs ResultMs
	//var resError ResultError

	// Create a Resty Client
	client := resty.New()
	resOK, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(sToken).
		//SetBody(sJson).
		SetResult(&resMs).
		//SetError(&resError).
		Get(sUrl)
	fmt.Println(resOK)
	if err != nil {
		logs.EscribirLineaLog("msSearchUser_Err: " + err.Error())
		fmt.Println("Error: " + err.Error())
		return "", false
	}
	//sTmp := resOK
	fmt.Println(resOK)
	if resMs.Status_code == 200 {
		return resMs.Msg, true
	} else {
		return "", false
	}
	return "", true
}
