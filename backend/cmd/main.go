package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"cpsu/internal/config"
	"cpsu/internal/connectdb"
	"cpsu/internal/handler"
	"cpsu/internal/repository"
	"cpsu/internal/service"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := connectdb.NewPostgresDatabase(cfg.GetConnectionString())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	cpsuRepo := repository.NewCPSURepository(db.GetDB())
	cpsuService := service.NewCPSUService(cpsuRepo)
	cpsuHandler := handler.NewCPSUHandler(cpsuService)

	go func() {
		for {
			time.Sleep(10 * time.Second)
			if err := db.Ping(); err != nil {
				log.Printf("Database connection lost: %v", err)
				if reconnErr := db.Reconnect(cfg.GetConnectionString()); reconnErr != nil {
					log.Printf("Failed to reconnect: %v", reconnErr)
				} else {
					log.Printf("Successfully reconnected to the database")
				}
			}
		}
	}()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(TimeoutMiddleware(5 * time.Second))

	r.GET("/health", func(c *gin.Context) {
		if err := connectdb.CheckDBConnection(db.GetDB()); err != nil {
			c.JSON(503, gin.H{"detail": "Database connection failed"})
			return
		}
		c.JSON(200, gin.H{"status": "healthy", "database": "connected"})
	})

	v1 := r.Group("/api/v1")
	{
		news := v1.Group("admin/news")
		{
			news.GET("", cpsuHandler.GetAllNews)
			news.GET("/:id", cpsuHandler.GetNewsDetail)
			news.POST("", cpsuHandler.CreateNews)
			news.PUT("/:id", cpsuHandler.UpdateNews)
			news.DELETE("/:id", cpsuHandler.DeleteNews)
		}
	}

	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
