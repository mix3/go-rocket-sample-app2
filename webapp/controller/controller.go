package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"code.google.com/p/go-uuid/uuid"

	"github.com/acidlemon/rocket"
)

func TopPage(ctx rocket.CtxData) {
	ctx.Res().StatusCode = http.StatusOK
	ctx.Render("template/top.html", rocket.RenderVars{})
}

func UploadPage(ctx rocket.CtxData) {
	ctx.Res().StatusCode = http.StatusOK
	ctx.Render("template/upload.html", rocket.RenderVars{})
}

func UploadAsyncPage(ctx rocket.CtxData) {
	key, err := uploadStart()
	if err != nil {
		ctx.Res().StatusCode = http.StatusInternalServerError
		ctx.RenderText(err.Error())
		return
	}
	go uploading(key)
	ctx.Res().StatusCode = http.StatusOK
	ctx.Render("template/async.html", rocket.RenderVars{
		"Key": key,
	})
}

func UploadSyncPage(ctx rocket.CtxData) {
	key, err := uploadStart()
	if err != nil {
		ctx.Res().StatusCode = http.StatusInternalServerError
		ctx.RenderText(err.Error())
		return
	}
	uploading(key)
	ctx.Res().StatusCode = http.StatusOK
	ctx.Render("template/poling.html", rocket.RenderVars{
		"Key": key,
	})
}

func UploadPolingPage(ctx rocket.CtxData) {
	key := ctx.Params().Get("key")
	var err error
	for {
		if exists(dir + key + "_tmp") {
			// uploading
		} else if exists(dir + key) {
			// complete
			break
		} else {
			// invalid key
			err = fmt.Errorf("invalid key")
			break
		}
	}
	if err != nil {
		ctx.Res().StatusCode = http.StatusInternalServerError
		ctx.RenderText(err.Error())
		return
	}
	ctx.Res().StatusCode = http.StatusOK
	ctx.Render("template/poling.html", rocket.RenderVars{
		"Key": key,
	})
}

var dir = "tmp/"

func uploadStart() (string, error) {
	key := uuid.New()
	err := ioutil.WriteFile(dir+key+"_tmp", []byte{}, 0644)
	return key, err
}

func uploading(key string) {
	time.Sleep(time.Second * 10)
	os.Remove(dir + key + "_tmp")
	ioutil.WriteFile(dir+key, []byte{}, 0644)
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
