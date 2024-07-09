package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/ruziba3vich/redis_practice/handler"
	customRedis "github.com/ruziba3vich/redis_practice/redis"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	myRedis := customRedis.NewMyRedis(redisClient)

	handler := handler.New(myRedis)

	router := gin.Default()

	router.PUT("/set/:key/:value", handler.SetKeyValueHandler)
	router.GET("/get/:key", handler.GetKeyValueHandler)
	router.DELETE("/delete/:key", handler.DeleteKeyHandler)
	router.GET("/exists/:key", handler.ExistsKeyHandler)

	router.PUT("/set/add/:setname/:member", handler.AddToSetHandler)
	router.GET("/set/get/:setname/:member", handler.GetFromSetHandler)
	router.DELETE("/set/remove/:setname/:member", handler.RemoveFromSetHandler)

	router.PUT("/hash/add/:hashname/:key/:value", handler.AddToHashHandler)
	router.DELETE("/hash/remove/:hashname/:key", handler.RemoveFromHashHandler)
	router.GET("/hash/exists/:hashname/:key", handler.ExistsInHashHandler)
	router.GET("/hash/getall/:hashname", handler.GetAllFromHashHandler)

	router.PUT("/list/lpush/:listname/:value", handler.LeftPushHandler)
	router.PUT("/list/rpush/:listname/:value", handler.RightPushHandler)
	router.GET("/list/lpop/:listname", handler.PopLeftHandler)
	router.GET("/list/rpop/:listname", handler.PopRightHandler)
	router.GET("/list/length/:listname", handler.ListLengthHandler)
	router.GET("/list/range/:listname/:from/:to", handler.GetRangeElementsHandler)

	router.Run(":7777")
}
