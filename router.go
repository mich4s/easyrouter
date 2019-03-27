package router

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Router struct {
	registry  map[string]reflect.Type
	Router    *mux.Router
	db        *gorm.DB
	debugMode bool
}

func New(db *gorm.DB) *Router {
	return &Router{
		registry:  make(map[string]reflect.Type),
		Router:    mux.NewRouter(),
		debugMode: false,
	}
}

func (r *Router) DebugMode(flag bool) {
	r.debugMode = flag
}

func (r *Router) AddRegistry(name string, typedNil interface{}) {
	r.registry[name] = reflect.TypeOf(typedNil).Elem()
}

func (router *Router) GET(path string, controller string, action string) {
	router.addRoute("GET", path, controller, action)
}

func (router *Router) HEAD(path string, controller string, action string) {
	router.addRoute("HEAD", path, controller, action)
}

func (router *Router) POST(path string, controller string, action string) {
	router.addRoute("POST", path, controller, action)
}

func (router *Router) PUT(path string, controller string, action string) {
	router.addRoute("PUT", path, controller, action)
}

func (router *Router) DELETE(path string, controller string, action string) {
	router.addRoute("DELETE", path, controller, action)
}

func (router *Router) CONNECT(path string, controller string, action string) {
	router.addRoute("CONNECT", path, controller, action)
}

func (router *Router) OPTIONS(path string, controller string, action string) {
	router.addRoute("OPTIONS", path, controller, action)
}

func (router *Router) TRACE(path string, controller string, action string) {
	router.addRoute("TRACE", path, controller, action)
}

func (router *Router) PATCH(path string, controller string, action string) {
	router.addRoute("PATCH", path, controller, action)
}

func (router *Router) RESOURCE(path string, controller string) {
	router.addRoute("GET", path, controller, "Index")
	router.addRoute("GET", path+"/{id:[0-9]+}", controller, "Index")
	router.addRoute("POST", path, controller, "Store")
	router.addRoute("PUT", path+"/{id:[0-9]+}", controller, "Update")
	router.addRoute("DELETE", path, controller, "Delete")
}

func (router *Router) addRoute(method string, path string, controller string, action string) {
	router.Router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		controller := reflect.ValueOf(reflect.New(router.registry[controller]).Interface())
		dbPointer := controller.Elem().FieldByName("DB")
		if dbPointer.IsValid() {
			dbPointer.Set(reflect.ValueOf(router.db))
		}
		method := controller.MethodByName(action)
		if method.Interface() == nil {
			method = controller.MethodByName("NotFound")
		}
		in := make([]reflect.Value, method.Type().NumIn())
		in[0] = reflect.ValueOf(r)
		response, err, responseCode := router.prepareResponse(method.Call(in))
		if err != nil {
			router.errorResponse(err, w)
			return
		}
		jsonResponse, jsonErr := json.Marshal(response)
		if jsonErr != nil {
			router.errorResponse(jsonErr, w)
			return
		}
		w.WriteHeader(responseCode)
		w.Write(jsonResponse)
	}).Methods(method)
}

func (router *Router) prepareResponse(values []reflect.Value) (interface{}, error, int) {
	headerCode := 200
	var err error
	var response interface{}
	if len(values) > 2 {
		headerCode = int(values[2].Int())
	}
	if len(values) > 1 {
		err = values[1].Interface().(error)
	}
	if len(values) > 0 {
		response = values[0].Interface()
	}

	return response, err, headerCode
}

func (router *Router) errorResponse(err error, w http.ResponseWriter) {
	errorString := "Internal server error"
	if router.debugMode {
		errorString = err.Error()
	}
	response, err := json.Marshal(map[string]string{
		"error": errorString,
	})
	w.WriteHeader(500)
	w.Write(response)
}
