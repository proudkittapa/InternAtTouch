package main

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/app"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/config"
	elasRepo "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/repository/elastic"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/repository/kafka"
	msgBrokerService "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/implement"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/msgbroker/msgbrokerin"
	userService "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/implement"
	validatorService "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/validator"
	"log"
	"time"
)

//func newApp(appConfig *config.Config) *app.App {
//	ctx := context.Background()
//	//uRepo, err := userRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBHeroTableName)
//	//panicIfErr(err)
//	kRepo, err := kafka.New(configKafka(appConfig))
//	panicIfErr(err)
//	//validator := validatorService.New(uRepo)
//
//	//user := userService.New(validator, uRepo, kRepo)
//	//msgService := msgBrokerService.New(kRepo, user)
//	//wg.Add(1)
//	//msgService.Receiver(topics)
//	//time.Sleep(10 * time.Second)
//	//return app.New(user)
//}

func newApp(appConfig *config.Config) (*app.App) {
	//ctx := context.Background()
	//uRepo, err := userRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBHeroTableName)
	//panicIfErr(err)
	elasRepo, err := elasRepo.New(appConfig.ElasticDBEndpoint, appConfig.ElasticDBUsername, appConfig.ElasticDBPassword)
	kRepo, err := kafka.New(configKafka(appConfig))
	panicIfErr(err)
	validator := validatorService.New(elasRepo)

	user := userService.New(validator, elasRepo, kRepo)
	msgService := msgBrokerService.New(kRepo, user)
	//wg.Add(1)
	msgService.Receiver(topics)
	time.Sleep(10 * time.Second)
	return app.New(user)
	//return a
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

