module go.formulabun.club/replays

go 1.19

require (
	github.com/go-sql-driver/mysql v1.7.0
	github.com/gorilla/mux v1.8.0
	go.formulabun.club/srb2kart v0.0.0-20221230121727-74d7ef86b554
)

require go.formulabun.club/functional v0.0.0-20221218180115-a1d18adcf950 // indirect

replace (
  go.formulabun.club/srb2kart => ../srb2kart
)
