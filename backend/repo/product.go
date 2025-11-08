package repo

import (
	"database/sql"
	"fmt"

	"ecommerce/domain"
	"ecommerce/product"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

// constructor or constructor function
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {

	//fmt.Println("product : ", p)

	query := `
	  INSERT INTO products (
       title,
	   description,
	   price,
	   img_url
     ) VALUES(
     $1,
	 $2,
	 $3,
	 $4
    )
	 RETURNING id
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)

	if err != nil {
		return nil, err
	}

	return &p, nil

}

func (r *productRepo) Get(id int) (*domain.Product, error) {
	var prd domain.Product

	query := `
	 SELECT 
	 id ,
	 title,
	 description,
	 price,
	 img_url
	 from products
	 where id = $1;
	 `
	err := r.db.Get(&prd, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}
	return &prd, nil

}

func (r *productRepo) List(page, limit int64) ([]*domain.Product, error) {

	offset := ((page - 1) * limit)

	var prdList []*domain.Product

	query := `
	 SELECT 
	  id ,
	  title,
	  description,
	  price,
	  img_url
	  from products
	  LIMIT $1
	  OFFSET $2;
	 `
	err := r.db.Select(&prdList, query, limit, offset)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return prdList, nil
}

func (r *productRepo) Count() (int64, error) {

	query := `
	 SELECT 
	   COUNT(*)
	  from products;
	 `
	var count int64
	err := r.db.QueryRow(query).Scan(&count)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return count, nil
}

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
	query := `
	 UPDATE products
	   SET 
		title = $2,
		description = $3,
		price = $4,
		img_url = $5
	   WHERE id = $1
	   RETURNING id, title, description, price, img_url;
	`

	row := r.db.QueryRow(query, p.ID, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.ImgUrl)

	if err != nil {
		fmt.Println("show the error update not successfull : ", err)
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Delete(id int) error {
	query := `
	 DELETE FROM products WHERE id =$1 
	`
	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
