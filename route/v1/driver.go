package v1

import (
	"github.com/Vioneta/VionetaOS-Common/model"
	"github.com/Vioneta/VionetaOS-Common/utils/common_err"
	"github.com/Vioneta/VionetaOS-LocalStorage/internal/op"
	"github.com/gin-gonic/gin"
)

func ListDriverInfo(c *gin.Context) {
	c.JSON(common_err.SUCCESS, model.Result{Success: common_err.SUCCESS, Message: common_err.GetMsg(common_err.SUCCESS), Data: op.GetDriverInfoMap()})
}
