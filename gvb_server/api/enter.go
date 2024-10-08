package api

import (
	advert_api "gvb_server/api/advert_api"
	images_api "gvb_server/api/images_api"
	"gvb_server/api/menu_api"
	settings_api "gvb_server/api/settings_api"
	"gvb_server/api/tag_api"
	"gvb_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
	TagApi      tag_api.TagApi
}

var ApiGroupApp = new(ApiGroup)
