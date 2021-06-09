package main

import (
	"io"
	"log"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/config"

	// validatorService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/validator"

	validatorService "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/validator"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/app"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"

	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/app"

	jaegerConf "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/config"
	// compRepo "github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/company"
	// staffRepo "github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/staff"
	// companyService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/implement"
	// staffService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/implement"
)

func setupJaeger(appConfig *config.Config) io.Closer {
	cfg, err := jaegerConf.FromEnv()
	panicIfErr(err)

	cfg.ServiceName = appConfig.AppName
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter = &jaegerConf.ReporterConfig{LogSpans: true}

	jLogger := jaegerLog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegerConf.Logger(jLogger),
		jaegerConf.Metrics(jMetricsFactory),
	)
	panicIfErr(err)
	opentracing.SetGlobalTracer(tracer)

	return closer
}

func newApp(appConfig *config.Config) *app.App {
	// ctx := context.Background()

	// cRepo, err := compRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBCompanyTableName)
	// panicIfErr(err)
	// sRepo, err := staffRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBStaffTableName)
	// panicIfErr(err)

	validator := validatorService.New()
	// generateID, err := util.NewUUID()
	// panicIfErr(err)

	// company := companyService.New(validator, cRepo, generateID)
	// staff := staffService.New(validator, sRepo, generateID)
	user := userService.New(validator)

	return app.New(staff, company)
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
