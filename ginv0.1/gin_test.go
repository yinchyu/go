package gin

var DB = make(map[string]string)

func main() {

	r := Default()

	// Ping test
	r.GET("/ping", func(c *Context) {
		c.String(200, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *Context) {
		//user := c.Params.ByName("name")
		user := "ycy"
		value, ok := DB[user]
		if ok {
			c.JSON(200, H{"user": user, "value": value})
		} else {
			c.JSON(200, H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(BasicAuth(Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", BasicAuth(Accounts{
		{"foo", "bar"},  //1. user:foo password:bar
		{"manu", "123"}, //2. user:manu password:123
	}))

	authorized.POST("admin", func(c *Context) {
		user := c.Get("user").(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}
		if c.EnsureBody(&json) {
			DB[user] = json.Value
			c.JSON(200, H{"status": "ok"})
		}
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8081")
}
