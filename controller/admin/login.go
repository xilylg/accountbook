package admin

// type UserInfo struct {
// 	Username string `form:"username" binding:"required,username"`
// 	Password string `form:"password" binding:"required,password"`
// }

// func (u *UserInfo) GetError(err validator.ValidationErrors) string {
// 	for _, val := range err {
// 		switch val.StructField() {
// 		case "Username":
// 			switch val.Tag() {
// 			case "required":
// 				return "请输入用户名"
// 			case "username":
// 				return "请正确书写用户名"
// 			}
// 		case "Password":
// 			switch val.Tag() {
// 			case "password":
// 				fmt.Printf("dadsafds\n")
// 				return "密码错误"
// 			}
// 		}
// 	}
// 	return ""
// }

// func Login(c *gin.Context) {
// 	var query UserInfo
// 	if err := c.ShouldBindWith(&query, binding.Query); err == nil {
// 		middleware.Ok(c, "admin pong")
// 	} else {
// 		middleware.Fail(c, "100", err.Error())
// 	}

// }
