package repository

import (
	"Case_Study_1_Building_a_Blog_Management_System/model"
	"database/sql"
	"errors"
	"time"
)

type BlogRepository struct {
    DB *sql.DB
}

func NewBlogRepository(db *sql.DB) *BlogRepository {
    return &BlogRepository{DB: db}
}

func (r *BlogRepository) CreateBlog(blog *model.Blog) (*model.Blog, error) {
    stmt, err := r.DB.Prepare("INSERT INTO blogs (title, content, author, timestamp) VALUES (?, ?, ?, ?)")
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    result, err := stmt.Exec(blog.Title, blog.Content, blog.Author, time.Now().String())
    if err != nil {
        return nil, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return nil, err
    }

    blog.ID = int(id)
    return blog, nil
}

func (r *BlogRepository) GetBlog(id int) (*model.Blog, error) {
    var blog model.Blog
    err := r.DB.QueryRow("SELECT id, title, content, author, timestamp FROM blogs WHERE id = ?", id).
        Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
    if err == sql.ErrNoRows {
        return nil, errors.New("blog not found")
    }
    if err != nil {
        return nil, err
    }
    return &blog, nil
}

func (r *BlogRepository) GetAllBlogs() ([]model.Blog, error) {
    rows, err := r.DB.Query("SELECT id, title, content, author, timestamp FROM blogs ORDER BY timestamp DESC")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var blogs []model.Blog
    for rows.Next() {
        var blog model.Blog
        if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp); err != nil {
            return nil, err
        }
        blogs = append(blogs, blog)
    }
    return blogs, nil
}

func (r *BlogRepository) UpdateBlog(id int, blog *model.Blog) (*model.Blog, error) {
    stmt, err := r.DB.Prepare("UPDATE blogs SET title = ?, content = ?, author = ? WHERE id = ?")
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    result, err := stmt.Exec(blog.Title, blog.Content, blog.Author, id)
    if err != nil {
        return nil, err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return nil, err
    }
    if rowsAffected == 0 {
        return nil, errors.New("blog not found")
    }

    blog.ID = id
    return blog, nil
}

func (r *BlogRepository) DeleteBlog(id int) error {
    result, err := r.DB.Exec("DELETE FROM blogs WHERE id = ?", id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rowsAffected == 0 {
        return errors.New("blog not found")
    }

    return nil
}