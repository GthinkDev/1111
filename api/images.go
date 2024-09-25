package api

import (
	"fmt"
	"gin-blog/global"
	"gin-blog/models/res"
	"github.com/gin-gonic/gin"
	"path"
)

type ImagesApi struct {
}

type FileUploadRes struct {
	FileName  string `json:"filename"`
	Msg       string `json:"msg"`
	IsSuccess bool   `json:"is_success"`
}

func (ImagesApi) ImagesUploadView(c *gin.Context) {
	// 获取上传的文件
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	// 判断路径是否存在
	files, ok := form.File["image"]
	//fileHeader, err := c.FormFile("image")
	if !ok {
		res.FailWithMessage("文件不存在!", c)
		return
	}

	// 不存在就创建文件
	var resList []FileUploadRes

	for _, file := range files {
		filePath := path.Join("uploads", file.Filename)
		// 判断大小
		fileSize := float64(file.Size) / float64(1024*1024)
		if fileSize >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadRes{
				FileName:  file.Filename,
				Msg:       fmt.Sprintf("图片大小超过限制大小%dMB,当前文件大小为:%.2fMB", global.Config.Upload.Size, fileSize),
				IsSuccess: false,
			})
			continue
		}

		resList = append(resList, FileUploadRes{
			FileName:  filePath,
			Msg:       "上传成功",
			IsSuccess: true,
		})
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			continue
		}
	}
	res.OkWithData(resList, c)

}
