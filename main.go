package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("", func(ctx *gin.Context) {
		//  ctx.String(http.StatusOK, "hello Abu bakar")
		ctx.JSON(http.StatusOK, gin.H{ //! defult json pass
			"status":    "success",
			"name":      "Abu bakar",
			"roll":      "936943",
			"dep":       "futter",
			"available": true,
		})

	})

	// r.GET("/name",func(ctx *gin.Context) {   //! params a data pass korar jonno
	// 	var val =ctx.Param("name")
	// 	ctx.JSON(http.StatusOK,gin.H{
	// 		"value":val,

	// 	})
	// })

	// post method

	r.POST("/add", func(ctx *gin.Context) {
		var data TestModel
		myslice := map[string]string{
			"name":   "Abu Bakar",
			"dep":    "Bangladesh",
			"user":   "Abu",
			"pass":   "4565122",
			"status": "true",
		}
		if err := ctx.ShouldBind(&data); err != nil {
			fmt.Println(err)

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})

		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data": data,
				"sub":  myslice,
			})

		}

	})

	r.Run("127.0.0.1:8080")

}

//post Api model create

type TestModel struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
