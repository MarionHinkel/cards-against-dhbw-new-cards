package shelf

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/yannickkirschen/cards-against-dhbw/data"
	"github.com/yannickkirschen/cards-against-dhbw/game"
	"github.com/yannickkirschen/cards-against-dhbw/play"
	"github.com/yannickkirschen/cards-against-dhbw/utils"
)

var GlobalShelf *GameShelf

var ErrNotFound = errors.New("game does not exist")

type GameShelf struct {
	games map[string]*play.Play
	r     *rand.Rand
}

func New() *GameShelf {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	return &GameShelf{games: make(map[string]*play.Play), r: r}
}

func (gs *GameShelf) CreateGame(name string) (string, string, error) {
	log.Printf("Player %s wants to create a new game!", name)

	gameId := gs.newGameId()
	log.Printf("Created a new game ID for the game player %s wants to create. ID: %s", name, gameId)

	blacks, whites, err := data.ReadCards()
	if err != nil {
		return "", "", err
	}

	p := play.New(game.New(gameId, blacks, whites))
	player, _ := p.Game.CreatePlayer(name)
	p.Game.Mod = player

	gs.games[gameId] = p
	return gameId, player.Name, nil
}

func (gs *GameShelf) Play(id string) (*play.Play, error) {
	gp, exists := gs.games[id]
	if exists {
		return gp, nil
	}

	return nil, ErrNotFound
}

func (gs *GameShelf) JoinGame(gameId string, name string) (string, error) {
	log.Printf("Player %s wants to join game %s!", name, gameId)

	game, ok := gs.games[gameId]
	if ok {
		player, err := game.Game.CreatePlayer(name)
		if err != nil {
			return "", err
		}

		return player.Name, nil
	}

	return "", ErrNotFound
}

func (gs *GameShelf) newGameId() string {
	gameId := utils.RandString(gs.r, 4)
	for {
		_, exists := gs.games[gameId]

		if exists {
			gameId = utils.RandString(gs.r, 4)
		} else {
			return gameId
		}
	}
}