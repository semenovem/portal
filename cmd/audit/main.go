package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v9"
	"github.com/semenovem/portal/config"
	auditApp "github.com/semenovem/portal/internal/app/audit"
	"github.com/semenovem/portal/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		sig         = make(chan os.Signal)
		ll, setter  = logger.New()
		cfg         config.Audit
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

	if err := env.Parse(&cfg); err != nil {
		ll.Named("env.Parse").Errorf("can't parse env: ", err)
		cancel()
		return
	}

	setter.SetCli(true)
	setter.SetShowTime(true)
	setter.SetLevel(cfg.Base.LogLevel)

	if err := auditApp.New(ctx, ll, cfg); err != nil {
		ll.Named("create.app").Nested(err.Error())
		//cancel()
	}

	<-ctx.Done()
}
