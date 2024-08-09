package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sudheermurari-07/shorten-url/database"
)

func ResolveURL(ctx *gin.Context) {
	url := ctx.Param("url")
	r := database.CreateClient(0)

	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()

	if err == redis.Nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "short not found in db",
		})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot connect to DB",
		})
		return
	}

	rIncr := database.CreateClient(1)
	defer rIncr.Close()

	_ = rIncr.Incr(database.Ctx, "counter")

	ctx.Redirect(301, value)

}
