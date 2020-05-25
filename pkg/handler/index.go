package handler

import (
	"encoding/json"
	"github.com/nitram509/blil-web/pkg/info"
	"github.com/nitram509/blil-web/pkg/led"
	"net/http"
	"strconv"
)

type link struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}

type selfLink struct {
	Self link `json:"self"`
}

type ledInfoResource struct {
	led.LedInfo            //embedded
	Links       []selfLink `json:"_links"`
}

type embedded struct {
	Led []ledInfoResource `json:"leds"`
}

type indexResource struct {
	Version  string   `json:"version"`
	Name     string   `json:"name"`
	Embedded embedded `json:"_embedded"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	embeddedleds := embedded{
		Led: make([]ledInfoResource, 0),
	}

	leds := led.DetectAllLeds().Leds
	for l := range leds {
		aLed := leds[l]
		lir := ledInfoResource{
			Links: make([]selfLink, 0),
		}
		lir.Links = append(lir.Links, selfLink{Self: link{
			Href:  "http://" + r.Host + "/led/" + strconv.Itoa(aLed.Number),
			Title: "Set or get color on this LED",
		}})
		lir.Number = aLed.Number
		lir.Path = aLed.Path
		lir.Type = aLed.Type
		embeddedleds.Led = append(embeddedleds.Led, lir)
	}

	index := &indexResource{
		Version:  info.VERSION,
		Name:     info.NAME,
		Embedded: embeddedleds,
	}

	if err := json.NewEncoder(w).Encode(index); err != nil {
		panic(err)
	}
}
