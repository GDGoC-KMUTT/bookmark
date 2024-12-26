package routes

import (
	"backend/internals/config"
	"backend/internals/controllers"
	"backend/internals/db"
	"backend/internals/entities/response"
	"backend/internals/repositories"
	"backend/internals/routes/handler"
	"backend/internals/routes/middleware"
	"backend/internals/services"
	services2 "backend/internals/utils/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
)

func SetupRoutes() {
	// * Repositories
	var userRepo = repositories.NewUserRepository(db.Gorm)
	var courseRepo = repositories.NewCourseRepository(db.Gorm)
	var stepEvalRepo = repositories.NewStepEvaluateRepository(db.Gorm)
	var userEvalRepo = repositories.NewUserEvaluateRepo(db.Gorm)
	var stepCommentRepo = repositories.NewStepCommentRepository(db.Gorm)
	var stepCommentUpVoteRepo = repositories.NewStepCommentUpVote(db.Gorm)

	// * third party
	var oauthService = services2.NewOAuthService(config.Env)
	var jwtService = services2.NewJwtService()

	// * Services
	var loginService = services.NewLoginService(userRepo, oauthService, jwtService)
	var profileService = services.NewProfileService(userRepo)
	var courseService = services.NewCourseService(courseRepo)
	var progressService = services.NewProgressService(userRepo, courseRepo)
	var stepService = services.NewStepService(stepEvalRepo, userEvalRepo, userRepo, stepCommentRepo, stepCommentUpVoteRepo)

	// * Controller
	var loginController = controllers.NewLoginController(config.Env, loginService)
	var profileController = controllers.NewProfileController(profileService)
	var courseController = controllers.NewCourseController(courseService)
	var progressController = controllers.NewProgressController(progressService)
	var stepController = controllers.NewStepController(stepService)

	serverAddr := fmt.Sprintf("%s:%d", *config.Env.ServerHost, *config.Env.ServerPort)

	// Initialize fiber instance
	app := NewFiberApp()

	// Register root endpoint
	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(response.InfoResponse[string]{
			Success: true,
			Message: "BOOKMARK_API_ROOT",
		})
	})

	// * cores
	app.Use(middleware.Cors)

	// * Recover
	app.Use(Recover())

	api := app.Group("/api")
	api.Static("/static", "./resources/static")

	login := api.Group("/login")
	login.Get("/redirect", loginController.LoginRedirect)
	login.Post("/callback", loginController.LoginCallBack)

	profile := api.Group("/profile", middleware.Jwt())
	profile.Get("/info", profileController.ProfileUserInfo)
	profile.Get("/totalgems", profileController.GetUserGems)

	course := api.Group("/courses", middleware.Jwt())
	course.Get("/field/:field_id", courseController.GetCoursesByFieldId)
	course.Get("/current", courseController.GetCurrentCourse)
	course.Get("/:courseId/total-steps", courseController.GetTotalStepsByCourseId)

	progress := api.Group("/progress", middleware.Jwt())
	progress.Get("/:courseId/percentage", progressController.GetCompletionPercentage)

	step := api.Group("/step", middleware.Jwt())
	step.Get("/gem/:stepId", stepController.GetGemEachStep)
	step.Get("/comment/:stepId", stepController.GetStepComment)

	// Custom handler to set Content-Type header based on file extension
	api.Use("/static", func(c *fiber.Ctx) error {
		filePath := c.Path()
		contentType := getContentType(filePath)
		c.Set("Content-Type", contentType)
		return c.Next()
	})

	// * swagger
	api.Get("/swagger/*", swagger.HandlerDefault)

	// * Not found
	api.Use(handler.NotFoundHandler)

	ListenAndServe(app, serverAddr)
}

func getContentType(filename string) string {
	extension := strings.ToLower(filepath.Ext(filename))
	switch extension {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	// Add more cases for other file types if needed
	default:
		return "application/octet-stream"
	}
}

func NewFiberApp() *fiber.App {
	fiberConfig := fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	}

	app := fiber.New(fiberConfig)

	return app
}

func ListenAndServe(app *fiber.App, serverAddr string) {
	err := app.Listen(serverAddr)
	if err != nil {
		panic(fmt.Errorf("[Server] Unable to start server: %w", err))
	}
	logrus.Printf("[Server] Server started successfully")
}
