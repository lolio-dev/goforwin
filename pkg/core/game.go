package core

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"slices"
)

type Grid [][]string

type Game struct {
	Grid    Grid
	Players []Player
}

type PawnPosition struct{ C, R int }

func (g *Game) Lauch() {
	fmt.Println("Lauching game...")
}

func NewGame() *Game {
	game := &Game{
		Grid:    GenerateGrid(),
		Players: []Player{},
	}

	return game
}

func (g *Game) PlacePawn(col int, player *Player) error {
	if col > 7 {
		return errors.New("game has only 7 columns")
	}
	if !slices.Contains(g.Players, *player) {
		return errors.New("player not in game")
	}

	for i := range g.Grid {
		i = len(g.Grid) - 1 - i
		if g.Grid[i][col] == "_" {
			g.Grid[i][col] = player.ID.String()
			return nil
		}
	}

	return errors.New("column already full")
}

func (g *Game) CheckWin() (*uuid.UUID, error) {
	for ri, r := range g.Grid {
		for ci, p := range r {
			pawn := &PawnPosition{ci, ri}
			if g.GetPositionNeighbours(pawn) == nil {
				pUuid, err := uuid.Parse(p)
				if err != nil {
					return nil, errors.New("win but can't parse uuid")
				}
				return &pUuid, nil
			}
		}
		fmt.Println()
	}

	return nil, nil
}

func (g *Game) GetPositionNeighbours(position *PawnPosition) []PawnPosition {
	// vérifer dans les 9 cases autour si il y a des pions appartenant au même joueur
	// Prendre chaque pion
	// 		augmenter le conteur de 1
	// 		Réappeller la fonction pour le pion
	// Si il y a encore un pion de la même couleur
	// sauf celui précédent augmenter le conteur de 1 et continuer
	var pawnNeighbours []PawnPosition

	for _, c := range []int{position.C - 1, position.C, position.C + 1} {
		for _, r := range []int{position.R - 1, position.R, position.R + 1} {
			newNeighbour := &PawnPosition{c, r}
			if !reflect.DeepEqual(newNeighbour, position) && c != -1 && c != 7 && r != -1 && r != 6 {
				pawnNeighbours = append(pawnNeighbours, PawnPosition{c, r})
			}
		}
	}

	return pawnNeighbours
}

func (g *Game) CheckPawnWin(pawn *PawnPosition) (bool, error) {
	playerUUID := g.Grid[pawn.R][pawn.C]

	if playerUUID == "_" {
		return false, errors.New("no pawn in this position")
	}

	// Check row
	if LongestConsecutiveOccurrenceLength[string](g.Grid[pawn.R], playerUUID) == 4 {
		return true, nil
	}

	// Check Col
	var col []string
	for _, r := range g.Grid {
		col = append(col, r[pawn.C])
	}
	if LongestConsecutiveOccurrenceLength[string](col, playerUUID) == 4 {
		return true, nil
	}

	// check diagonals
	var firstDiag []string
	for i := range []int{0, 1, 2, 3, 4, 5} {
		if i == 0 {
			firstDiag = append(firstDiag, g.Grid[pawn.R][pawn.C])
			continue
		}
		if pawn.R-i >= 0 && pawn.C+i <= 6 {
			firstDiag = append(firstDiag, g.Grid[pawn.R-i][pawn.C+i])
		}
		if pawn.R+i <= 5 && pawn.C-i >= 0 {
			firstDiag = append(firstDiag, g.Grid[pawn.R+i][pawn.C-i])
		}
	}
	if LongestConsecutiveOccurrenceLength[string](firstDiag, playerUUID) == 4 {
		return true, nil
	}

	var secondDiag []string
	for i := range []int{0, 1, 2, 3, 4, 5} {
		if i == 0 {
			secondDiag = append(secondDiag, g.Grid[pawn.R][pawn.C])
			continue
		}
		if pawn.R-i >= 0 && pawn.C-i >= 0 {
			secondDiag = append(secondDiag, g.Grid[pawn.R-i][pawn.C-i])
		}
		if pawn.R+i <= 5 && pawn.C+i <= 6 {
			secondDiag = append(secondDiag, g.Grid[pawn.R+i][pawn.C+i])
		}
	}
	if LongestConsecutiveOccurrenceLength[string](secondDiag, playerUUID) == 4 {
		return true, nil
	}

	return false, nil
}

func LongestConsecutiveOccurrenceLength[T comparable](arr []T, e T) int {
	count := 0
	maxCount := 0
	for _, tabE := range arr {
		if tabE == e {
			count += 1
		} else {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		}
	}

	return maxCount
}

func GenerateGrid() Grid {
	grid := make([][]string, 6)
	for i := range grid {
		grid[i] = []string{"_", "_", "_", "_", "_", "_", "_"}
	}
	return grid
}
