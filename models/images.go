package models

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

type Image struct {
	ID      uint   `gorm:"primarykey" json:"-"`
	Name    string `json:"name" gorm:"index:idx_name,unique"`
	Source  string `json:"source"`
	Version string `json:"version"`
}

func GetAllImages() (imgs []*Image, err error) {
	err = db.Find(&imgs).Error

	return
}

func GetImage(name string) (img *Image, err error) {
	err = db.Where("name = ?", name).First(&img).Error
	return
}

func StoreImage(img *Image) (err error) {
	i, err := GetImage(img.Name)
	if err != nil {
		err = db.Save(i).Error
		return
	}

	err = db.Create(img).Error
	return
}
