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

	// ----------------- Store -----------------
	auth.GET("/store/:store_path", r.storeCnt.Load)
	auth.POST("/store/:store_path", r.storeCnt.Store)
	auth.DELETE("/store/:store_path", r.storeCnt.Delete)

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
	auth.DELETE("/people/:user_id", r.peopleCnt.DeleteUser)

	// Employee
	auth.POST("/people/employee", r.peopleCnt.CreateEmployee)
	auth.PATCH("/people/employee/:user_id", r.peopleCnt.UpdateEmployee)

	auth.GET("/people/self/profile", r.peopleCnt.SelfProfile)
	auth.GET("/people/:user_id/profile", r.peopleCnt.UserProfile)
	auth.GET("/people/:user_id/profile/public", r.peopleCnt.UserPublicProfile)

	auth.GET("/people/free-login/:login_name", r.peopleCnt.CheckLogin)

	unauth.GET("/people/handbook", r.peopleCnt.Handbook)

	// People Position
	auth.GET("/people/position", r.peopleCnt.UserPublicProfile)
	auth.GET("/people/position/:position_id", r.peopleCnt.UserPublicProfile)

	// People Dept
	auth.GET("/people/dept", r.peopleCnt.UserPublicProfile)
	auth.GET("/people/dept/:dept_id", r.peopleCnt.UserPublicProfile)

	// ----------------- Vehicle -----------------
	auth.GET("/vehicles", r.vehicleCnt.Search)
	auth.GET("/vehicles/:vehicle_id", r.vehicleCnt.Get)
	auth.POST("/vehicles/:vehicle_id", r.vehicleCnt.Create)
	auth.PUT("/vehicles/:vehicle_id", r.vehicleCnt.Upd)
	auth.DELETE("/vehicles/:vehicle_id", r.vehicleCnt.Del)

	//
}
