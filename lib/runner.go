package lib

import (
	"fmt"
	sakura "github.com/yamamoto-febc/sakura-iot-go"
	"log"
	"net/http"
	"os"
)

func Start(option *Option) error {

	option.PrintOptionValues()

	log.SetFlags(log.Ldate | log.Ltime)
	log.SetOutput(os.Stdout)
	out := log.Printf

	balusHandler := NewWebhookHandler(option, out)

	handler := &sakura.WebhookHandler{
		Secret:     option.ReceiveSecret,
		HandleFunc: balusHandler.HandleRequest,
		Debug:      option.Debug,
	}

	addr := fmt.Sprintf("%s:%d", "", option.ReceivePort)

	out("[INFO] start ListenAndServe. addr:[%s] path:[%s] secret:[%s]\n", addr, option.ReceivePath, option.ReceiveSecret)
	http.Handle(option.ReceivePath, handler)
	return http.ListenAndServe(addr, nil)
}
