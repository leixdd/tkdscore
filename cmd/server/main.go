package main

import (
	gameserver "github.com/leixddd/jtkdscore/internal/server"
	tcpserver "github.com/leixddd/jtkdscore/internal/tcp/server"
)

func main() {
	tcpServer := tcpserver.NewTCPServer(":3111")
	gameServer := gameserver.NewGameServer()

	gameServer.GameLobby = gameserver.NewLobby()
	gameServer.TCPServer = tcpServer

	go gameServer.ListenEvents()
	go gameServer.TCPServer.Start()

}
