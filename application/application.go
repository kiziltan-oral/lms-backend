package application

import (
	"log"
	"net/http"

	"lms-web-services-main/database/datasources"
	"lms-web-services-main/repositories"
	"lms-web-services-main/routers"
	services "lms-web-services-main/services"

	"github.com/LGYtech/lgo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// Load Environment Variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading environment: %v", err)
		return
	}
	log.Println(".env file loaded successfully")

	// Setup CORS
	setupCORS()

	// Setup Router
	addRoutes()

	// Start Server
	log.Println("LMS Service is running on port 8080")
	router.Run(":8080")
}

func setupCORS() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"} // Update as per your frontend's origin
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Token"}
	router.Use(cors.New(corsConfig))
}

func addRoutes() {
	// #region Initialize repositories and services
	systemUserRepo := repositories.NewSystemUserRepository(datasources.Database)
	systemUserService := services.NewSystemUserService(systemUserRepo)

	systemUserSettingRepo := repositories.NewSystemUserSettingRepository(datasources.Database)
	systemUserSettingService := services.NewSystemUserSettingService(systemUserSettingRepo)

	clientRepo := repositories.NewClientRepository(datasources.Database)
	clientService := services.NewClientService(clientRepo)

	clientProjectRepo := repositories.NewClientProjectRepository(datasources.Database)
	clientProjectService := services.NewClientProjectService(clientProjectRepo)

	timingRepo := repositories.NewTimingRepository(datasources.Database)
	timingService := services.NewTimingService(timingRepo)

	// #endregion Initialize repositories and services

	// #region Add Routes
	// Korumasız rotalar
	openRoutes := router.Group("/")
	routers.NonProtectedRoutes(openRoutes, systemUserService)

	// Korunan rotalar
	protectedRoutes := router.Group("/")
	protectedRoutes.Use(authenticationMiddleware())
	routers.SystemUserRoutes(protectedRoutes, systemUserService)
	routers.SystemUserSettingRoutes(protectedRoutes, systemUserSettingService)
	routers.ClientRoutes(protectedRoutes, clientService)
	routers.ClientProjectRoutes(protectedRoutes, clientProjectService)
	routers.TimingRoutes(protectedRoutes, timingService)
	// #endregion Add Routes
}

func authenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// #region Get User Token
		userToken := c.GetHeader("X-Token")
		log.Println("Gelen Token:", userToken)

		if len(userToken) == 0 {
			log.Println("HATA: Token Header boş.")
			c.JSON(http.StatusUnauthorized, lgo.NewAuthError())
			c.Abort()
			return
		}

		authResult := services.CacheService.AuthenticateSystemUser(userToken)
		if authResult == nil || !authResult.IsSuccess() {
			log.Println("HATA: AuthenticateSystemUser başarısız.")
			c.JSON(http.StatusUnauthorized, lgo.NewAuthError())
			c.Abort()
			return
		}

		if !authResult.ReturnObject.(bool) {
			log.Println("HATA: Token geçersiz.")
			c.JSON(http.StatusUnauthorized, lgo.NewAuthError())
			c.Abort()
			return
		}
		// #endregion Get User Token

		// #region Set Context
		c.Set("usertoken", userToken)
		log.Println("Başarılı Token:", userToken)
		// #endregion Set Context

		c.Next()
	}
}

// #endregion Middleware
