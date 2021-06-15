package main

import (
	"context"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/app"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/config"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/repository/kafka"
	//userRepo "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/repository/user"
	msgBrokerService "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/implement"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	userService "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/implement"
	validatorService "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/validator"
	"log"
)

func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()
	//uRepo, err := userRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBHeroTableName)
	panicIfErr(err)
	kRepo, err := kafka.New(configKafka(appConfig))
	panicIfErr(err)
	//validator := validatorService.New(uRepo)

	//user := userService.New(validator, uRepo, kRepo)
	msgService := msgBrokerService.New(kRepo, user)
	//wg.Add(1)
	msgService.Receiver(topics)
	//time.Sleep(10 * time.Second)

	return app.New(user)
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
	msgbrokerin.TopicCreate,
	msgbrokerin.TopicUpdate,
	msgbrokerin.TopicDelete,
}

