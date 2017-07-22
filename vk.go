package vkcover

import (
	"io"
	"strconv"

	"bitbucket.org/PeterHueter/go-vk"
)

// Upload загружает изображение image шириной w и высотой h
// в группу gid, при этом токен может быть как токен пользователя
// или как токен групппы с правами photo
func Upload(gid int64, token string, img io.Reader, w, h int) error {
	// Получить адрес сервера загрузки
	p := vk.Params{
		"access_token": token,
		"group_id":     strconv.FormatInt(gid, 10),
		"crop_x2":      strconv.Itoa(w),
		"crop_y2":      strconv.Itoa(h),
	}

	url, err := vk.Photos_getOwnerCoverPhotoUploadServer(p)
	if err != nil {
		return err
	}

	// загрузить
	cp, err := vk.Photos_UploadCoverImage(url, img)
	if err != nil {
		return err
	}

	p = vk.Params{
		"access_token": token,
		"hash":         cp.Hash,
		"photo":        cp.Photo,
	}
	_, err = vk.Photos_saveOwnerCoverPhoto(p)
	return err
}
