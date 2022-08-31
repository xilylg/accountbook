package admin

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/gin-gonic/gin/binding"
// 	"github.com/xilylg/accountbook/middleware"
// 	"github.com/xilylg/accountbook/model"
// 	"github.com/xilylg/accountbook/service"
// )

// func AddUser(c *gin.Context) {
// 	us := &model.User{Nickname: "aaa", Password: "aaaa"}
// 	service.User.Add(c, us)
// }

// func UserList(c *gin.Context) {
// 	var query UserList
// 	if err := c.ShouldBindWith(&query, binding.Query); err == nil {
// 		middleware.Ok(c, "admin pong")
// 	} else {
// 		middleware.Fail(c, "100", err.Error())
// 	}
// 	service.User.List(c, us)
// }
