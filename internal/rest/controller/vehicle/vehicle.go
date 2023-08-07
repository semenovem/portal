package vehicle

import (
	"github.com/labstack/echo/v4"
)

// Search docs
//
//	@Summary	Ищет автомобили по фильтру
//	@Description
//	@Produce	json
//	@Param		payload	query		SearchForm	false	"Параметры запроса. Все поля опциональные"
//	@Success	200		{object}	ListResponse
//	@Failure	400		{object}	failing.Response
//	@Router		/vehicles [GET]
//	@Tags		vehicles
//	@Security	ApiKeyAuth
func (ct *Controller) Search(c echo.Context) error {
	return nil
}

// Get docs
//
//	@Summary	Получает данные автомобиля по ID
//	@Description
//	@Produce	json
//	@Param		vehicle_id	path		uint32	true	"ID автомобиля"
//	@Success	200			{object}	view.Vehicle
//	@Failure	400			{object}	failing.Response
//	@Router		/vehicles/:vehicle_id [GET]
//	@Tags		vehicles
//	@Security	ApiKeyAuth
func (ct *Controller) Get(c echo.Context) error {
	return nil
}

// Upd docs
//
//	@Summary	Обновляет данные автомобиля
//	@Description
//	@Produce	json
//	@Param		vehicle_id	path		uint32	true	"ID автомобиля"
//	@Success	200			{object}	view.Vehicle
//	@Failure	400			{object}	failing.Response
//	@Router		/vehicles/:vehicle_id [PUT]
//	@Tags		vehicles
//	@Security	ApiKeyAuth
func (ct *Controller) Upd(c echo.Context) error {
	return nil
}

// Create   docs
//
//	@Summary	Создает новый автомобиль
//	@Description
//	@Produce	json
//	@Param		vehicle_id	path		uint32	true	"ID автомобиля"
//	@Success	201			{object}	view.Vehicle
//	@Failure	400			{object}	failing.Response
//	@Router		/vehicles/:vehicle_id [POST]
//	@Tags		vehicles
//	@Security	ApiKeyAuth
func (ct *Controller) Create(c echo.Context) error {
	return nil
}

// Del   docs
//
//	@Summary	Удаляет автомобиля
//	@Description
//	@Produce	json
//	@Param		vehicle_id	path		uint32	true	"ID автомобиля"
//	@Success	204			"no content"
//	@Failure	400			{object}	failing.Response
//	@Router		/vehicles/:vehicle_id [POST]
//	@Tags		vehicles
//	@Security	ApiKeyAuth
func (ct *Controller) Del(c echo.Context) error {
	return nil
}
