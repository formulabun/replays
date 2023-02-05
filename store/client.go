package store

import (
	"encoding/hex"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"go.formulabun.club/srb2kart/lump/replay"
)

type columnScanner interface {
	Scan(dest ...any) error
}

func ConvertToStoreData(replay replay.ReplayRaw) (r Replay, files []ReplayFile, err error) {
	if len(replay.PlayerEntries) != 1 {
		err = errors.New("Replay file doesn't contain exactly 1 player")
		return
	}
	player := replay.PlayerEntries[0]
	r = Replay{
		GameMap: replay.GameMap,
		Time:    replay.Time,
		BestLap: replay.Lap,
		Speed:   player.Speed,
		Weight:  player.Weight,
	}
	copy(r.PlayerName[:], player.Name[:16])
	copy(r.PlayerSkin[:], player.Skin[:16])
	copy(r.PlayerColor[:], player.Color[:16])

	files = make([]ReplayFile, len(replay.WadEntries))
	for i, file := range replay.WadEntries {
		files[i].FileName = file.FileName
		copy(files[i].FileCheckSum[:16], file.WadMd5[:])
	}

	return r, files, nil
}

func (db *Client) InsertReplayRaw(replay replay.ReplayRaw) (id int64, err error) {
	r, files, err := ConvertToStoreData(replay)
	if err != nil {
		return 0, err
	}
	return db.InsertReplayInfo(r, files)
}

func (db *Client) InsertReplayInfo(replay Replay, files []ReplayFile) (id int64, err error) {
	t, err := db.Begin()
	defer t.Commit()
	if err != nil {
		return
	}

	res, err := t.Exec(
		`
INSERT INTO replays
(GameMap, Time, BestLap, PlayerName, PlayerSkin, PlayerColor, Speed, Weight)
VALUE
(?, ?, ?, ?, ?, ?, ?, ?)
`,
		replay.GameMap, replay.Time, replay.BestLap, fmt.Sprintf("%s", replay.PlayerName), fmt.Sprintf("%s", replay.PlayerSkin), fmt.Sprintf("%s", replay.PlayerColor), replay.Speed, replay.Weight)
	if err != nil {
		t.Rollback()
		return
	}

	id, err = res.LastInsertId()

	for _, file := range files {
    checksum := hex.EncodeToString(file.FileCheckSum[:])
		_, err = t.Exec(
			`
INSERT INTO replayfiles
(ReplayID, FileName, FileCheckSum)
VALUE
(?, ?, ?)
`,
    id, file.FileName, checksum)
		if err != nil {
			t.Rollback()
			return
		}

	}

	return
}

func scanRow(row columnScanner) (Replay, error) {
	var result Replay
	var name, skin, color string
	err := row.Scan(&result.ReplayID, &result.GameMap, &result.Time, &result.BestLap, &name, &skin, &color, &result.Speed, &result.Weight)

	copy(result.PlayerName[:], name)
	copy(result.PlayerSkin[:], skin)
	copy(result.PlayerColor[:], color)

	return result, err
}

func (db *Client) GetReplayById(id int64) (result Replay, err error) {
	row := db.QueryRow(
		`
SELECT ReplayID, GameMap, Time, BestLap, PlayerName, PlayerSkin, PlayerColor, Speed, Weight from replays 
WHERE ReplayId = ?
`, id)
	return scanRow(row)
}

func (db *Client) FindReplay(r Replay) ([]Replay, error) {
	var where = "WHERE GameMap %s ? AND Time %s ? AND BestLap %s ? AND Speed %s ? AND Weight %s ?"
	var cMap, cTime, cLap, cSpeed, cWeight string
	if r.GameMap == 0 {
		cMap = ">"
	} else {
		cMap = "="
	}
	if r.Time == 0 {
		cTime = ">"
	} else {
		cTime = "<="
	}
	if r.BestLap == 0 {
		cLap = ">"
	} else {
		cLap = "<="
	}
	if r.Speed == 0 {
		cSpeed = ">"
	} else {
		cSpeed = "="
	}
	if r.Weight == 0 {
		cWeight = ">"
	} else {
		cWeight = "="
	}

	where = fmt.Sprintf(where, cMap, cTime, cLap, cSpeed, cWeight)
	rows, err := db.Query(
		fmt.Sprintf(`SELECT * FROM replays
    %s`, where),
		r.GameMap, r.Time, r.BestLap, r.Speed, r.Weight)

	var result = []Replay{}
	if err != nil {
		return result, err
	}

	for rows.Next() {
		item, err := scanRow(rows)
		if err != nil {
			return result, err
		}
		result = append(result, item)
	}
	return result, err

}
