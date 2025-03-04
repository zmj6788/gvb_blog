package api

import (
	advert_api "gvb_server/api/advert_api"
	"gvb_server/api/article_api"
	"gvb_server/api/comment_api"
	"gvb_server/api/digg_api"
	images_api "gvb_server/api/images_api"
	"gvb_server/api/menu_api"
	"gvb_server/api/message_api"
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
	MessageApi  message_api.MessageApi
	ArticleApi  article_api.ArticleApi
	DiggApi     digg_api.DiggApi
	CommentApi  comment_api.CommentApi
}

var ApiGroupApp = new(ApiGroup)
