package router

import (
	"errors"
	"net/http"
)

func (r *Router) Start() {
	go func() {
		<-r.ctx.Done()
		if err := r.server.Close(); err != nil {
			r.logger.Named("Close").Error(err.Error())
		}
	}()

	r.logger.Infof("router start on %s", r.addr)

	r.server.HidePort = true
	r.server.HideBanner = true

	if err := r.server.Start(r.addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
		r.logger.Named("Start").Error(err.Error())
	}
}
