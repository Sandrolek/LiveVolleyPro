package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"volley/internal/handlers"
	"volley/internal/repositories"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)
	playerHandler := handlers.NewPlayerHandler(db)
	teamHandler := handlers.NewTeamHandler(db)
	gameHandler := handlers.NewGameHandler(db)
	roundHandler := handlers.NewRoundHandler(db)
	champHandler := handlers.NewChampionshipHandler(db)
	gamePlayerHandler := handlers.NewGamePlayerHandler(db)
	setHandler := handlers.NewSetHandler(db)
	setActionHandler := handlers.NewSetActionHandler(db)

	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("/", userHandler.GetAllUsers)
			users.GET("/:user_id", userHandler.GetUserByID)
			users.POST("/", userHandler.CreateUser)
			users.PUT("/:user_id", userHandler.UpdateUser)
			users.DELETE("/:user_id", userHandler.DeleteUser)
		}

		teams := api.Group("/teams")
		{
			teams.POST("/", teamHandler.CreateTeam)
			teams.GET("/", teamHandler.GetAllTeams)
			teams.GET("/:id", teamHandler.GetTeamByID)
			teams.PUT("/:id", teamHandler.UpdateTeam)
			teams.DELETE("/:id", teamHandler.DeleteTeam)
		}

		players := api.Group("/players")
		{
			players.POST("/", playerHandler.CreatePlayer)
			players.GET("/", playerHandler.GetAllPlayers)
			players.GET("/:id", playerHandler.GetPlayerByID)
			players.PUT("/:id", playerHandler.UpdatePlayer)
			players.DELETE("/:id", playerHandler.DeletePlayer)
		}

		games := r.Group("/games")
		{
			games.POST("", gameHandler.CreateGame)
			games.GET("", gameHandler.GetAllGames)
			games.GET("/:id", gameHandler.GetGameByID)
			games.PUT("/:id", gameHandler.UpdateGame)
			games.DELETE("/:id", gameHandler.DeleteGame)
		}

		rounds := r.Group("/rounds")
		{
			rounds.POST("", roundHandler.CreateRound)
			rounds.GET("", roundHandler.GetAllRounds)
			rounds.GET("/:id", roundHandler.GetRoundByID)
			rounds.PUT("/:id", roundHandler.UpdateRound)
			rounds.DELETE("/:id", roundHandler.DeleteRound)
		}

		champs := r.Group("/championships")
		{
			champs.POST("", champHandler.CreateChampionship)
			champs.GET("", champHandler.GetAllChampionships)
			champs.GET("/:id", champHandler.GetChampionshipByID)
			champs.PUT("/:id", champHandler.UpdateChampionship)
			champs.DELETE("/:id", champHandler.DeleteChampionship)
		}

		gamePlayers := r.Group("/game_players")
		{
			gamePlayers.POST("", gamePlayerHandler.CreateGamePlayer)
			gamePlayers.GET("", gamePlayerHandler.GetAllGamePlayers)
			gamePlayers.GET("/:id", gamePlayerHandler.GetGamePlayerByID)
			gamePlayers.PUT("/:id", gamePlayerHandler.UpdateGamePlayer)
			gamePlayers.DELETE("/:id", gamePlayerHandler.DeleteGamePlayer)
		}

		sets := r.Group("/sets")
		{
			sets.POST("", setHandler.CreateSet)
			sets.GET("", setHandler.GetAllSets)
			sets.GET("/:id", setHandler.GetSetByID)
			sets.PUT("/:id", setHandler.UpdateSet)
			sets.DELETE("/:id", setHandler.DeleteSet)
		}

		setActions := r.Group("/set_actions")
		{
			setActions.POST("", setActionHandler.CreateSetAction)
			setActions.GET("", setActionHandler.GetAllSetActions)
			setActions.GET("/:id", setActionHandler.GetSetActionByID)
			setActions.PUT("/:id", setActionHandler.UpdateSetAction)
			setActions.DELETE("/:id", setActionHandler.DeleteSetAction)
		}
	}

	return r
}
