package handler

import (
	"encoding/json"
	"github.com/nitram509/blil-web/pkg/led"
	"net/http"

	"github.com/gorilla/mux"
	"strconv"
)

func LedGetColor(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ledNr, err := strconv.Atoi(vars["ledNr"])
	if err != nil {
		ledNr = -1
	}

	if ledNr < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	aLed := &led.Led{Number: ledNr}
	if err := json.NewEncoder(w).Encode(aLed); err != nil {
		panic(err)
	}
}
