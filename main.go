package main

import (
    Main "blog/routers"

    "net/http"
	"fmt"
	"blog/pkg/setting"
)

func main() {
    router := Main.InitRouter()

    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeout,
        WriteTimeout:   setting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }

    s.ListenAndServe()
}
