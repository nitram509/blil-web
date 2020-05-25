package handler

import (
	"encoding/json"
	"fmt"
	"github.com/nitram509/blil-web/pkg/led"
	"net/http"

	"github.com/gorilla/mux"
	"strconv"
)

func LedSetColor(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ledNr, err := strconv.Atoi(vars["ledNr"])
	if err != nil {
		ledNr = -1
	}
	col := MapColor(vars["color"])

	if ledNr < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if col == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	led.SetLedColor(ledNr, col)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	r, g, b, _ := col.RGBA()
	aLed := &led.Led{Number: ledNr, Color: fmt.Sprintf("%.2x%.2x%.2x", uint8(r), uint8(g), uint8(b))}
	if err := json.NewEncoder(w).Encode(aLed); err != nil {
		panic(err)
	}
}
