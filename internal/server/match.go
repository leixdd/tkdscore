package gameserver

import (
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

type Match struct {
	ID          string
	BluePlayer  Player
	RedPlayer   Player
	Judges      map[string]Judge
	MatchStatus MatchStatusType
	TimeLimit   time.Duration
	MatchTimer  *time.Timer

	JoinChannel chan Judge
}

func NewMatch() *Match {

	timeLimit := (MAX_TIME_DURATION * time.Minute)
	mTimer := time.NewTimer(timeLimit)

	return &Match{
		ID:          ulid.Make().String(),
		Judges:      make(map[string]Judge),
		MatchStatus: MATCH_STATUS_PENDING,
		TimeLimit:   timeLimit,
		MatchTimer:  mTimer,
	}
}

func (m *Match) SetTimer(v float64) {
	m.TimeLimit = time.Duration(v * float64(time.Minute))
}

func (m *Match) ResetTimer() {
	m.MatchTimer.Reset(m.TimeLimit)
}

func (m *Match) SetRedPlayer(player Player) {
	player.PlayerType = PLAYER_RED
	m.SetRedPlayer(player)
}

func (m *Match) SetBluePlayer(player Player) {
	player.PlayerType = PLAYER_BLUE
	m.SetBluePlayer(player)
}

func (m *Match) StartSearchingForJudges() {
	go m.listenForJudgeJoiningMatch()
}

func (m *Match) CancelTheMatch() {
	m.MatchStatus = MATCH_STATUS_CANCELLED

	//send a channel event that the match is cancelled to lobby
}

func (m *Match) listenForJudgeJoiningMatch() {

	for {

		if m.MatchStatus != MATCH_STATUS_PENDING {
			fmt.Println("Match stopped searching for judges match: ", m.ID)
			break
		}

		select {

		case judge := <-m.JoinChannel:
			m.JoinMatch(judge)

		}
	}
}

func (m *Match) JoinMatch(judge Judge) {
	_, ok := m.Judges[judge.Name]

	if ok {
		fmt.Println("Judge is already existing, cannot join match: ", m.ID)
		return
	}

	m.Judges[judge.Name] = judge
	return
}
