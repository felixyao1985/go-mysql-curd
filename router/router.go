package router

import "restfulApi/lib/httprouter"

//import (
//	"restfulApi/app"
//	"restfulApi/lib/httprouter"
//)
//
//func NewRouter() *httprouter.Router {
//
//	router := httprouter.New()
//	routes := app.GetRoutes()
//	for _, route := range routes {
//		router.Handle(route.Method, route.Pattern, route.Handle)
//	}
//	return router
//}

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handle  httprouter.Handle
}

type Routes []Route