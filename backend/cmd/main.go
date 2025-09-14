package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"cpsu/internal/config"
	"cpsu/internal/connectdb"

	newsHandler "cpsu/internal/news/handler"
	newsRepo "cpsu/internal/news/repository"
	newsService "cpsu/internal/news/service"

	courseHandler "cpsu/internal/course/handler"
	courseRepo "cpsu/internal/course/repository"
	courseService "cpsu/internal/course/service"

	structureHandler "cpsu/internal/course_structure/handler"
	structureRepo "cpsu/internal/course_structure/repository"
	structureService "cpsu/internal/course_structure/service"

	roadmapHandler "cpsu/internal/roadmap/handler"
	roadmapRepo "cpsu/internal/roadmap/repository"
	roadmapService "cpsu/internal/roadmap/service"

	subjectHandler "cpsu/internal/subject/handler"
	subjectRepo "cpsu/internal/subject/repository"
	subjectService "cpsu/internal/subject/service"
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

	newsRepo := newsRepo.NewNewsRepository(db.GetDB())
	newsService := newsService.NewNewsService(newsRepo, cfg.AWSRegion, cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, cfg.S3BucketName)
	newsHandler := newsHandler.NewNewsHandler(newsService)

	courseRepo := courseRepo.NewCourseRepository(db.GetDB())
	courseService := courseService.NewCourseService(courseRepo)
	courseHandler := courseHandler.NewCourseHandler(courseService)

	structureRepo := structureRepo.NewCourseStructureRepository(db.GetDB())
	structureService := structureService.NewCourseStructureService(structureRepo, cfg.AWSRegion, cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, cfg.S3BucketName)
	structureHandler := structureHandler.NewCourseStructureHandler(structureService)

	roadmapRepo := roadmapRepo.NewRoadmapRepository(db.GetDB())
	roadmapService := roadmapService.NewRoadmapService(roadmapRepo, cfg.AWSRegion, cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, cfg.S3BucketName)
	roadmapHandler := roadmapHandler.NewRoadmapHandler(roadmapService)

	subjectRepo := subjectRepo.NewSubjectRepository(db.GetDB())
	subjectService := subjectService.NewSubjectService(subjectRepo)
	subjectHandler := subjectHandler.NewSubjectHandler(subjectService)

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
		newsAdmin := v1.Group("admin/news")
		{
			newsAdmin.GET("", newsHandler.GetAllNews)
			newsAdmin.GET("/:id", newsHandler.GetNewsByID)
			newsAdmin.POST("", newsHandler.CreateNews)
			newsAdmin.PUT("/:id", newsHandler.UpdateNews)
			newsAdmin.DELETE("/:id", newsHandler.DeleteNews)
		}

		courseAdmin := v1.Group("admin/course")
		{
			courseAdmin.GET("", courseHandler.GetAllCourses)
			courseAdmin.GET("/:id", courseHandler.GetCourseByID)
			courseAdmin.POST("", courseHandler.CreateCourse)
			courseAdmin.PUT("/:id", courseHandler.UpdateCourse)
			courseAdmin.DELETE("/:id", courseHandler.DeleteCourse)
		}

		structureAdmin := v1.Group("admin/structure")
		{
			structureAdmin.GET("", structureHandler.GetAllCourseStructure)
			structureAdmin.GET("/:id", structureHandler.GetCourseStructureByID)
			structureAdmin.POST("", structureHandler.CreateCourseStructure)
			structureAdmin.DELETE("/:id", structureHandler.DeleteCourseStructure)
		}

		roadmapAdmin := v1.Group("admin/roadmap")
		{
			roadmapAdmin.GET("", roadmapHandler.GetAllRoadmap)
			roadmapAdmin.GET("/:id", roadmapHandler.GetRoadmapByID)
			roadmapAdmin.POST("", roadmapHandler.CreateRoadmap)
			roadmapAdmin.DELETE("/:id", roadmapHandler.DeleteRoadmap)
		}

		subjectAdmin := v1.Group("admin/subject")
		{
			subjectAdmin.GET("", subjectHandler.GetAllSubjects)
			subjectAdmin.GET("/:id", subjectHandler.GetSubjectByID)
			subjectAdmin.POST("", subjectHandler.CreateSubject)
			subjectAdmin.PUT("/:id", subjectHandler.UpdateSubject)
			subjectAdmin.DELETE("/:id", subjectHandler.DeleteSubject)
		}
	}

	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
