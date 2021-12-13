package middleware

import (
	"log"
	"net/http"
	"os"
)

func LoggerMW(next http.Handler) http.Handler {

	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {

		f, err := os.OpenFile("log/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			log.Fatal("Error opening log file: %d", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(req.Method, req.URL, req.RemoteAddr)

		next.ServeHTTP(wr, req)
	})
}
