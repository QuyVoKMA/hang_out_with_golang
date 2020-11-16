func CORSMiddleware(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Access-Token")
		if c.Request.Method == "OPTIONS" || c.Request.Method == "HEAD" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
