FROM golang:1.18 AS build
WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o replayinfo server/main.go

FROM scratch AS runtime
COPY --from=build /go/src/replayinfo ./
EXPOSE 8080/tcp
ENTRYPOINT ["./replayinfo"]
