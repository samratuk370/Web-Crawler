
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
)

func main() {
    r := gin.Default()

    r.Use(authMiddleware())

    r.POST("/api/urls", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "URL submitted"})
    })

    r.GET("/api/urls", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"urls": []string{}})
    })

    r.POST("/api/urls/:id/start", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Crawl started"})
    })

    r.GET("/api/urls/:id", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"url": gin.H{"id": c.Param("id"), "title": "Demo"}})
    })

    r.Run(":" + os.Getenv("PORT"))
}

func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.GetHeader("Authorization") != "Bearer "+os.Getenv("AUTH_TOKEN") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            return
        }
        c.Next()
    }
}
