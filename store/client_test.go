package store

import (
	"testing"

	"go.formulabun.club/srb2kart/lump/replay"
)

func TestInsertAndGet(t *testing.T) {
	client, err := newTestClient()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	r := replay.ReplayRaw{}
	player := replay.PlayerEntryData{Speed: 4, Weight: 1}
	copy(player.Name[:], []byte("Fl_GUI"))
	copy(player.Skin[:], []byte("cream"))
	copy(player.Color[:], []byte("lemonade"))
	r.PlayerEntries = replay.PlayerEntries{replay.PlayerEntry{0, player}}
	r.WadEntries = replay.WadEntries{replay.WadEntry{"bonusfiles.kart", [16]byte{}}}

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

func TestMultipleSearch(t *testing.T) {
	names := [][16]byte{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{41, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{42, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{43, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	dataset := []Replay{
		Replay{0, 1, 1500, 500, names[1], names[1], names[1], 1, 1},
		Replay{0, 1, 1000, 400, names[1], names[1], names[1], 1, 1},
		Replay{0, 1, 1300, 400, names[2], names[1], names[1], 1, 1},
		Replay{0, 1, 1700, 600, names[2], names[1], names[1], 1, 1},
		Replay{0, 2, 1500, 500, names[1], names[1], names[1], 1, 1},
		Replay{0, 2, 1700, 700, names[1], names[1], names[1], 1, 1},
		Replay{0, 3, 1500, 500, names[1], names[1], names[1], 1, 1},
		Replay{0, 4, 1500, 500, names[1], names[1], names[1], 1, 1},
	}
	testData := []struct {
		data     Replay
		expected []int
	}{
		{Replay{0, 1, 0, 0, names[0], names[0], names[0], 0, 0}, []int{0,2}},
		{Replay{0, 2, 0, 0, names[0], names[0], names[0], 0, 0}, []int{4}},
	}

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

	for _, data := range dataset {
		_, err := client.InsertReplayInfo(data, []ReplayFile{})
		if err != nil {
			t.Fatal(err)
		}
	}

	for i, test := range testData {
		replays, err := client.FindReplay(test.data)
		if err != nil {
			t.Fatal(err)
		}
		if len(replays) != len(test.expected) {
			t.Errorf("test nr %v: expected %v replays but got %v", i, len(test.expected), len(replays))
      t.Log("Gotten:")
			for _, r := range replays {
				t.Log(r)
			}
		}
	}
}
