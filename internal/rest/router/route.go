package router

func (r *Router) addRoutes() {
	var (
		unauth = r.unauth
		auth   = r.auth
	)

	// ----------------- Авторизация -----------------
	unauth.POST("/auth/login", r.authCnt.Login)
	unauth.POST("/auth/logout", r.authCnt.Logout)
	unauth.POST("/auth/refresh", r.authCnt.Refresh)
	unauth.POST("/auth/onetime/:entry_id", r.authCnt.LoginOnetimeLink)
	auth.POST("/auth/onetime", r.authCnt.CreateOnetimeLink)

	// ----------------- Media -----------------
	auth.POST("/media/file", r.mediaCnt.FileUpload)
	auth.GET("/media/file/:file_id", r.mediaCnt.FileUpload)
	auth.DELETE("/media/file/:file_id", r.mediaCnt.FileUpload)

	// Media box
	auth.POST("/media/box", r.mediaCnt.FileUpload)
	auth.GET("/media/box/:box_id", r.mediaCnt.FileUpload)
	auth.PUT("/media/box/:box_id", r.mediaCnt.FileUpload)
	//auth.DELETE("/media/box/:box_id", r.mediaCnt.FileUpload)

	// ----------------- People -----------------
	auth.POST("/people", r.peopleCnt.CreateUser)

	auth.GET("/people/self/profile", r.peopleCnt.SelfProfile)
	auth.GET("/people/:user_id/profile", r.peopleCnt.Profile)

	auth.DELETE("/people/:user_id", r.peopleCnt.DeleteUser)

	// People Position
	auth.GET("/people/positions", r.peopleCnt.Profile)

	// ----------------- Store -----------------
	auth.GET("/store/:store_path", r.storeCnt.Load)
	auth.POST("/store/:store_path", r.storeCnt.Store)
	auth.DELETE("/store/:store_path", r.storeCnt.Delete)

	// ----------------- Vehicle -----------------
	auth.GET("/vehicles", r.vehicleCnt.Search)
	auth.GET("/vehicles/:vehicle_id", r.vehicleCnt.Get)
	auth.POST("/vehicles/:vehicle_id", r.vehicleCnt.Create)
	auth.PUT("/vehicles/:vehicle_id", r.vehicleCnt.Upd)
	auth.DELETE("/vehicles/:vehicle_id", r.vehicleCnt.Del)

	//
}
