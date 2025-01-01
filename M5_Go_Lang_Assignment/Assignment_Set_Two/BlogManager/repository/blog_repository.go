package repository

import (
	"blogmanager/model"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type BlogRepository struct {
	DB *sql.DB
}

func NewBlogRepository(db *sql.DB) *BlogRepository {
	return &BlogRepository{DB: db}
}

func (repo *BlogRepository) CreateBlog(blog *model.Blog) (*model.Blog, error) {
	stmt, err := repo.DB.Prepare("INSERT INTO blogs (title, content, author, timestamp) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(blog.Title, blog.Content, blog.Author, time.Now().String())
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err

	}

	blog.ID = int(id)
	return blog, nil
}

func (repo *BlogRepository) GetBlog(id int) (*model.Blog, error) {
	row := repo.DB.QueryRow("SELECT id, title, content, author, timestamp FROM blogs WHERE id = ?", id)
	blog := &model.Blog{}
	err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.TimeStamp)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (repo *BlogRepository) GetAllBlogs() ([]model.Blog, error) {
	rows, err := repo.DB.Query("SELECT id, title, content, author, timestamp FROM blogs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []model.Blog
	for rows.Next() {
		var blog model.Blog
		err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.TimeStamp)
		if err != nil {
			log.Fatal(err)
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func (repo *BlogRepository) UpdateBlog(blog *model.Blog) (*model.Blog, error) {
	stmt, err := repo.DB.Prepare("UPDATE blogs SET title = ?, content = ?, author = ?, timestamp = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(blog.Title, blog.Content, blog.Author, time.Now().String(), blog.ID)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully updated blog with ID:", blog.ID)
	return blog, nil
}

func (repo *BlogRepository) DeleteBlog(id int) error {
	stmt, err := repo.DB.Prepare("DELETE FROM blogs WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	fmt.Println("Successfully deleted blog with ID:", id)
	return nil
}
