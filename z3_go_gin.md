* start

```
	r := gin.Default() //此处就是new+了log
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{ //H 是 map的缩写
			"message": "pong",
		})
	})
	r.Run() //默认监听 8080
```



* form data

```
	r := gin.Default()
	r.POST("form", func(c *gin.Context) {
		val1 := c.DefaultPostForm("key1", "val1")
		val2 := c.PostForm("key2")
		c.JSON(200, gin.H{
			"key1": val1,
			"key2": val2,
		})
	})
	r.Run()
```



