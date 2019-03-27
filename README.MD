# easyrouter
Simple string based Golang router. Any help is welcomed

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


#### Methods
Router handle all HTTP methods that are supported by gorilla/mux, syntax for routes are the same
To create resource route you can use 
`r.RESOURCE(path, controller)`
which is shortcut for:
```
	router.addRoute("GET", path, controller, "Index")
	router.addRoute("GET", path+"/{id:[0-9]+}", controller, "Index")
	router.addRoute("POST", path, controller, "Store")
	router.addRoute("PUT", path+"/{id:[0-9]+}", controller, "Update")
	router.addRoute("DELETE", path, controller, "Delete")
```


### Notes
1. Router is build upon gorilla/mux since it's stable and popular router.
2. For now router needs gorm.DB pointer to initialize, I'm working on conditional dependency injecting so anyone will be able to customize router core
