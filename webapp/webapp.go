package webapp

import (
	"fmt"
	"net"

	"github.com/acidlemon/rocket"
	"github.com/mix3/go-rocket-sample-app2/webapp/controller"
)

type WebApp struct {
	rocket.WebApp
}

var view = &rocket.View{}

func New() WebApp {
	app := WebApp{}
	app.Init()

	app.AddRoute(
		"/upload/poling/:key",
		controller.UploadPolingPage,
		view,
	)

	app.AddRoute(
		"/upload/sync",
		controller.UploadSyncPage,
		view,
	)

	app.AddRoute(
		"/upload/async",
		controller.UploadAsyncPage,
		view,
	)

	app.AddRoute(
		"/upload",
		controller.UploadPage,
		view,
	)

	app.AddRoute(
		"/",
		controller.TopPage,
		view,
	)

	app.AddRoute(
		"/favicon.ico",
		controller.TopPage,
		view,
	)

	app.BuildRouter()
	return app
}

func Start(listener net.Listener) {
	app := New()
	fmt.Println("Launch succeeded!")
	app.Start(listener)
}
