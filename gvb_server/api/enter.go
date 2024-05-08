package api

<<<<<<< HEAD
import (
	images_api "gvb_server/api/images_api"
	settings_api "gvb_server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
=======
import settings_api "gvb_server/api/settings_api"

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
>>>>>>> 2f9e4d1a6a0ab0002a002517dace0301441cd6ca
}

var ApiGroupApp = new(ApiGroup)
