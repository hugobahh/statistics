package utils

import (
	"fmt"

	"strings"

	"github.com/spf13/viper"
)

//=====================  Archivo Conf ==========================================
func GetIpPuerto() string {
	sIpPort := ""
	fmt.Println(sIpPort)
	viper.SetConfigFile("configHttp.json")
	if err := viper.ReadInConfig(); err != nil {
		//panic(fmt.Errorf("Fatal error in config json file: %s ", err))
		return sIpPort
	}

	sPort := fmt.Sprintf(":%s", viper.GetString("HTTP.port"))
	sPort = strings.Replace(fmt.Sprint(sPort), ":", "", 1)
	sIp := fmt.Sprintf(":%s", viper.GetString("HTTP.ip"))
	sIp = strings.Replace(fmt.Sprint(sIp), ":", "", 1)
	if sPort == "0" || sPort == "" {
		sIpPort = fmt.Sprint(sIp)
	} else {
		sIpPort = fmt.Sprint(sIp) + ":" + fmt.Sprint(sPort)
	}
	return sIpPort
} //FIN de ObtenerIpPuerto

func ObtenerPuerto() string {
	sPort := ""
	fmt.Println(sPort)
	viper.SetConfigFile("configHttp.json")
	if err := viper.ReadInConfig(); err != nil {
		//panic(fmt.Errorf("Fatal error in config json file: %s ", err))
		return sPort
	}

	sPort = fmt.Sprintf(":%s", viper.GetString("HTTP.port"))
	return sPort
} //FIN de ObtenerPort

//=====================  Archivo Conf ==========================================
func ObtenerUrlMsSearchUsr() string {
	sUrl := ""
	viper.SetConfigFile("configMsSearchUsr.json")
	if err := viper.ReadInConfig(); err != nil {
		//panic(fmt.Errorf("Fatal error in config json file: %s ", err))
		return sUrl
	}

	sUrl = fmt.Sprintf(":%s", viper.GetString("SEARCH_USR.url"))
	sUrl = strings.Replace(fmt.Sprint(sUrl), ":", "", 1)

	return sUrl
}
