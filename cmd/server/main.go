package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/hatech/backup/cmd/server/config"
	"github.com/hatech/backup/docs"
	"github.com/hatech/backup/pkg/handler"
	"github.com/hatech/backup/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"k8s.io/klog/v2"
	"os"
)

var cmd = &cobra.Command{
	Use:   "hatech kubernetes backup service",
	Short: "hatech kubernetes backup service",
	RunE:  StartBackupService,
}

var cfg config.Config
var cfgFile string

func main() {
	klog.Info(version.PrintVersion())
	InitConfig()
	if err := cmd.Execute(); err != nil {
		klog.Fatalf("hatech backup service cmd execute failed, err: %s", err.Error())
	}
}

func InitConfig() {
	klog.InitFlags(nil)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.yaml", "config file name")

	cobra.OnInitialize(func() {
		v := viper.New()
		v.SetConfigType("yaml")
		v.SetConfigFile(cfgFile)
		v.AutomaticEnv()
		if err := v.ReadInConfig(); err != nil {
			klog.Errorf("load viper config err :%s", err.Error())
			os.Exit(1)
		}

		if err := v.Unmarshal(&cfg); err != nil {
			klog.Errorf("viper config json unmarshal err :%s", err.Error())
			os.Exit(1)
		}
		klog.Infof("server config : %v", cfg)
	})
}

func StartBackupService(cmd *cobra.Command, args []string) error {
	r := gin.Default()
	r.NoRoute(handler.NoRouteHandler)
	StartSwagger(r)
	handler.Register(r.Group("/"))
	return r.Run(cfg.HatechServer.Listen)
}

func StartSwagger(g *gin.Engine) {
	docs.SwaggerInfo.Title = "Hatech Kubernetes Backup Service API"
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
