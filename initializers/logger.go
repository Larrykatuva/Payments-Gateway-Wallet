package initializers

import (
	"log"
	"os"
	"strconv"
	"time"
)

var (
	WarningLog *log.Logger
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
)

func StartLogger() {
	year, month, day := time.Now().Date()
	date := strconv.Itoa(year) + "/" + month.String() + "/" + strconv.Itoa(day)
	logFile, err := os.OpenFile(date+"-logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logToFile, _ := strconv.ParseBool(os.Getenv("LogToFile"))
	if logToFile {
		log.SetOutput(logFile)
	}
	InfoLog = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(logFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
