package core

import (
	"errors"
	"github.com/google/uuid"
	"slices"
)

type Player struct {
	ID       uuid.UUID
	Nickname string
}

func NewPlayer(nickname string) *Player {
	return &Player{
		ID:       uuid.New(),
		Nickname: nickname,
	}
}

func (p *Player) JoinGame(game *Game) error {
	if len(game.Players) >= 2 {
		return errors.New("game already full")
	} else if slices.Contains(game.Players, *p) {
		return errors.New("player already connected")
	}
	game.Players = append(game.Players, *p)
	return nil
}
