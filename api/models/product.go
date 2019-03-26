package models

import (
	"errors"
	"fmt"

	"github.com/paramahastha/pretest/api/db"
	"github.com/paramahastha/pretest/api/forms"
)

// Product ...
type Product struct {
	ID          int64  `db:"id, primarykey, autoincrement" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Enable      *bool  `db:"enable" json:"enable"`
	Category    string `db:"category" json:"category"`
	Image       string `db:"image" json:"image"`
}

//ProductModel ...
type ProductModel struct{}

//Create ...
func (m ProductModel) Create(form forms.ProductForm) (err error) {
	getDb := db.GetDB()

	lastID := 0

	err = getDb.QueryRow("INSERT INTO public.product(name, description, enable) VALUES($1, $2, $3) RETURNING id", form.Name, form.Description, form.Enable).Scan(&lastID)

	if err != nil {
		return err
	}

	_, err = getDb.Exec("INSERT INTO public.category_product(product_id, category_id) VALUES($1, $2)", lastID, form.CategoryID)

	if err != nil {
		return err
	}

	_, err = getDb.Exec("INSERT INTO public.product_image(product_id, image_id) VALUES($1, $2)", lastID, form.ImageID)
	fmt.Println("LOG", form.ImageID)
	if err != nil {
		return err
	}

	return err
}

//One ...
func (m ProductModel) One(id int64) (product Product, err error) {
	err = db.GetDB().SelectOne(&product, `SELECT p.id, p.name, p.description, p.enable, 
	string_agg(DISTINCT c.name, ',')  AS category , 
	string_agg(DISTINCT i.file, ',') AS image 
	FROM product p 
	INNER JOIN category_product cp ON p.id = cp.product_id 
	INNER JOIN product_image pi ON p.id = pi.product_id 
	INNER JOIN category c ON cp.category_id = c.id
	INNER JOIN image i ON pi.image_id = i.id WHERE p.id=$1 GROUP BY p.id LIMIT 1`, id)
	return product, err
}

//All ...
func (m ProductModel) All() (products []Product, err error) {
	_, err = db.GetDB().Select(&products, `SELECT p.id, p.name, p.description, p.enable, 
	string_agg(DISTINCT c.name, ',')  AS category , 
	string_agg(DISTINCT i.file, ',') AS image 
	FROM product p 
	INNER JOIN category_product cp ON p.id = cp.product_id 
	INNER JOIN product_image pi ON p.id = pi.product_id 
	INNER JOIN category c ON cp.category_id = c.id
	INNER JOIN image i ON pi.image_id = i.id GROUP BY p.id
	`)

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
