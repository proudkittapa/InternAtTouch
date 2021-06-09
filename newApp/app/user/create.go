package user

import (
	"fmt"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/app/view"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	goxid "github.com/touchtechnologies-product/xid"

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
	initID := goxid.New()
	input.ID = initID.Gen()
	// _, err := ctrl.service.Create(c, input)
	err := ctrl.service.Create(c, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeCreatedResp(c, input.ID)
}
