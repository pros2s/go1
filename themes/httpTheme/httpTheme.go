package httptheme

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"go1/themes/httpTheme/pay"
	"go1/themes/httpTheme/query"
)

var (
	PORT  = ":9000"
	money atomic.Int64
)

func HTTPTest() {
	money.Add(1000)

	http.HandleFunc("/pay", pay.PayRequest(&money))
	http.HandleFunc("/query", query.TestQuery)

	fmt.Printf("Server on http://localhost%s\n", PORT)

	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
