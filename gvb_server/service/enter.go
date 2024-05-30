package service

import (
	image_service "gvb_server/service/image_service"
)

type ServiceGroup struct {
	ImageService image_service.ImageService
}

var Services = new(ServiceGroup)