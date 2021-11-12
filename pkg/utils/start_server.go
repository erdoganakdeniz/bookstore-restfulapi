package utils

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
)

func StartServerWithGracefulShutdown(a *fiber.App) {
	idleConnsClosed:=make(chan struct{})

	go func() {
		sigint:=make(chan os.Signal,1)
		signal.Notify(sigint,os.Interrupt)
		<-sigint
		if err:=a.Shutdown();err != nil {
			log.Printf("Sunucu duraklatıldı ! Hata : %v",err)
		}
		close(idleConnsClosed)
	}()

	if err:=a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Sunucu başlatılamadı ! Hata : %v",err)
	}
	<-idleConnsClosed
}
func StartServer(a *fiber.App) {
	if err:=a.Listen(os.Getenv("SERVER_URL"));err!=nil {
		log.Printf("Sunucu başlatılamadı ! Hata : %v",err)
	}
}
