package router

import (
	"volley/internal/handlers"
	"volley/internal/middlewares"

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

	api := r.Group("/api")

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
		//users.POST("/", userHandler.Create)
		//users.GET("/", userHandler.GetAll)
		users.GET("/", userHandler.Get)
		users.PUT("/", userHandler.Update)
		users.DELETE("/", userHandler.Delete)
	}

	teams := api.Group("/teams")
	teams.Use(middlewares.AuthMiddleware(db))

	{
		teams.POST("/", teamHandler.Create)
		teams.GET("/", teamHandler.GetAll)
		teams.GET("/:id", teamHandler.Get)
		teams.PUT("/", teamHandler.Update)
		teams.DELETE("/", teamHandler.Delete)
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
}
