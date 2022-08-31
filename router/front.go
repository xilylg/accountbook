package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/controller/front"
	"github.com/xilylg/accountbook/middleware"
)

func frontRouter(r *gin.Engine) {
	// r.GET("/login", cadmin.Login)
	r.Use(middleware.UserAuth())
	{
		memberC := new(front.MemberC)
		member := r.Group("/member")
		{
			member.POST("/", memberC.Add)
			// member.PUT("/:mid", memberC.Update)
			// member.GET("/list", memberC.List)
			// member.GET("/:mid", memberC.One)
			// member.DELETE("/:mid", memberC.Delete)
		}
		// accountbookC := new
		// accountbook := front.Group("/accountbook")
		// {
		// 	member.POST("/", front.Accountbook.Add)
		// 	member.PUT("/:aid", front.Accountbook.Update)
		// 	member.GET("/list", front.Accountbook.List)
		// 	member.GET("/:aid", front.Accountbook.One)
		// 	member.DELETE("/:aid", front.Accountbook.Delete)
		// }
	}
}
