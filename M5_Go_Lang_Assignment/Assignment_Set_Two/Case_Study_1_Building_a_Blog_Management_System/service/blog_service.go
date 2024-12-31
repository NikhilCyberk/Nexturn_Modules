package service

import (
	"Case_Study_1_Building_a_Blog_Management_System/model"
	"Case_Study_1_Building_a_Blog_Management_System/repository"
)

type BlogService struct {
    BlogRepository *repository.BlogRepository
}

func NewBlogService(blogRepository *repository.BlogRepository) *BlogService {
    return &BlogService{BlogRepository: blogRepository}
}

func (s *BlogService) CreateBlog(blog *model.Blog) (*model.Blog, error) {
    return s.BlogRepository.CreateBlog(blog)
}

func (s *BlogService) GetBlog(id int) (*model.Blog, error) {
    return s.BlogRepository.GetBlog(id)
}

func (s *BlogService) GetAllBlogs() ([]model.Blog, error) {
    return s.BlogRepository.GetAllBlogs()
}

func (s *BlogService) UpdateBlog(id int, blog *model.Blog) (*model.Blog, error) {
    return s.BlogRepository.UpdateBlog(id, blog)
}

func (s *BlogService) DeleteBlog(id int) error {
    return s.BlogRepository.DeleteBlog(id)
}