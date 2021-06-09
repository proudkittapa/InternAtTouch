package user

import (
	"fmt"
	"touch/app/view"
	"touch/service/user/userin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Create(c *gin.Context) {
	// span, ctx := opentracing.StartSpanFromContextWithTracer(
	// 	c.Request.Context(),
	// 	opentracing.GlobalTracer(),
	// 	"handler.staff.Create",
	// )
	// defer span.Finish()

	input := &userin.CreateInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		// view.MakeErrResp(c, err)
		fmt.Println("error")
		return
	}

	_, err := ctrl.service.Create(input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	// view.MakeCreatedResp(c, ID)
}
