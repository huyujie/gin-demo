// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package demo1

import (
	"gin-demo/pkg/logger"
	"gin-demo/pkg/service/demo_api/v1/demo1/handler"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	apiv1 := r.Group("/api/:zone/v1")
	apiv1.Use(ginzap.Ginzap(logger.GetLogger(), time.RFC3339, false))
	{
		// demo
		apiv1.POST("/GetProjects", handler.GetProjects)

		// hello world
		apiv1.GET("/GetHelloWorld", handler.GetHelloWorld)

	}
}
