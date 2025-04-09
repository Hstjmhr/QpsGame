package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

// ModelUsableListView 可用的大模型列表接口
func (BigModelApi) ModelUsableListView(c *gin.Context) {
	res.OkWithData(global.Config.BigModel.ModelList, c)
	return
}
