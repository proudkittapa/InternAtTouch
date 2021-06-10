package main

import (
	"context"
	"log"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/config"

	// validatorService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/validator"

	validatorService "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/validator"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/app"

	userRepo "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/repository/user"
	userService "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/implement"
	"github.com/sirupsen/logrus"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/app"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/config"
	// compRepo "github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/company"
	// staffRepo "github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/staff"
	// companyService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/implement"
	// staffService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/implement"
)

func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()
	uRepo, err := userRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBHeroTableName)
	// cRepo, err := compRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBCompanyTableName)
	panicIfErr(err)
	// sRepo, err := staffRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBStaffTableName)
	// panicIfErr(err)

	validator := validatorService.New()
	// generateID, err := util.NewUUID()
	// panicIfErr(err)

	// company := companyService.New(validator, cRepo, generateID)
	// staff := staffService.New(validator, sRepo, generateID)
	user := userService.New(validator, uRepo)

	return app.New(user)
}

func setupLog() *logrus.Logger {
	lr := logrus.New()
	lr.SetFormatter(&logrus.JSONFormatter{})

	return lr
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
