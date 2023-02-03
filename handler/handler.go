package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func ErrRouter(c *gin.Context) {
	c.String(http.StatusBadRequest, "url err")
}

func NIP05(c *gin.Context) {
	name2pubkey := map[string]string{
		"kirito": "2f7caa968b0ec9bacd55a07cfaf6206aab5a62387c76303c311db949dec8bc57",
	}
	user := c.Query("name")
	fmt.Println("nip05 verify request", user)
	if v, ok := name2pubkey[user]; ok {
		resp := NIP05Resp{}
		resp.Names[user] = v
		c.JSON(http.StatusOK, resp)
	}
	c.Status(http.StatusNotFound)
	return
}

type NIP05Resp struct {
	Names map[string]string `json:"names"`
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origi", "*")
	c.Next()
}
