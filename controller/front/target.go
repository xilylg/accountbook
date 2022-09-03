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

type TargetC struct{}

var Target *TargetC

func (m *TargetC) Add(c *gin.Context) {
	post := new(model.TargetAddReq)
	err := c.ShouldBindWith(post, binding.Form)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}
	post.Uid = c.GetInt64("uid")
	post.Status = trans.STATUS_NORMAL
	code := service.Target.Add(c, post)
	Resp(c, code, nil)
}

func (m *TargetC) Update(c *gin.Context) {
	post := new(model.TargetUpdateReq)
	err := c.ShouldBindWith(post, binding.Form)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}

	TargetId, _ := strconv.ParseInt(c.Param("target_id"), 10, 64)
	target, code := service.CheckOwner(c, TargetId, "target")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}
	//超过三天不能修改原来目标
	if time.Now().Unix()-target.(*model.Target).CreateTime > trans.TDAY {
		Resp(c, trans.ERROR_OP_EXPIRED, trans.OP_EXPIRED_3D)
		return
	}
	post.TargetId = TargetId
	code = service.Target.Update(c, post)
	Resp(c, code, nil)
}

func (t *TargetC) Review(c *gin.Context) {
	post := new(model.TargetReviewReq)
	err := c.ShouldBindWith(post, binding.Form)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}

	TargetId, _ := strconv.ParseInt(c.Param("target_id"), 10, 64)
	target, code := service.CheckOwner(c, TargetId, "target")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}
	//超过三天不能修改原来目标
	if target.(*model.Target).Status >= trans.STATUS_REVIEW {
		Resp(c, trans.ERROR_OP_EXPIRED, trans.OP_ED)
		return
	}
	post.TargetId = TargetId
	code = service.Target.UpdateReview(c, post)
	Resp(c, code, nil)
}

func (m *TargetC) One(c *gin.Context) {
	TargetId, _ := strconv.ParseInt(c.Param("target_id"), 10, 64)
	Target, code := service.CheckOwner(c, TargetId, "target")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}
	Resp(c, code, Target)
}

func (m *TargetC) List(c *gin.Context) {
	query := new(model.TargetListReq)
	err := c.ShouldBindWith(query, binding.Query)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}

	rst, count, code := service.Target.List(c, query)
	Resp(c, code, map[string]interface{}{"list": rst, "total": count})
}

func (m *TargetC) Delete(c *gin.Context) {
	TargetId, _ := strconv.ParseInt(c.Param("target_id"), 10, 64)
	_, code := service.CheckOwner(c, TargetId, "target")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}

	code = service.Target.Delete(c, TargetId)
	Resp(c, code, "")
}
