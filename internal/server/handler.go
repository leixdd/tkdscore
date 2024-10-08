package gameserver

import (
	"context"

	"github.com/leixddd/jtkdscore/_common"
	tcpserver "github.com/leixddd/jtkdscore/internal/tcp/server"
)

type PlayerType = uint8
type ScoreTypeValue = uint8
type ScoreTypeName = string
type MatchStatusType = uint8

const (
	PLAYER_BLUE PlayerType = 1
	PLAYER_RED  PlayerType = 2

	SCORE_NAME_PUNCH     ScoreTypeName = "PUNCH"
	SCORE_NAME_BODY_KICK ScoreTypeName = "BODY"
	SCORE_NAME_HEAD_KICK ScoreTypeName = "HEAD"

	SCORE_PUNCH     ScoreTypeValue = 1
	SCORE_BODY_KICK ScoreTypeValue = 2
	SCORE_HEAD_KICK ScoreTypeValue = 3

	PENALTY_DEDUCTION_COST uint8 = 1

	MATCH_STATUS_PENDING   MatchStatusType = 0
	MATCH_STATUS_READY     MatchStatusType = 1
	MATCH_STATUS_ONGOING   MatchStatusType = 2
	MATCH_STATUS_COMPLETED MatchStatusType = 3
	MATCH_STATUS_CANCELLED MatchStatusType = 4

	MAX_JUDGES        uint8 = 3
	MAX_TIME_DURATION       = 2
)

type Player struct {
	ID         string
	Scores     []Score
	Name       string
	Penalties  uint8
	PlayerType PlayerType
}

type Score struct {
	Player Player
	Judge  Judge
	Type   string
	Value  ScoreTypeValue
}

type ScoreHistory struct {
	Score     Score
	Timestamp uint64
	PlayerID  string
}

type GameServer struct {
	GameLobby         *Lobby
	GameServerContext context.Context
	TCPServer         *tcpserver.TCPServer
}

type GameServerEvent struct {
	EventType _common.GameServerEventType
	Data      string
}

func NewGameServer() *GameServer {
	return &GameServer{
		GameServerContext: context.Background(),
	}
}

func (gs *GameServer) ListenEvents() {
	for {
		select {
		case eventMessage := <-gs.TCPServer.ClientMessages:
			gs.handleTCPCommands(eventMessage)
		}
	}
}

func (gs *GameServer) handleTCPCommands(clientMessage *tcpserver.ClientMessage) {

	size := len(clientMessage.Message)

	if size == 0 {
		return
	}

	switch clientMessage.Message[0] {
	case _common.SERVER_EVENT_JOIN_ROOM:
		payload := string(clientMessage.Message[1:size])

	}

}
