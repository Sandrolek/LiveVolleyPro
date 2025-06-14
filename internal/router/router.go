package router

import (
	"volley/internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {

	// userRepo := repositories.NewUserRepository(db)
	// userHandler := handlers.NewUserHandler(userRepo)
	// playerHandler := handlers.NewPlayerHandler(db)
	// teamHandler := handlers.NewTeamHandler(db)
	// gameHandler := handlers.NewGameHandler(db)
	// roundHandler := handlers.NewRoundHandler(db)
	// champHandler := handlers.NewChampionshipHandler(db)
	// gamePlayerHandler := handlers.NewGamePlayerHandler(db)
	// setHandler := handlers.NewSetHandler(db)
	// setActionHandler := handlers.NewSetActionHandler(db)

	api := r.Group("/api")

	championshipHandler := handlers.NewChampionshipHandler(db)
	roundHandler := handlers.NewRoundHandler(db)

	api.POST("/championships", championshipHandler.Create)
	api.GET("/championships", championshipHandler.GetAll)
	api.GET("/championships/:id", championshipHandler.Get)
	api.PUT("/championships/:id", championshipHandler.Update)
	api.DELETE("/championships/:id", championshipHandler.Delete)

	api.POST("/rounds", roundHandler.Create)
	api.GET("/rounds", roundHandler.GetAll)
	api.GET("/rounds/:id", roundHandler.Get)
	api.PUT("/rounds/:id", roundHandler.Update)
	api.DELETE("/rounds/:id", roundHandler.Delete)

	// api := r.Group("/api/v1")
	// {
	// users := api.Group("/users")
	// {
	// 	users.GET("/", userHandler.GetAllUsers)
	// 	users.GET("/:user_id", userHandler.GetUserByID)
	// 	users.POST("/", userHandler.CreateUser)
	// 	users.PUT("/:user_id", userHandler.UpdateUser)
	// 	users.DELETE("/:user_id", userHandler.DeleteUser)
	// }

	// teams := api.Group("/teams")
	// {
	// 	teams.POST("/", teamHandler.CreateTeam)
	// 	teams.GET("/", teamHandler.GetAllTeams)
	// 	teams.GET("/:id", teamHandler.GetTeamByID)
	// 	teams.PUT("/:id", teamHandler.UpdateTeam)
	// 	teams.DELETE("/:id", teamHandler.DeleteTeam)
	// }

	// players := api.Group("/players")
	// {
	// 	players.POST("/", playerHandler.CreatePlayer)
	// 	players.GET("/", playerHandler.GetAllPlayers)
	// 	players.GET("/:id", playerHandler.GetPlayerByID)
	// 	players.PUT("/:id", playerHandler.UpdatePlayer)
	// 	players.DELETE("/:id", playerHandler.DeletePlayer)
	// }

	// games := r.Group("/games")
	// {
	// 	games.POST("", gameHandler.CreateGame)
	// 	games.GET("", gameHandler.GetAllGames)
	// 	games.GET("/:id", gameHandler.GetGameByID)
	// 	games.PUT("/:id", gameHandler.UpdateGame)
	// 	games.DELETE("/:id", gameHandler.DeleteGame)
	// }

	// rounds := r.Group("/rounds")
	// {
	// 	rounds.POST("", roundHandler.CreateRound)
	// 	rounds.GET("", roundHandler.GetAllRounds)
	// 	rounds.GET("/:id", roundHandler.GetRoundByID)
	// 	rounds.PUT("/:id", roundHandler.UpdateRound)
	// 	rounds.DELETE("/:id", roundHandler.DeleteRound)
	// }

	// champs := r.Group("/championships")
	// {
	// 	champs.POST("", championshipHandler.Create)
	// 	champs.GET("", championshipHandler.GetAll)
	// 	champs.GET("/:id", championshipHandler.Get)
	// 	champs.PUT("/:id", championshipHandler.Update)
	// 	champs.DELETE("/:id", championshipHandler.Delete)
	// }

	// gamePlayers := r.Group("/game_players")
	// {
	// 	gamePlayers.POST("", gamePlayerHandler.CreateGamePlayer)
	// 	gamePlayers.GET("", gamePlayerHandler.GetAllGamePlayers)
	// 	gamePlayers.GET("/:id", gamePlayerHandler.GetGamePlayerByID)
	// 	gamePlayers.PUT("/:id", gamePlayerHandler.UpdateGamePlayer)
	// 	gamePlayers.DELETE("/:id", gamePlayerHandler.DeleteGamePlayer)
	// }

	// sets := r.Group("/sets")
	// {
	// 	sets.POST("", setHandler.CreateSet)
	// 	sets.GET("", setHandler.GetAllSets)
	// 	sets.GET("/:id", setHandler.GetSetByID)
	// 	sets.PUT("/:id", setHandler.UpdateSet)
	// 	sets.DELETE("/:id", setHandler.DeleteSet)
	// }

	// setActions := r.Group("/set_actions")
	// {
	// 	setActions.POST("", setActionHandler.CreateSetAction)
	// 	setActions.GET("", setActionHandler.GetAllSetActions)
	// 	setActions.GET("/:id", setActionHandler.GetSetActionByID)
	// 	setActions.PUT("/:id", setActionHandler.UpdateSetAction)
	// 	setActions.DELETE("/:id", setActionHandler.DeleteSetAction)
	// }
	// }
}
