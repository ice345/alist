package middlewares

import (
	"github.com/NewAlist/alist/internal/conf"
	"github.com/NewAlist/alist/internal/errs"
	"github.com/NewAlist/alist/internal/setting"
	"github.com/NewAlist/alist/server/common"
	"github.com/gin-gonic/gin"
)

func SearchIndex(c *gin.Context) {
	mode := setting.GetStr(conf.SearchIndex)
	if mode == "none" {
		common.ErrorResp(c, errs.SearchNotAvailable, 500)
		c.Abort()
	} else {
		c.Next()
	}
}
