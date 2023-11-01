package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"haloKv/src/config"
	_ "haloKv/src/config"
	"haloKv/src/pb"
	"net/http"
)
import "haloKv/src/raft"
import s "haloKv/src/server"

func main() {
	c := config.ReadConfig()

	raftServer := raft.GetRaftInstance(&c)
	serv := s.GetServer()

	raftServer.StartRaftServer(&serv.LogChan)

	r := gin.Default()

	r.GET("/s/get", func(context *gin.Context) {
		key := context.Query("key")
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"key":     key,
			"value":   serv.Get(key),
		})
	})

	r.GET("/s/set", func(context *gin.Context) {
		key := context.Query("key")
		value := context.Query("value")
		serv.Put(key, value)
		log := pb.Log{}
		*log.Key = key
		*log.Value = value

		context.JSON(http.StatusOK, gin.H{
			"message": raftServer.PutLog(&log),
		})
	})
	err := r.Run(string(c.ServerConfig.Port))
	if err != nil {
		fmt.Println("Start Panic : {}", err)
	}
}
