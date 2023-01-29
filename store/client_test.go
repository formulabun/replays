package store

import (
	"testing"

	"go.formulabun.club/srb2kart/lump"
)

func TestInsertAndGet(t *testing.T) {
	client, err := newTestClient()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	r := lump.ReplayRaw{}
	player := lump.PlayerEntryData{Speed: 4, Weight: 1}
	copy(player.Name[:], []byte("Fl_GUI"))
	copy(player.Skin[:], []byte("cream"))
	copy(player.Color[:], []byte("lemonade"))
	r.PlayerEntries = lump.PlayerEntries{lump.PlayerEntry{0, player}}
	r.WadEntries = lump.WadEntries{lump.WadEntry{"bonusfiles.kart", [16]byte{}}}

	id, err := client.InsertReplayRaw(r)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := client.GetReplayById(id)
	if err != nil {
		t.Fatal(err)
	}

	if id != ret.ReplayID {
		t.Error("Returned id is not the same")
	}
	t.Log(ret)
}

func TestSearch(t *testing.T) {
	testData := []struct {
		insert Replay
		search Replay
	}{
		{Replay{0, 1, 300, 100, [16]byte{}, [16]byte{}, [16]byte{}, 4, 1},
			Replay{0, 1, 0, 0, [16]byte{}, [16]byte{}, [16]byte{}, 0, 0}},
		{Replay{0, 1, 300, 100, [16]byte{}, [16]byte{}, [16]byte{}, 4, 1},
			Replay{0, 0, 500, 0, [16]byte{}, [16]byte{}, [16]byte{}, 0, 0}},
		{Replay{0, 1, 300, 100, [16]byte{}, [16]byte{}, [16]byte{}, 4, 1},
			Replay{0, 0, 0, 200, [16]byte{}, [16]byte{}, [16]byte{}, 0, 0}},
		{Replay{0, 1, 300, 100, [16]byte{}, [16]byte{}, [16]byte{}, 4, 1},
			Replay{0, 0, 0, 0, [16]byte{}, [16]byte{}, [16]byte{}, 4, 0}},
		{Replay{0, 1, 300, 100, [16]byte{}, [16]byte{}, [16]byte{}, 4, 1},
			Replay{0, 0, 0, 0, [16]byte{}, [16]byte{}, [16]byte{}, 0, 1}},
		{Replay{0, 1, 300, 100, [16]byte{}, [16]byte{}, [16]byte{}, 4, 1},
			Replay{0, 1, 500, 0, [16]byte{}, [16]byte{}, [16]byte{}, 4, 1}},
	}
	for _, data := range testData {
		client, err := newTestClient()
		if err != nil {
			t.Fatal(err)
		}
		defer client.Close()
		_, err = client.Exec(`delete from replayfiles`)
		if err != nil {
			t.Fatal(err)
		}
		_, err = client.Exec(`delete from replays`)
		if err != nil {
			t.Fatal(err)
		}

		id, _ := client.InsertReplayInfo(data.insert, []ReplayFile{})

		replays, err := client.FindReplay(data.search)
		if err != nil {
			t.Fatal(err)
		}

		if len(replays) != 1 {
			t.Fatal("Expected only one result")
		}
		if id != replays[0].ReplayID {
			t.Fatal("Unexpected return id")
		}
	}
}
