package service

import (
	image_service "gvb_server/service/image_service"
	 "gvb_server/service/user_service"
)

type ServiceGroup struct {
	ImageService image_service.ImageService
	UserService  user_service.UserService
}

var Services = new(ServiceGroup)
