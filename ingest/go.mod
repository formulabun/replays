module go.formulabun.club/replays/ingest

go 1.19

require (
	github.com/gorilla/mux v1.8.0
	golang.org/x/oauth2 v0.3.0
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	go.formulabun.club/functional v0.0.0-20221218180115-a1d18adcf950 // indirect
	go.formulabun.club/replays/store v0.0.0-00010101000000-000000000000 // indirect
	go.formulabun.club/srb2kart v0.0.0-20221230121727-74d7ef86b554 // indirect
	golang.org/x/net v0.3.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace go.formulabun.club/replays/store => ../store
