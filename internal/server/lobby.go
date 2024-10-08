package gameserver

import (
	"fmt"
	"math/rand/v2"
)

type Lobby struct {
	Matches map[uint]*Match
}

func NewLobby() *Lobby {
	return &Lobby{
		Matches: make(map[uint]*Match),
	}
}

func (lobby *Lobby) CreateNewMatch() uint {

	m := NewMatch()

	initRand := func() uint {
		MAX_UINT16 := 999999
		MIN_UINT := 100000

		return rand.UintN(uint(MAX_UINT16)-uint(MIN_UINT)) + uint(MIN_UINT)
	}

	curSeed := initRand()

	_, isExists := lobby.Matches[curSeed]

	for isExists {
		curSeed = initRand()
		_, isExists = lobby.Matches[curSeed]
	}

	lobby.Matches[curSeed] = m

	return curSeed
}

func (Lobby *Lobby) JoinMatch(id uint, cornerJudge string) *Match {

	match, isExists := Lobby.Matches[id]

	if !isExists {
		fmt.Println("Match is not existing")
		return nil
	}

	return match
}
