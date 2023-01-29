package store

import (
	"database/sql"
	"fmt"
)

type Client struct {
	*sql.DB
}

type ReplayFile struct {
	FileName     string
	FileCheckSum [16]byte
}

type Replay struct {
	ReplayID    int64
	GameMap     uint16
	Time        uint32
	BestLap     uint32
	PlayerName  [16]byte
	PlayerSkin  [16]byte
	PlayerColor [16]byte
	Speed       byte
	Weight      byte
}

func makeClient(databaseName string) (Client, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("root:pass@tcp(replaydb)/%s", databaseName))
	if err != nil {
		return Client{}, err
	}
	if err = db.Ping(); err != nil {
		return Client{}, err
	}
	return Client{db}, nil
}

func NewClient() (Client, error) {
  return makeClient("replays")
}

func newTestClient() (Client, error) {
  return makeClient("testReplays")
}

func (r Replay) String() string {
	return fmt.Sprintf("Replay %d on map %d with time %d, best lap %d by %s driving a %s %s with stats (%d,%d)", r.ReplayID, r.GameMap, r.Time, r.BestLap, r.PlayerName, r.PlayerSkin, r.PlayerColor, r.Speed, r.Weight)
}
