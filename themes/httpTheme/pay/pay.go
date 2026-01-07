package pay

import (
	"encoding/json"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/k0kubun/pp"
)

var mtx sync.Mutex

func PayRequest(money *atomic.Int64) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var payment Payment
		if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		payment.Println()
		int64Paid := int64(payment.Paid)

		mtx.Lock()
		if int64Paid > money.Load() {
			w.WriteHeader(http.StatusPaymentRequired)

			errorMessage := "/pay money limit error"
			printWritingError(w, []byte(errorMessage))
			mtx.Unlock()
			return
		}

		money.Add(-int64Paid)

		paymentResponse := PaymentResponse{
			Payment:  payment,
			Time:     time.Now().UTC().Unix(),
			ActMoney: int(money.Load()),
			Sl:       []int{1, 2, 3, 4, 5},
		}

		value, err := json.MarshalIndent(paymentResponse, "", "	")
		if err != nil {
			w.WriteHeader(http.StatusPaymentRequired)

			errMessage := "Error with json Marshal: " + string(value)
			printWritingError(w, []byte(errMessage))
			mtx.Unlock()
			return
		}

		printWritingError(w, value)
		mtx.Unlock()
	}
}

func printWritingError(w http.ResponseWriter, message []byte) {
	if _, err := w.Write(message); err != nil {
		pp.Println("Error with server body: ", err)
	}
}
