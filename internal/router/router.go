package router

import (
	"volley/internal/handlers"
	"volley/internal/middlewares"
	"volley/internal/services"

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

	championshipHandler := handlers.NewChampionshipHandler(db)
	roundHandler := handlers.NewRoundHandler(db)
	userHandler := handlers.NewUserHandler(db)
	teamHandler := handlers.NewTeamHandler(db)
	gameHandler := handlers.NewGameHandler(db)
	playerHandler := handlers.NewPlayerHandler(db)
	setHandler := handlers.NewSetHandler(db)
	ampluaHandler := handlers.NewAmpluaHandler(db)
	actionHandler := handlers.NewActionHandler(db)
	actionRateHandler := handlers.NewActionRateHandler(db)
	authHandler := handlers.NewAuthHandler(db)

	setActionHandler := handlers.NewSetActionHandler(services.NewSetActionService(db))
	statsHandler := handlers.NewStatsHandler(services.NewStatsService(db))

	api := r.Group("/api")

	record := api.Group("/record")
	{
		record.POST("/action", setActionHandler.Create)
	}

	stats := api.Group("/stats")
	{
		stats.POST("/player", statsHandler.GetPlayerStats)
	}

	auth := api.Group("/")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	championships := api.Group("/championships")
	championships.Use(middlewares.AuthMiddleware(db))
	{
		championships.POST("/", championshipHandler.Create)
		championships.GET("/", championshipHandler.GetAll)
		championships.GET("/:id", championshipHandler.Get)
		championships.PUT("/:id", championshipHandler.Update)
		championships.DELETE("/:id", championshipHandler.Delete)
	}

	rounds := api.Group("/rounds")
	rounds.Use(middlewares.AuthMiddleware(db))
	{
		rounds.POST("/", roundHandler.Create)
		rounds.GET("/", roundHandler.GetAll)
		rounds.GET("/:id", roundHandler.Get)
		rounds.PUT("/:id", roundHandler.Update)
		rounds.DELETE("/:id", roundHandler.Delete)
	}

	users := api.Group("/users")
	users.Use(middlewares.AuthMiddleware(db))
	{
		users.POST("/", userHandler.Create)
		users.GET("/", userHandler.GetAll)
		users.GET("/:id", userHandler.Get)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}

	teams := api.Group("/teams")
	teams.Use(middlewares.AuthMiddleware(db))

	{
		teams.POST("/", teamHandler.Create)
		teams.GET("/", teamHandler.GetAll)
		teams.GET("/:id", teamHandler.Get)
		teams.PUT("/:id", teamHandler.Update)
		teams.DELETE("/:id", teamHandler.Delete)
	}

	players := api.Group("/players")
	players.Use(middlewares.AuthMiddleware(db))

	{
		players.POST("/", playerHandler.Create)
		players.GET("/", playerHandler.GetAll)
		players.GET("/:id", playerHandler.Get)
		players.PUT("/:id", playerHandler.Update)
		players.DELETE("/:id", playerHandler.Delete)
	}

	sets := api.Group("/sets")
	sets.Use(middlewares.AuthMiddleware(db))

	{
		sets.POST("/", setHandler.Create)
		sets.GET("/", setHandler.GetAll)
		sets.GET("/:id", setHandler.Get)
		sets.PUT("/:id", setHandler.Update)
		sets.DELETE("/:id", setHandler.Delete)
	}

	games := api.Group("/games")
	games.Use(middlewares.AuthMiddleware(db))

	{
		games.POST("/", gameHandler.Create)
		games.GET("/", gameHandler.GetAll)
		games.GET("/:id", gameHandler.Get)
		games.PUT("/:id", gameHandler.Update)
		games.DELETE("/:id", gameHandler.Delete)
	}

	ampluas := api.Group("/ampluas")
	ampluas.Use(middlewares.AuthMiddleware(db))

	{
		ampluas.POST("/", ampluaHandler.Create)
		ampluas.GET("/", ampluaHandler.GetAll)
		ampluas.GET("/:id", ampluaHandler.Get)
		ampluas.PUT("/:id", ampluaHandler.Update)
		ampluas.DELETE("/:id", ampluaHandler.Delete)
	}

	actions := api.Group("/actions")
	actions.Use(middlewares.AuthMiddleware(db))

	{
		actions.POST("/", actionHandler.Create)
		actions.GET("/", actionHandler.GetAll)
		actions.GET("/:id", actionHandler.Get)
		actions.PUT("/:id", actionHandler.Update)
		actions.DELETE("/:id", actionHandler.Delete)
	}

	actionRates := api.Group("/action_rates")
	actionRates.Use(middlewares.AuthMiddleware(db))
	{
		actionRates.POST("/", actionRateHandler.Create)
		actionRates.GET("/", actionRateHandler.GetAll)
		actionRates.GET("/:id", actionRateHandler.Get)
		actionRates.PUT("/:id", actionRateHandler.Update)
		actionRates.DELETE("/:id", actionRateHandler.Delete)
	}

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
