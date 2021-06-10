package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/app/view"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
)

func (ctrl *Controller) Update(c *gin.Context) {
	id := c.Param("id")

	input := &userin.UpdateInput{}
	input.ID = id

	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, err)
		// fmt.Println("error")
		return
	}
	fmt.Println("user input update", input)

	a, err := ctrl.service.Update(c, input)
	fmt.Println("a, err:", a, err)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeCreatedResp(c, id)
}
