package handler

import "github.com/gin-gonic/gin"

func (i *impl) GetEvents(c *gin.Context) {
	i.logger.Info("GetEvents")

	res, err := i.eventSvc.GetAllEvents(c)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, res)
}

func (i *impl) GetEventDetails(c *gin.Context) {

}
