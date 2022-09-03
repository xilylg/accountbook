package front

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/service"
	"github.com/xilylg/accountbook/trans"
)

type FlowC struct{}

var Flow *FlowC

func (m *FlowC) Add(c *gin.Context) {
	post := new(model.FlowAddReq)
	err := c.ShouldBindWith(post, binding.Form)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}
	post.Uid = c.GetInt64("uid")
	code := service.Flow.Add(c, post)
	Resp(c, code, nil)
}

func (m *FlowC) Update(c *gin.Context) {
	post := new(model.FlowUpdateReq)
	err := c.ShouldBindWith(post, binding.Form)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}

	flowId, _ := strconv.ParseInt(c.Param("flow_id"), 10, 64)
	flow, code := service.CheckOwner(c, flowId, "flow")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}
	//超过三天不能修改原来目标
	if time.Now().Unix()-flow.(*model.Flow).CreateTime > trans.TDAY {
		Resp(c, trans.ERROR_OP_EXPIRED, trans.OP_EXPIRED_3D)
		return
	}
	post.FlowId = flowId
	code = service.Flow.Update(c, post)
	Resp(c, code, nil)
}

func (t *FlowC) Review(c *gin.Context) {
	post := new(model.FlowReviewReq)
	err := c.ShouldBindWith(post, binding.Form)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}

	flowId, _ := strconv.ParseInt(c.Param("flow_id"), 10, 64)
	flow, code := service.CheckOwner(c, flowId, "flow")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}

	if flow.(*model.Flow).Status >= trans.STATUS_REVIEW {
		Resp(c, trans.ERROR_OP_EXPIRED, trans.OP_ED)
		return
	}
	post.FlowId = flowId
	code = service.Flow.UpdateReview(c, post)
	Resp(c, code, nil)
}

func (m *FlowC) One(c *gin.Context) {
	flowId, _ := strconv.ParseInt(c.Param("flow_id"), 10, 64)
	Flow, code := service.CheckOwner(c, flowId, "flow")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}
	Resp(c, code, Flow)
}

func (m *FlowC) List(c *gin.Context) {
	query := new(model.FlowListReq)
	err := c.ShouldBindWith(query, binding.Query)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}

	rst, count, code := service.Flow.List(c, query)
	Resp(c, code, map[string]interface{}{"list": rst, "total": count})
}

func (m *FlowC) Delete(c *gin.Context) {
	flowId, _ := strconv.ParseInt(c.Param("Flow_id"), 10, 64)
	_, code := service.CheckOwner(c, flowId, "flow")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}

	code = service.Flow.Delete(c, flowId)
	Resp(c, code, "")
}
