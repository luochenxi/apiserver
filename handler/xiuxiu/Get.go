package xiuxiu

import (
	"fmt"

	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

//Get a similar result by xiuxiu images retrieval
func Get(c *gin.Context) {
	log.Info("Xiuxiu Get function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r GetRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	// a GET request to /xiuxiu/john
	admin2 := c.Param("image") // image=john
	log.Infof("URL image: %s", admin2)
	// GET /path?desc=1234
	// DefaultQuery()：类似 Query()，但是如果 key 不存在，会返回默认值
	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)
	// GetHeader()：获取 HTTP 头。
	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("image is: [%s], password is [%s]", r.Image, r.Top)
	if r.Image == "" {
		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("image can not found in xiuxiu:xx.xx.xx.xx")), nil)
		return
	}

	if r.Top == "" {
		SendResponse(c, fmt.Errorf("password is empty"), nil)
	}

	req := GetResponse{
		Image: r.Image,
	}

	//Show the Image infomation.
	SendResponse(c, nil, req)
}
