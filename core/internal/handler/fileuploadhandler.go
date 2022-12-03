package handler

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		// 判断是否已达用户容量上限
		// userIdentity := r.Header.Get("UserIdentity")
		// ub := new(models.UserBasic)
		// has, err := svcCtx.Engine.Where("identity = ?", userIdentity).Select("now_volume, total_volume").Get(ub)
		// if err != nil {
		// 	httpx.Error(w, err)
		// 	return
		// }
		// if fileHeader.Size+ub.NowVolume > ub.TotalVolume {
		// 	httpx.Error(w, errors.New("已超出当前容量"))
		// 	return
		// }
		// 判断文件是否存在
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)
		has, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		fmt.Println("2\n")

		if has {
			httpx.OkJson(w, &types.FileUploadReply{Identity: rp.Identity})
			return
		}

		cosPath, err := helper.CosUpload(r)
		if err != nil {
			return
		}
		fmt.Println(cosPath, "cos upload\n")
		// 判断使用的存储引擎，默认使用COS
		// if define.ObjectStorageType == "minio" {
		// 	filePath, err = helper.MinIOUpload(r)
		// } else {
		// 	filePath, err = helper.CosUpload(r)
		// }
		// if err != nil {
		// 	httpx.Error(w, err)
		// 	return
		// }

		// 往 logic 中传递 request
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = cosPath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
