package router

func (r *Router) addRoutes() {
	// Авторизация
	r.unauth.POST("/auth/login", r.authCnt.Login)
	r.unauth.POST("/auth/logout", r.authCnt.Logout)
	r.unauth.POST("/auth/refresh", r.authCnt.Refresh)
	r.auth.POST("/auth/onetime", r.authCnt.CreateOnetimeLink)
	r.unauth.POST("/auth/onetime/:entry_id", r.authCnt.LoginOnetimeLink)

	//
	r.auth.GET("/vehicles", r.vehicleCnt.Search)

	r.auth.GET("/vehicles/:vehicle_id", r.vehicleCnt.Get)
	r.auth.POST("/vehicles/:vehicle_id", r.vehicleCnt.Create)
	r.auth.PUT("/vehicles/:vehicle_id", r.vehicleCnt.Upd)
	r.auth.DELETE("/vehicles/:vehicle_id", r.vehicleCnt.Del)

	//
}
