package models

import (
	"errors"

	"github.com/paramahastha/pretest/api/db"
	"github.com/paramahastha/pretest/api/forms"
)

//Product ...
type Product struct {
	ID          int64  `db:"id, primarykey, autoincrement" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Enable      *bool  `db:"enable" json:"enable"`
}

//ProductModel ...
type ProductModel struct{}

//Create ...
func (m ProductModel) Create(form forms.ProductForm) (err error) {
	getDb := db.GetDB()

	_, err = getDb.Exec("INSERT INTO public.product(name, description, enable) VALUES($1, $2, $3) RETURNING id", form.Name, form.Description, form.Enable)
	// fmt.Printf(result)
	if err != nil {
		return err
	}

	return err
}

//One ...
func (m ProductModel) One(id int64) (product Product, err error) {
	err = db.GetDB().SelectOne(&product, "SELECT * FROM public.product WHERE id=$1 LIMIT 1", id)
	return product, err
}

//All ...
func (m ProductModel) All() (products []Product, err error) {
	_, err = db.GetDB().Select(&products, "SELECT * FROM public.product ORDER BY name ASC")

	return products, err
}

//Update ...
func (m ProductModel) Update(id int64, form forms.ProductForm) (err error) {
	_, err = m.One(id)

	if err != nil {
		return errors.New("Product not found")
	}

	_, err = db.GetDB().Exec("UPDATE public.product SET name=$1, description=$2, enable=$3 WHERE id=$4", form.Name, form.Description, form.Enable, id)

	return err
}

//Delete ...
func (m ProductModel) Delete(id int64) (err error) {
	_, err = m.One(id)

	if err != nil {
		return errors.New("Product not found")
	}

	_, err = db.GetDB().Exec("DELETE FROM public.product WHERE id=$1", id)

	return err
}
