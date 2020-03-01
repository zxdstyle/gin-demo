package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func All(c *gin.Context)  {
	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("%c",data)
}
