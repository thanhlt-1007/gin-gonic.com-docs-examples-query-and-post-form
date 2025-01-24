package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func postPagingHandler(context *gin.Context) {
    page := context.Query("page")
    per := context.DefaultQuery("per", "10")
    message := context.PostForm("message")

    context.JSON(
        http.StatusOK,
        gin.H {
            "page": page,
            "per": per,
            "message": message,
        },
    )
}

func main() {
    engine := gin.Default()
    engine.POST("/paging", postPagingHandler)
    engine.Run()
}
