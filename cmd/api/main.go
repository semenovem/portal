package main

import (
	"context"
	"fmt"
	"github.com/semenovem/portal/config"
	apiApp "github.com/semenovem/portal/internal/app/platform"
	"github.com/semenovem/portal/internal/router"
	"github.com/semenovem/portal/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
	//"github.com/semenovem/portal/internal/api"
)

//	@title			portal API
//	@version		1.0
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	semenovem@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/[v1]/

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Description for what is this security definition being used
func main() {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		sig         = make(chan os.Signal)
		ll, setter  = logger.New()
	)

	defer func() {
		cancel()
		ll.Info("exiting")
		fmt.Println("exiting")
		time.Sleep(time.Millisecond * 500)
	}()

	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		cancel()
	}()

	cfg, err := config.ParseAPI()
	if err != nil {
		ll.Named("env.Parse").Errorf("can't parse env: ", err)
		cancel()
		return
	}

	setter.SetCli(true)
	setter.SetShowTime(true)
	setter.SetLevel(cfg.Base.LogLevel)
	setter.SetRequestIDExtractor(router.ExtractRequestID)

	if err = apiApp.New(ctx, ll, cfg); err != nil {
		_ = ll.NestedWith(err, "can't start app")
		//cancel()
	}

	<-ctx.Done()

}
