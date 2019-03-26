package models

import (
	"errors"

	"github.com/paramahastha/pretest/api/db"
	"github.com/paramahastha/pretest/api/forms"
)

//Category ...
type Category struct {
	ID     int64  `db:"id, primarykey, autoincrement" json:"id"`
	Name   string `db:"name" json:"name"`
	Enable *bool  `db:"enable" json:"enable"`
}

//CategoryModel ...
type CategoryModel struct{}

//Create ...
func (m CategoryModel) Create(form forms.CategoryForm) (err error) {
	getDb := db.GetDB()

	_, err = getDb.Exec("INSERT INTO public.category(name, enable) VALUES($1, $2) RETURNING id", form.Name, form.Enable)
	// fmt.Printf(result)
	if err != nil {
		return err
	}

	// articleID, err = getDb.SelectInt("SELECT id FROM public.article WHERE user_id=$1 ORDER BY id DESC LIMIT 1", userID)

	return err
}

//One ...
func (m CategoryModel) One(id int64) (category Category, err error) {
	err = db.GetDB().SelectOne(&category, "SELECT * FROM public.category WHERE id=$1 LIMIT 1", id)
	return category, err
}

//All ...
func (m CategoryModel) All() (categories []Category, err error) {
	_, err = db.GetDB().Select(&categories, "SELECT * FROM public.category ORDER BY name ASC")

	return categories, err
}

//Update ...
func (m CategoryModel) Update(id int64, form forms.CategoryForm) (err error) {
	_, err = m.One(id)

	if err != nil {
		return errors.New("Category not found")
	}

	_, err = db.GetDB().Exec("UPDATE public.category SET name=$1, enable=$2 WHERE id=$3", form.Name, form.Enable, id)

	return err
}

//Delete ...
func (m CategoryModel) Delete(id int64) (err error) {
	_, err = m.One(id)

	if err != nil {
		return errors.New("Category not found")
	}

	_, err = db.GetDB().Exec("DELETE FROM public.category WHERE id=$1", id)

	return err
}
