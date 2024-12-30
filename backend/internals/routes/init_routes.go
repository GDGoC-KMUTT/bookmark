package routes

import (
	_ "backend/docs"
	"backend/internals/config"
	"backend/internals/controllers"
	"backend/internals/db"
	"backend/internals/entities/response"
	"backend/internals/minio"
	"backend/internals/repositories"
	"backend/internals/routes/handler"
	"backend/internals/routes/middleware"
	"backend/internals/services"
	services2 "backend/internals/utils/services"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"
)

func SetupRoutes() {
	// * Repositories
	var userRepo = repositories.NewUserRepository(db.Gorm)
	var courseRepo = repositories.NewCourseRepository(db.Gorm)
	var coursePageRepo = repositories.NewCoursePageRepository(db.Gorm)
	var fieldTypeRepo = repositories.NewFieldTypeRepository(db.Gorm)
	var articleRepo = repositories.NewArticleRepository(db.Gorm)
	var moduleRepo = repositories.NewModuleRepository(db.Gorm)
	var moduleStepRepo = repositories.NewStepRepository(db.Gorm)
	var enrollRepo = repositories.NewEnrollRepository(db.Gorm)
	var stepEvalRepo = repositories.NewStepEvaluateRepository(db.Gorm)
	var userEvalRepo = repositories.NewUserEvaluateRepo(db.Gorm)
	var stepRepo = repositories.NewStepRepository(db.Gorm)
	var stepCommentRepo = repositories.NewStepCommentRepository(db.Gorm)
	var stepCommentUpVoteRepo = repositories.NewStepCommentUpVote(db.Gorm)
	var stepAuthorRepo = repositories.NewStepAuthorRepository(db.Gorm)
	var courseContentRepo = repositories.NewCourseContentRepository(db.Gorm)

	// * third party
	var oauthService = services2.NewOAuthService(config.Env)
	var jwtService = services2.NewJwtService()
	var minioService = services2.NewMinioService(minio.MinioClient)

	// * Services
	var loginService = services.NewLoginService(userRepo, oauthService, jwtService)
	var profileService = services.NewProfileService(userRepo)
	var courseService = services.NewCourseService(courseRepo, fieldTypeRepo)
	var coursePageService = services.NewCoursePageService(&coursePageRepo, courseRepo)
	var progressService = services.NewProgressService(userRepo, courseRepo)
	var stepService = services.NewStepService(
		stepRepo,
		stepEvalRepo,
		stepCommentRepo,
		stepCommentUpVoteRepo,
		stepAuthorRepo,
		userRepo,
		userEvalRepo,
		courseContentRepo,
		moduleRepo)
	var articleService = services.NewArticleService(articleRepo)
	var moduleService = services.NewModuleService(&moduleRepo)
	var moduleStepService = services.NewModuleStepService(&moduleStepRepo)
	var enrollService = services.NewEnrollService(enrollRepo)

	// * Controller
	var loginController = controllers.NewLoginController(config.Env, loginService)
	var profileController = controllers.NewProfileController(profileService)
	var courseController = controllers.NewCourseController(courseService)
	var coursePageController = controllers.NewCoursePageController(coursePageService)
	var articleController = controllers.NewArticleController(articleService)
	var progressController = controllers.NewProgressController(progressService)
	var moduleController = controllers.NewModuleController(moduleService)
	var moduleStepController = controllers.NewModuleStepController(moduleStepService)
	var enrollController = controllers.NewEnrollController(enrollService)
	var stepController = controllers.NewStepController(stepService, config.Env, minioService)

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

	// * Course routes
	course := api.Group("/courses", middleware.Jwt())
	course.Get("/field/:fieldId", courseController.GetCoursesByFieldId)
	course.Get("/field-types", courseController.GetAllFieldTypes)
	course.Get("/current", courseController.GetCurrentCourse)
	course.Get("/:courseId/total-steps", courseController.GetTotalStepsByCourseId)
	course.Get("/enrolled", courseController.GetEnrollCourseByUserId)
	course.Get("/:coursePageId/info", coursePageController.GetCoursePageInfo)
	course.Get("/:coursePageId/content", coursePageController.GetCoursePageContent)
	course.Get("/suggest/:fieldId", coursePageController.GetSuggestCoursesByFieldId)

	// * Module routes
	module := api.Group("/module", middleware.Jwt())
	module.Get("/:moduleId/info", moduleController.GetModuleInfo)

	// * Step routes
	step := api.Group("/step", middleware.Jwt())
	step.Get("/:moduleId/info", moduleStepController.GetModuleSteps)

	// * Article routes
	article := api.Group("/article", middleware.Jwt())
	article.Get("", articleController.GetAllArticles)

	// * Progress routes
	progress := api.Group("/progress", middleware.Jwt())
	progress.Get("/:courseId/percentage", progressController.GetCompletionPercentage)

	// * Enroll routes
	enroll := api.Group("/enroll", middleware.Jwt())
	enroll.Post("/:courseId", enrollController.EnrollInCourse)
	
	step := api.Group("/step", middleware.Jwt())
	step.Get("/:stepId", stepController.GetStepInfo)
	step.Get("/gem/:stepId", stepController.GetGemEachStep)

	stepEval := step.Group("/stepEval")
	stepEval.Post("/submit", stepController.SubmitStepEval)
	stepEval.Get("/status", stepController.CheckStepEvalStatus)
	stepEval.Post("/submit-type-check", stepController.SubmitStepEvalTypCheck)
	stepEval.Get("/:stepId", stepController.GetStepEvaluate)

	stepComment := step.Group("/comment")
	stepComment.Post("/create", stepController.CommentOnStep)
	stepComment.Post("/upvote", stepController.UpVoteStepComment)
	stepComment.Get("/:stepId", stepController.GetStepComment)

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
