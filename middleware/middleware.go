package middlewares

import (
	"net/http"
	"fmt"
	"github.com/astaxie/beego"
)


type MainMiddleware struct{
	beego.Controller
}

// Authorization middleware
func Authorization() *beego.Router {
	return func(c *MainMiddleware) {
		// token := strings.Replace(authHeader, "Bearer ", "", 1)
		authHeader := c.Header("Authorization","token")
		// fmt.Println(token,"=====>TOKEN")
		// jwt := jwt.JWT{}

		// if userID, err := jwt.ValidateToken(token); err == nil {
		// 	c.Set("UserID", userID)
		// 	c.Next()
		// } else {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"status": "failed",
		// 		"error":  "Unauthorized",
		// 	})
		// 	c.Abort()
		// 	return
		// }
		fmt.Println("ini middleware jalan")
	}
}
