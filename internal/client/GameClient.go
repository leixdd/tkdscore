package client

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/leixddd/jtkdscore/_common"
	tcpclient "github.com/leixddd/jtkdscore/internal/tcp/client"
)

type GameClient struct {
	input       chan rune
	event_input chan []byte
	tcpClient   *tcpclient.TCPClient
	User        string
	TargetRoom  string
}

func NewGameClient(tcpClient *tcpclient.TCPClient) *GameClient {

	return &GameClient{
		input:       make(chan rune),
		tcpClient:   tcpClient,
		event_input: make(chan []byte),
	}
}

func (g *GameClient) Init() {

	defer func() {
		g.tcpClient.ConnState.Close()
		close(g.event_input)
		close(g.input)
	}()

	go g.gameLoop()
	g.Setup()
}

func (g *GameClient) Setup() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Judge Name: ")
	v, err := r.ReadString('\n')

	v = strings.TrimSpace(v)

	if err != nil {
		slog.Error("unable to read", "err", err)
		return
	}

	g.User = v

	cls()

	opts := make(map[uint8]string)

	opts[1] = "Connect to match server"
	opts[2] = "Quit"

	for i, v := range opts {
		fmt.Printf("[%d] %s \n", i, v)
	}

	fmt.Print("Choice: ")

	c, err := r.ReadString('\n')

	if err != nil {
		slog.Error("unable to read", "err", err)
		return
	}

	c = strings.TrimSpace(c)

	switch c {
	case "1":
		cls()
		fmt.Print("Room Name: ")

		room, err := r.ReadString('\n')

		if err != nil {
			slog.Error("unable to read", "err", err)
			return
		}

		g.TargetRoom = strings.TrimSpace(room)
		evtParam := []byte(fmt.Sprintf("%s:%s", g.TargetRoom, g.User))

		serverEventParam := []byte{
			_common.SERVER_EVENT_JOIN_ROOM,
		}

		serverEventParam = append(serverEventParam, evtParam...)

		g.event_input <- serverEventParam

		g.inputLoop()

	case "2":
		return
	default:
		cls()
		g.Setup()
	}
}

func cls() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *GameClient) inputLoop() {
	fmt.Println("Connected to Match: ", g.TargetRoom)
	keyEvents, err := keyboard.GetKeys(10)

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		ev := <-keyEvents

		if ev.Err != nil {
			panic(ev.Err)
		}

		g.input <- ev.Rune

		if ev.Key == keyboard.KeyEsc {
			break
		}

	}
}

func (g *GameClient) HandleServerMessage(message []byte) {
	fmt.Println(string(message), "hatdog")
}

func (g *GameClient) gameLoop() {

	for {
		select {
		case key := <-g.input:
			g.tcpClient.Writer([]byte{byte(key)})
		case eventMessage := <-g.event_input:
			g.tcpClient.Writer(eventMessage)
		case serverMessage := <-g.tcpClient.ServerMessages:
			g.HandleServerMessage(serverMessage)

		}
	}

}
