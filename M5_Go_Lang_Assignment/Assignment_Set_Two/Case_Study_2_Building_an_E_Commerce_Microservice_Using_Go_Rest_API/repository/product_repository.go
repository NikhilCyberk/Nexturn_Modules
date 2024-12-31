package repository

import (
	"Case_Study_2_Building_an_E_Commerce_Microservice_Using_Go_Rest_API/model"
	"database/sql"
	"errors"
	"time"
)

type ProductRepository struct {
    db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
    return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *model.Product) error {
    stmt, err := r.db.Prepare(`
        INSERT INTO products (name, description, price, stock, category_id)
        VALUES (?, ?, ?, ?, ?)
    `)
    if err != nil {
        return err
    }
    defer stmt.Close()

    result, err := stmt.Exec(
        product.Name,
        product.Description,
        product.Price,
        product.Stock,
        product.CategoryID,
    )
    if err != nil {
        return err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return err
    }

    product.ID = int(id)
    return nil
}

func (r *ProductRepository) GetByID(id int) (*model.Product, error) {
    var product model.Product
    err := r.db.QueryRow(`
        SELECT id, name, description, price, stock, category_id, created_at, updated_at 
        FROM products WHERE id = ?
    `, id).Scan(
        &product.ID,
        &product.Name,
        &product.Description,
        &product.Price,
        &product.Stock,
        &product.CategoryID,
        
    )
    if err == sql.ErrNoRows {
        return nil, errors.New("product not found")
    }
    if err != nil {
        return nil, err
    }
    return &product, nil
}

func (r *ProductRepository) GetAll(page, limit int) ([]model.Product, int, error) {
    offset := (page - 1) * limit

    // Get total count
    var total int
    err := r.db.QueryRow("SELECT COUNT(*) FROM products").Scan(&total)
    if err != nil {
        return nil, 0, err
    }

    // Get paginated results
    rows, err := r.db.Query(`
        SELECT id, name, description, price, stock, category_id, created_at, updated_at 
        FROM products LIMIT ? OFFSET ?
    `, limit, offset)
    if err != nil {
        return nil, 0, err
    }
    defer rows.Close()

    var products []model.Product
    for rows.Next() {
        var p model.Product
        if err := rows.Scan(
            &p.ID,
            &p.Name,
            &p.Description,
            &p.Price,
            &p.Stock,
            &p.CategoryID,
        ); err != nil {
            return nil, 0, err
        }
        products = append(products, p)
    }

    return products, total, nil
}

func (r *ProductRepository) Update(product *model.Product) error {
    stmt, err := r.db.Prepare(`
        UPDATE products 
        SET name = ?, description = ?, price = ?, stock = ?, category_id = ?
        WHERE id = ?
    `)
    if err != nil {
        return err
    }
    defer stmt.Close()

    result, err := stmt.Exec(
        product.Name,
        product.Description,
        product.Price,
        product.Stock,
        product.CategoryID,
        time.Now(),
        product.ID,
    )
    if err != nil {
        return err
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rows == 0 {
        return errors.New("product not found")
    }

    return nil
}

func (r *ProductRepository) UpdateStock(id, stock int) error {
    result, err := r.db.Exec(
        "UPDATE products SET stock = ?, WHERE id = ?",
        stock,
        time.Now(),
        id,
    )
    if err != nil {
        return err
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rows == 0 {
        return errors.New("product not found")
    }

    return nil
}

func (r *ProductRepository) Delete(id int) error {
    result, err := r.db.Exec("DELETE FROM products WHERE id = ?", id)
    if err != nil {
        return err
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rows == 0 {
        return errors.New("product not found")
    }

    return nil
}