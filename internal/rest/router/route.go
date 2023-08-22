package router

func (r *Router) addRoutes() {
	// Авторизация
	r.unauth.POST("/auth/login", r.authCnt.Login)
	r.unauth.POST("/auth/logout", r.authCnt.Logout)
	r.unauth.POST("/auth/refresh", r.authCnt.Refresh)
	r.auth.POST("/auth/onetime", r.authCnt.CreateOnetimeLink)
	r.unauth.POST("/auth/onetime/:entry_id", r.authCnt.LoginOnetimeLink)

	// People
	r.auth.POST("/people", r.peopleCnt.CreateUser)

	r.auth.GET("/people/self/profile", r.peopleCnt.SelfProfile)
	
	r.auth.GET("/people/:user_id/profile", r.peopleCnt.Profile)
	r.auth.DELETE("/people/:user_id", r.peopleCnt.DeleteUser)

	// People Position
	r.auth.GET("/people/positions", r.peopleCnt.Profile)

	// Store
	r.auth.GET("/store/:store_path", r.storeCnt.Load)
	r.auth.POST("/store/:store_path", r.storeCnt.Store)
	r.auth.DELETE("/store/:store_path", r.storeCnt.Delete)

	//
	r.auth.GET("/vehicles", r.vehicleCnt.Search)
	r.auth.GET("/vehicles/:vehicle_id", r.vehicleCnt.Get)
	r.auth.POST("/vehicles/:vehicle_id", r.vehicleCnt.Create)
	r.auth.PUT("/vehicles/:vehicle_id", r.vehicleCnt.Upd)
	r.auth.DELETE("/vehicles/:vehicle_id", r.vehicleCnt.Del)

	//
}
