package ping

import (
	"context"
	"log"
	"net/http"
)

var srv = http.Server{Addr: ":80"}

func StartPingServer() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[HTTP]accept http request,host:%v,remote addr:%v", r.Host, r.RemoteAddr)
		w.Write([]byte("pong"))
	})
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("[Error]Listen and serve:%v", err)
		}
	}()
	log.Printf("Ping HTTP Server start successfully")
}

func StopPingServer(ctx context.Context) {
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("[Error]Fail to stop HTTP server gracefully:%v", err)
	}
}
