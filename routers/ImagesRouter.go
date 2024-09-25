package routers

import "gin-blog/api"

func (r *RouterGroup) ImagesRouter() {
	images := api.ApiGroup.Images
	r.POST("images", images.ImagesUploadView)
}
