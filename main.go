package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// struct
type Game struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Size        string    `json:"size"`
	Platforms   []string  `json:"platforms"`
	ReleaseDate time.Time `json:"realease_date"`
	Publisher   string    `json:"publisher"`
	Description string    `json:"description"`
	Ratings     Rating    `json:"rating"`
}

type Rating struct {
	CriticRating float64 `json:"critic_rating"`
	UserRating   float64 `json:"user_rating"`
}

var games = []Game{
	{
		ID:          1,
		Title:       "Super Mario Odyssey",
		Size:        "5.3 GB",
		Platforms:   []string{"Nintendo Switch"},
		ReleaseDate: time.Date(2017, time.October, 27, 0, 0, 0, 0, time.UTC),
		Publisher:   "Nintendo",
		Description: "Super Mario Odyssey is a 3D platformer...",
		Ratings: Rating{
			CriticRating: 9.5,
			UserRating:   9.2,
		},
	},
	{
		ID:          2,
		Title:       "The Legend of Zelda: Breath of the Wild",
		Size:        "14.4 GB",
		Platforms:   []string{"Nintendo Switch"},
		ReleaseDate: time.Date(2017, time.March, 3, 0, 0, 0, 0, time.UTC),
		Publisher:   "Nintendo",
		Description: "The Legend of Zelda: Breath of the Wild...",
		Ratings: Rating{
			CriticRating: 9.7,
			UserRating:   9.4,
		},
	},
	{
		ID:          3,
		Title:       "Red Dead Redemption 2",
		Size:        "99 GB",
		Platforms:   []string{"PlayStation 4", "Xbox One"},
		ReleaseDate: time.Date(2018, time.October, 26, 0, 0, 0, 0, time.UTC),
		Publisher:   "Rockstar Games",
		Description: "Red Dead Redemption 2 is an action-adventure game...",
		Ratings: Rating{
			CriticRating: 9.8,
			UserRating:   9.1,
		},
	},
}

// GET
func getGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games)
}

func getGameById(id int) (*Game, error) {
	for i, g := range games {
		if g.ID == id {
			return &games[i], nil
		}
	}

	return nil, errors.New("game not found")
}

func getGame(c *gin.Context) {
	id := c.Param("id")
	gameID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	game, err := getGameById(gameID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, game)
}

// POST
func createGame(c *gin.Context) {
	var newGame Game

	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	games = append(games, newGame)
	c.IndentedJSON(http.StatusOK, games)
}

// PUT
func updateGame(c *gin.Context) {
	id := c.Param("id")
	gameID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// find game id
	index := -1
	for i, g := range games {
		if g.ID == gameID {
			index = i
			break
		}
	}

	// if not found
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	gameToUpdate := &games[index]

	// parse data json
	var updatedGame Game
	if err := c.ShouldBindJSON(&updatedGame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game data"})
		return
	}

	// Update atribut-atribut game
	gameToUpdate.Title = updatedGame.Title
	gameToUpdate.Size = updatedGame.Size
	gameToUpdate.Platforms = updatedGame.Platforms
	gameToUpdate.ReleaseDate = updatedGame.ReleaseDate
	gameToUpdate.Publisher = updatedGame.Publisher
	gameToUpdate.Description = updatedGame.Description
	gameToUpdate.Ratings = updatedGame.Ratings

	c.IndentedJSON(http.StatusOK, gameToUpdate)
}

// DELETE
func deleteGame(c *gin.Context) {
	id := c.Param("id")
	gameID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// find game
	index := -1
	for i, g := range games {
		if g.ID == gameID {
			index = i
			break
		}
	}

	// if not found
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	// slice game
	games = append(games[:index], games[index+1:]...)

	c.JSON(http.StatusOK, gin.H{"message": "Game deleted"})
}

func main() {
	router := gin.Default()
	router.GET("/games", getGames)
	router.GET("/games/:id", getGame)
	router.POST("/games", createGame)
	router.PUT("/games/:id", updateGame)
	router.DELETE("/games/:id", deleteGame)
	router.Run()
}
