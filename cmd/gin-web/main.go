package main

import (
	"github.com/shiyongxi/go-common/common"
	"github.com/shiyongxi/go-common/server"
	"github.com/shiyongxi/go-gin-web/utils"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"os"
)

type (
	Server struct {
		Config *utils.Config
		//Logger     logs.Factory
		BaseServer *server.GinServer
	}
)

func NewServer(cfg *utils.Config /*, log logs.Factory*/) *Server {
	return &Server{
		Config: cfg,
		/*Logger:     log,*/
		BaseServer: server.NewGinServer(cfg.System),
	}
}

func main() {
	cmd := NewCommands().Commands()

	cmd.CliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "run Commands",
			Action: func(c *cli.Context) {
				var cfg utils.Config

				err := common.NewConfig(&cfg).ReadYaml(cmd.PathFile)
				if err != nil {
					zap.S().Error(zap.Error(err))
				}
				utils.Set(&cfg)

				// log := logs.NewFactory(logs.NewStore(cfg.Logger, zap.InfoLevel).JsonEncoder())
				svc := NewServer(&cfg)

				//mysqlConn := mysql.NewMysql(cfg.Mysql).Connection()
				//				//defer func() {
				//				//	log.Bg().Info("[mysql conn close]", zap.Error(mysqlConn.Close()))
				//				//}()

				svc.BaseServer.I18nBundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
				svc.BaseServer.I18nBundle.MustLoadMessageFile("./locales/en-US.yaml")
				svc.BaseServer.I18nBundle.MustLoadMessageFile("./locales/zh-CN.yaml")

				//svc.BaseServer.LoopCall(NewManufacture(&cfg, log))

				svc.BaseServer.RegisterRoute = svc.Router

				svc.BaseServer.Run()
			},
		},
	}

	if err := cmd.CliApp.Run(os.Args); err != nil {
		zap.L().Error("[commands]", zap.Error(err))
	}
}
