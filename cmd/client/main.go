package main

import (
	"github.com/leixddd/jtkdscore/internal/client"
	tcpclient "github.com/leixddd/jtkdscore/internal/tcp/client"
)

func main() {

	tcpClient := tcpclient.NewTCPClient(":3111")
	tcpClient.Connect()

	defer tcpClient.ConnState.Close()

	gameClient := client.NewGameClient(tcpClient)
	gameClient.Init()

}
