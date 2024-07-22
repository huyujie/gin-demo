// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"gin-demo/pkg/config"
	"gin-demo/pkg/logger"
	"gin-demo/pkg/service/demo_api"
	"os"
	"os/signal"
	"syscall"

	// mysql "gin-demo/pkg/util/db"
	"github.com/gin-gonic/gin"
)

func main() {

	//init config
	cfg := flag.String("c", "/etc/gin-demo/config.yaml", "configuration file")
	flag.Parse()
	config.ParseConfig(*cfg)
	c := config.ReadConf()

	//init logger
	logger.Setup("gin-demo")

	//mysql
	// mysql.OpenDatabase(c)

	// init router
	gin.SetMode(c.Common.Gin.Mode)
	routes := gin.New()
	go demo_api.StartGin(c, routes)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		os.Exit(0)
	}()
	select {}

}
