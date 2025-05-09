package userapp

import (
	mysqluser "github.com/mreza-mz/userapp.git/repository/mysql"
	"shop/adapter/mysql"
	"shop/adapter/redis"
	"shop/core"
	cfg "shop/core/config"
	"shop/core/userapp/delivery/httpserver/otphandler"
	"shop/core/userapp/delivery/httpserver/userhandler"
	"shop/core/userapp/repository/mysql"
	"shop/core/userapp/repository/redis/redisotp"
	"shop/core/userapp/service/authservice"
	"shop/core/userapp/service/mockotpservice"
	"shop/core/userapp/service/otpservice"
	"shop/core/userapp/service/userservice"
	"shop/core/userapp/validator/otpvalidator"
	"shop/pkg/caching"
	"shop/pkg/notifier"
)

func NewBuildUserApp(cfg cfg.Config, server core.Server) {
	// adapters
	db := mysql.New(cfg.Mysql)
	redisAdapter := redis.New(cfg.Redis)
	caching.New(*redisAdapter)

	// repos
	repo := mysqluser.New(db)
	otpRedis := redisotp.New(redisAdapter)

	// services
	authSvc := authservice.New(cfg.Auth)
	userSvc := userservice.New(authSvc, repo)
	var otpSvc otphandler.OtpService
	if cfg.OTP.UseMock {
		otpSvc = mockotpservice.New()
	} else {
		svc := otpservice.New(cfg.OTP, notifier.New, otpRedis, userSvc)
		otpSvc = &svc
	}

	userSvc.WithOtpService(otpSvc.Verify)

	// validators
	otpValidator := otpvalidator.New()

	// handlers
	userHandler := userhandler.New(userSvc)
	otpHandler := otphandler.New(otpSvc, otpValidator)

	// set routes
	userHandler.SetRoutes(server.Router)
	otpHandler.SetRoutes(server.Router)
}
