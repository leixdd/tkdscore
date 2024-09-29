package gameserver

type Lobby struct {
	Matches map[uint16]Match
}

func NewLobby() *Lobby {
	return &Lobby{
		Matches: make(map[uint16]Match),
	}
}
