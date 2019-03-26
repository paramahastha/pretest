package models

import (
	"errors"

	"github.com/paramahastha/pretest/api/db"
	"github.com/paramahastha/pretest/api/forms"
)

//Image ...
type Image struct {
	ID     int64  `db:"id, primarykey, autoincrement" json:"id"`
	Name   string `db:"name" json:"name"`
	File   string `db:"file" json:"file"`
	Enable *bool  `db:"enable" json:"enable"`
}

//ImageModel ...
type ImageModel struct{}

//Create ...
func (m ImageModel) Create(form forms.ImageForm) (err error) {
	getDb := db.GetDB()

	_, err = getDb.Exec("INSERT INTO public.image(name, file, enable) VALUES($1, $2, $3) RETURNING id", form.Name, form.File, form.Enable)

	if err != nil {
		return err
	}

	return err
}

//One ...
func (m ImageModel) One(id int64) (image Image, err error) {
	err = db.GetDB().SelectOne(&image, "SELECT * FROM public.image WHERE id=$1 LIMIT 1", id)
	return image, err
}

//All ...
func (m ImageModel) All() (images []Image, err error) {
	_, err = db.GetDB().Select(&images, "SELECT * FROM public.image ORDER BY name ASC")

	return images, err
}

//Update ...
func (m ImageModel) Update(id int64, form forms.ImageForm) (err error) {
	_, err = m.One(id)

	if err != nil {
		return errors.New("Image not found")
	}

	_, err = db.GetDB().Exec("UPDATE public.image SET name=$1, file=$2, enable=$3 WHERE id=$4", form.Name, form.File, form.Enable, id)

	return err
}

//Delete ...
func (m ImageModel) Delete(id int64) (err error) {
	_, err = m.One(id)

	if err != nil {
		return errors.New("Image not found")
	}

	_, err = db.GetDB().Exec("DELETE FROM public.image WHERE id=$1", id)

	return err
}
