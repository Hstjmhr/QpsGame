package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

// AutoReplyListView 列表
func (BigModelApi) AutoReplyListView(c *gin.Context) {
	var cr models.PageInfo
	c.ShouldBindQuery(&cr)

	list, count, _ := common.ComList(models.AutoReplyModel{}, common.Option{
		PageInfo: cr,
		Likes:    []string{"name"},
	})
	res.OkWithList(list, count, c)
	return
}
