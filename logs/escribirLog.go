package logs

import (
	"log"
	"os"
)

func EscribirLineaLog(sTxt string) {
	fFile, err := os.OpenFile("App.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//log.Println("error opening file: %v", err)
		log.Println("error opening file " + sTxt)
	}
	log.SetOutput(fFile)
	defer fFile.Close()
	//wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(fFile)
	log.Output(1, sTxt)
}

func EscribirLineaErrorLog(sTxt string) {
	fFile, err := os.OpenFile("Err.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//log.Println("error opening file: %v", err)
		log.Println("error opening file " + sTxt)
	}
	log.SetOutput(fFile)
	defer fFile.Close()
	//wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(fFile)
	log.Output(1, sTxt)
}
