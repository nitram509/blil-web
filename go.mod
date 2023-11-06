module github.com/nitram509/blil-web

go 1.14

require (
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d // indirect
	github.com/boombuler/hid v0.0.0-20200303134931-8ff92ccd15a8 // indirect
	github.com/boombuler/led v0.0.0-20190225062837-d94ba02fda02
	github.com/gorilla/mux v1.8.1
	gopkg.in/alecthomas/kingpin.v1 v1.3.7
)

// replace because build error: go build github.com/boombuler/hid: invalid flag in #cgo LDFLAGS: -fconstant-cfstrings
replace github.com/boombuler/hid => github.com/nitram509/hid v0.0.0-20200516220657-af2d63ac9fab
