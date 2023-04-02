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
		"shishi": "ff99c98713102b462581f31caf810bd59342a13c308cd0e4fb48695e3fc7fc75",
		"admin":  "058d9cad80e91eb75f0691b63bce23d630cfefe8d5627cb7ee5b10c7ef6ada8d",
	}
	user := c.Query("name")
	fmt.Println("nip05 verify request", user, name2pubkey[user])
	if v, ok := name2pubkey[user]; ok {
		resp := NIP05Resp{Names: map[string]string{}}
		resp.Names[user] = v
		c.JSON(http.StatusOK, resp)
		return
	}
	c.Status(http.StatusNotFound)
	return
}

type NIP05Resp struct {
	Names map[string]string `json:"names"`
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
