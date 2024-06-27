package http

import (
	"github.com/autokz/go-http-server-helper/v2/httpHelper"
	"hwCalendar/gateway/server/http/handler/event/add"
	"hwCalendar/gateway/server/http/handler/event/all"
	"hwCalendar/gateway/server/http/handler/event/byid"
	"hwCalendar/gateway/server/http/handler/event/deleteEvent"
	"hwCalendar/gateway/server/http/handler/event/update"
	"hwCalendar/gateway/server/http/handler/oauth/refresh"
	"hwCalendar/gateway/server/http/handler/oauth/signin"
	"hwCalendar/gateway/server/http/handler/oauth/signoutall"
	"hwCalendar/gateway/server/http/handler/oauth/signup"
	"hwCalendar/gateway/server/http/handler/user/updateCreds"
	"hwCalendar/gateway/server/http/middleware"
	"log"
	"net/http"
)

func InitServer() {
	router := httpHelper.NewRouter()
	api := router.NewGroupRoute("/api")

	v1 := api.NewGroupRoute("/v1", httpHelper.JsonMiddleware)

	oauth := v1.NewGroupRoute("/auth")
	oauth.Post("/signup", signup.Handle)
	oauth.Post("/signin", signin.Handle)
	oauth.Post("/refresh", refresh.Handle).Middleware(middleware.RequireAuth)
	oauth.Post("/signoutall", signoutall.Handle).Middleware(middleware.RequireAuth)

	event := v1.NewGroupRoute("/event", middleware.RequireAuth)
	event.Post("/add", add.Handle)
	event.Get("/byid", byid.Handle)
	event.Get("/all", all.Handle)
	event.Delete("/delete", deleteEvent.Handle)
	event.Put("/update", update.Handle)

	user := v1.NewGroupRoute("/user")
	user.Put("", updateCreds.Handle).Middleware(middleware.RequireAuth)

	log.Fatal(http.ListenAndServe(":8080", router.Mux()))
}
