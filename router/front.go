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
			member.POST("", memberC.Add)
			member.PUT("/:mid", memberC.Update)
			member.GET("/list/:page", memberC.List)
			member.GET("/:mid", memberC.One)
			member.DELETE("/:mid", memberC.Delete)
		}

		bookC := new(front.BookC)
		book := r.Group("/book")
		{
			book.POST("", bookC.Add)
			book.PUT("/:book_id", bookC.Update)
			book.GET("/list/:page", bookC.List)
			book.GET("/:book_id", bookC.One)
			book.DELETE("/:book_id", bookC.Delete)
		}

		itemC := new(front.ItemC)
		item := r.Group("/item")
		{
			item.POST("", itemC.Add)
			item.PUT("/:item_id", itemC.Update)
			item.GET("/list/:page", itemC.List)
			item.GET("/:item_id", itemC.One)
			item.DELETE("/:item_id", itemC.Delete)
		}

		targetC := new(front.TargetC)
		target := r.Group("/target")
		{
			target.POST("", targetC.Add)
			target.PUT("/:target_id", targetC.Update)
			target.PUT("/review/:target_id", targetC.Review)
			target.GET("/list/:page", targetC.List)
			target.GET("/:target_id", targetC.One)
			target.DELETE("/:target_id", targetC.Delete)
		}

		flowC := new(front.FlowC)
		flow := r.Group("/flow")
		{
			flow.POST("", flowC.Add)
			flow.GET("/list/:page", flowC.List)
			flow.GET("/:flow_id", flowC.One)
			flow.PUT("/review/:flow_id", flowC.Review)
			flow.PUT("/:flow_id", flowC.Update)
			flow.DELETE("/:flow_id", flowC.Delete)
		}

		userC := new(front.UserC)
		user := r.Group("/user")
		{
			user.POST("", userC.Login)
			user.PUT("/:uid", userC.Update)
			user.GET("/:uid", userC.One)
		}
	}
}
