package route

import (
	"github.com/gin-gonic/gin"
	activationAuth "github.com/j2eevip/gin-restful-example/controllers/auth-controllers/activation"
	forgotAuth "github.com/j2eevip/gin-restful-example/controllers/auth-controllers/forgot"
	loginAuth "github.com/j2eevip/gin-restful-example/controllers/auth-controllers/login"
	registerAuth "github.com/j2eevip/gin-restful-example/controllers/auth-controllers/register"
	resendAuth "github.com/j2eevip/gin-restful-example/controllers/auth-controllers/resend"
	resetAuth "github.com/j2eevip/gin-restful-example/controllers/auth-controllers/reset"
	handlerActivation "github.com/j2eevip/gin-restful-example/handlers/auth-handlers/activation"
	handlerForgot "github.com/j2eevip/gin-restful-example/handlers/auth-handlers/forgot"
	handlerLogin "github.com/j2eevip/gin-restful-example/handlers/auth-handlers/login"
	handlerRegister "github.com/j2eevip/gin-restful-example/handlers/auth-handlers/register"
	handlerResend "github.com/j2eevip/gin-restful-example/handlers/auth-handlers/resend"
	handlerReset "github.com/j2eevip/gin-restful-example/handlers/auth-handlers/reset"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Auth
	*/
	LoginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(LoginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)

	registerRepository := registerAuth.NewRepositoryRegister(db)
	registerService := registerAuth.NewServiceRegister(registerRepository)
	registerHandler := handlerRegister.NewHandlerRegister(registerService)

	activationRepository := activationAuth.NewRepositoryActivation(db)
	activationService := activationAuth.NewServiceActivation(activationRepository)
	activationHandler := handlerActivation.NewHandlerActivation(activationService)

	resendRepository := resendAuth.NewRepositoryResend(db)
	resendService := resendAuth.NewServiceResend(resendRepository)
	resendHandler := handlerResend.NewHandlerResend(resendService)

	forgotRepository := forgotAuth.NewRepositoryForgot(db)
	forgotService := forgotAuth.NewServiceForgot(forgotRepository)
	forgotHandler := handlerForgot.NewHandlerForgot(forgotService)

	resetRepository := resetAuth.NewRepositoryReset(db)
	resetService := resetAuth.NewServiceReset(resetRepository)
	resetHandler := handlerReset.NewHandlerReset(resetService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)
	groupRoute.POST("/activation/:token", activationHandler.ActivationHandler)
	groupRoute.POST("/resend-token", resendHandler.ResendHandler)
	groupRoute.POST("/forgot-password", forgotHandler.ForgotHandler)
	groupRoute.POST("/change-password/:token", resetHandler.ResetHandler)

}
