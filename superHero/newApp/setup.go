package main

import (
	"context"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/repository/kafka"
	msgBrokerService "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/implement"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	"log"
	"time"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/config"

	// validatorService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/validator"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/app"
	validatorService "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/validator"

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
	panicIfErr(err)
	kRepo, err := kafka.New(configKafka(appConfig))
	panicIfErr(err)
	validator := validatorService.New(uRepo)

	user := userService.New(validator, uRepo, kRepo)
	msgService := msgBrokerService.New(kRepo, user)
	//wg.Add(1)
	go msgService.Receiver(topics)
	time.Sleep(10 * time.Second)
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

func configKafka(appConfig *config.Config) *kafka.Config {
	return &kafka.Config{
		BackOffTime:  appConfig.MessageBrokerBackOffTime,
		MaximumRetry: appConfig.MessageBrokerMaximumRetry,
		Host:         appConfig.MessageBrokerEndpoint,
		Group:        appConfig.MessageBrokerGroup,
		Version:      appConfig.MessageBrokerVersion,
	}
}
var topics = []msgbrokerin.TopicMsgBroker{
	msgbrokerin.TopicResponse,
	msgbrokerin.TopicCreate,
	//msgbrokerin.TopicOTP,
	//msgbrokerin.TopicVerify,
}

//hi
