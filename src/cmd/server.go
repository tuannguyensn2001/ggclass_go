package cmd

import (
	"contrib.go.opencensus.io/exporter/jaeger"
	"ggclass_go/src/config"
	"ggclass_go/src/middlewares"
	"ggclass_go/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"log"
	"net/http"
)

func server() *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {

			isProduction := config.Cfg.IsProduction

			if isProduction {
				gin.SetMode(gin.ReleaseMode)
			}

			r := gin.Default()

			r.Use(middlewares.Cors)
			r.Use(middlewares.Recover)

			routes.MatchRoutes(r)

			je, err := jaeger.NewExporter(jaeger.Options{
				AgentEndpoint:     "localhost:6831",
				CollectorEndpoint: "http://localhost:14268/api/traces",
				Process:           jaeger.Process{ServiceName: "ggclass"},
			})

			if err != nil {
				log.Fatalln("err init jaeger")
			}

			trace.RegisterExporter(je)
			trace.ApplyConfig(trace.Config{
				DefaultSampler: trace.ProbabilitySampler(1),
			})

			err = http.ListenAndServe(":"+config.Cfg.GetPort(), &ochttp.Handler{
				Handler: r,
			})
			if err != nil {
				log.Fatalln("err init server", err)
			}
		},
	}
}
