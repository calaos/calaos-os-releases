package models

import (
	"errors"

	"gorm.io/gorm"
)

/*
{
    "images": [
        {
            "name": "calaos_home",
            "image": "ghcr.io/calaos/calaos_home:4.2.6",
            "version": "4.2.6"
        },
        {
            "name": "calaos_base",
            "image": "ghcr.io/calaos/calaos_base:4.8.1",
            "version": "4.8.1"
        }
    ]
}
*/

var (
	Release     = "release"
	Development = "dev"
)

type Image struct {
	ID       uint   `gorm:"primarykey" json:"-"`
	Name     string `json:"name" gorm:"index:idx_name,unique"`
	Source   string `json:"source"`
	Version  string `json:"version"`
	RepoType string `json:"-"`
}

func GetAllImages(repoType string) (imgs []*Image, err error) {
	err = db.Where("repo_type = ?", repoType).Find(&imgs).Error

	return
}

func GetImage(name string, repoType string) (img *Image, err error) {
	err = db.Where("name = ? AND repo_type = ?", name, repoType).First(&img).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return
}

func StoreImage(img *Image) (err error) {
	i, err := GetImage(img.Name, img.RepoType)
	if err != nil {
		return err
	}

	if i != nil {
		i.Source = img.Source
		i.Version = img.Version
		err = db.Save(i).Error
		return
	}

	err = db.Create(img).Error
	return
}
