# easyrouter
Simple string based Golang router

This project is created to provide easy router for "Laravel switcher".

### Example

#### Routing
```
r := router.New(db)
r.DebugMode(true)
r.AddRegistry("HomeController", (*controllers.HomeController)(nil))
r.GET("/test", "HomeController", "Index")
```

#### Controller
```
type HomeController struct {
	router.Controller
}

func (h *HomeController) Index(r *http.Request) (string, error) {
	return "Executed", errors.New("Test")
}
```

#### Action return statements
To provide all needed features of response action function can returns three variables

`(interface{}, error, int)`

Where interface is text response that will be printed as json, error is any error that occurs while executing action and int is response http status code. By default 200.