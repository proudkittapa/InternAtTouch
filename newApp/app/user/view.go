package user

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/app/view"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) View(c *gin.Context) {
	id := c.Param("id")
	// span, ctx := opentracing.StartSpanFromContextWithTracer(
	// 	c.Request.Context(),
	// 	opentracing.GlobalTracer(),
	// 	"handler.staff.Create",
	// )
	// defer span.Finish()
	input := &userin.ViewInput{}
	// if err := c.ShouldBindJSON(input); err != nil {
	// 	// view.MakeErrResp(c, err)
	// 	fmt.Println("error")
	// 	return
	// }
	input.ID = id
	// _, err := ctrl.service.Create(c, input)
	_, err := ctrl.service.View(c, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeCreatedResp(c, id)
}
