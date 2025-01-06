import unittest
import json
from app import create_app
from services.book_service import BookService
import os
from config import Config

class TestBookRoutes(unittest.TestCase):
    """Test suite for BookBuddy API routes."""

    def setUp(self):
        """Set up test environment before each test."""
        self.app = create_app()
        self.client = self.app.test_client()
        self.db_path = 'test_bookbuddy.db'
        
        # Ensure clean database for each test
        if os.path.exists(self.db_path):
            os.remove(self.db_path)
            
        self.book_service = BookService(self.db_path)
        
        # Sample valid book data for testing
        self.valid_book = {
            "title": "Test Book",
            "author": "Test Author",
            "published_year": 2020,
            "genre": "Fiction"
        }

    def tearDown(self):
        """Clean up test environment after each test."""
        # Close all database connections
        for conn in sqlite3.connect(self.db_path).execute(
            "SELECT name FROM sqlite_master WHERE type='table'"
        ).fetchall():
            sqlite3.connect(self.db_path).close()
            
        # Remove test database
        try:
            os.remove(self.db_path)
        except PermissionError:
            # If file is locked, wait briefly and try again
            import time
            time.sleep(0.1)
            try:
                os.remove(self.db_path)
            except:
                pass  # If still can't remove, let the next test handle it

    def test_create_book_success(self):
        """Test successful book creation."""
        response = self.client.post('/books',
                                  data=json.dumps(self.valid_book),
                                  content_type='application/json')
        
        self.assertEqual(response.status_code, 201)
        data = response.get_json()
        self.assertIn('book_id', data)
        self.assertIn('message', data)
        self.assertEqual(data['message'], 'Book added successfully')

    def test_create_book_invalid_data(self):
        """Test book creation with invalid data."""
        # Missing required field
        invalid_book = self.valid_book.copy()
        del invalid_book['title']
        response = self.client.post('/books',
                                  data=json.dumps(invalid_book),
                                  content_type='application/json')
        
        self.assertEqual(response.status_code, 400)
        self.assertIn('error', response.get_json())

        # Invalid genre
        invalid_book = self.valid_book.copy()
        invalid_book['genre'] = 'InvalidGenre'
        response = self.client.post('/books',
                                  data=json.dumps(invalid_book),
                                  content_type='application/json')
        
        self.assertEqual(response.status_code, 400)
        self.assertIn('error', response.get_json())

        # Invalid year
        invalid_book = self.valid_book.copy()
        invalid_book['published_year'] = 3000
        response = self.client.post('/books',
                                  data=json.dumps(invalid_book),
                                  content_type='application/json')
        
        self.assertEqual(response.status_code, 400)
        self.assertIn('error', response.get_json())

    def test_get_books_empty(self):
        """Test getting books when database is empty."""
        response = self.client.get('/books')
        
        self.assertEqual(response.status_code, 200)
        data = response.get_json()
        self.assertEqual(len(data), 0)
        self.assertIsInstance(data, list)

    def test_get_books_with_data(self):
        """Test getting books when database has entries."""
        # First create some books
        self.client.post('/books',
                        data=json.dumps(self.valid_book),
                        content_type='application/json')
        
        second_book = self.valid_book.copy()
        second_book['title'] = 'Second Test Book'
        self.client.post('/books',
                        data=json.dumps(second_book),
                        content_type='application/json')

        # Get all books
        response = self.client.get('/books')
        
        self.assertEqual(response.status_code, 200)
        data = response.get_json()
        self.assertEqual(len(data), 2)
        self.assertEqual(data[0]['title'], 'Test Book')
        self.assertEqual(data[1]['title'], 'Second Test Book')

    def test_get_books_with_filters(self):
        """Test getting books with genre and author filters."""
        # Create books with different genres and authors
        self.client.post('/books',
                        data=json.dumps(self.valid_book),
                        content_type='application/json')
        
        mystery_book = {
            "title": "Mystery Book",
            "author": "Mystery Author",
            "published_year": 2020,
            "genre": "Mystery"
        }
        self.client.post('/books',
                        data=json.dumps(mystery_book),
                        content_type='application/json')

        # Test genre filter
        response = self.client.get('/books?genre=Fiction')
        self.assertEqual(response.status_code, 200)
        data = response.get_json()
        self.assertEqual(len(data), 1)
        self.assertEqual(data[0]['genre'], 'Fiction')

        # Test author filter
        response = self.client.get('/books?author=Mystery Author')
        self.assertEqual(response.status_code, 200)
        data = response.get_json()
        self.assertEqual(len(data), 1)
        self.assertEqual(data[0]['author'], 'Mystery Author')

        # Test combined filters
        response = self.client.get('/books?genre=Fiction&author=Test Author')
        self.assertEqual(response.status_code, 200)
        data = response.get_json()
        self.assertEqual(len(data), 1)
        self.assertEqual(data[0]['genre'], 'Fiction')
        self.assertEqual(data[0]['author'], 'Test Author')

    def test_get_book_by_id(self):
        """Test getting a specific book by ID."""
        # First create a book
        response = self.client.post('/books',
                                  data=json.dumps(self.valid_book),
                                  content_type='application/json')
        book_id = response.get_json()['book_id']

        # Get the book by ID
        response = self.client.get(f'/books/{book_id}')
        
        self.assertEqual(response.status_code, 200)
        data = response.get_json()
        self.assertEqual(data['title'], self.valid_book['title'])
        self.assertEqual(data['author'], self.valid_book['author'])

        # Test nonexistent book ID
        response = self.client.get('/books/999')
        self.assertEqual(response.status_code, 404)
        self.assertIn('error', response.get_json())

    def test_update_book(self):
        """Test updating a book's information."""
        # First create a book
        response = self.client.post('/books',
                                  data=json.dumps(self.valid_book),
                                  content_type='application/json')
        book_id = response.get_json()['book_id']

        # Update the book
        updated_data = {
            "title": "Updated Title",
            "published_year": 2021
        }
        response = self.client.put(f'/books/{book_id}',
                                 data=json.dumps(updated_data),
                                 content_type='application/json')
        
        self.assertEqual(response.status_code, 200)
        self.assertIn('message', response.get_json())

        # Verify the update
        response = self.client.get(f'/books/{book_id}')
        data = response.get_json()
        self.assertEqual(data['title'], "Updated Title")
        self.assertEqual(data['published_year'], 2021)
        # Original data should remain unchanged
        self.assertEqual(data['author'], self.valid_book['author'])

        # Test updating nonexistent book
        response = self.client.put('/books/999',
                                 data=json.dumps(updated_data),
                                 content_type='application/json')
        self.assertEqual(response.status_code, 404)

        # Test updating with invalid data
        invalid_update = {"published_year": 3000}
        response = self.client.put(f'/books/{book_id}',
                                 data=json.dumps(invalid_update),
                                 content_type='application/json')
        self.assertEqual(response.status_code, 400)

    def test_delete_book(self):
        """Test deleting a book."""
        # First create a book
        response = self.client.post('/books',
                                  data=json.dumps(self.valid_book),
                                  content_type='application/json')
        book_id = response.get_json()['book_id']

        # Delete the book
        response = self.client.delete(f'/books/{book_id}')
        self.assertEqual(response.status_code, 200)
        self.assertIn('message', response.get_json())

        # Verify the book is deleted
        response = self.client.get(f'/books/{book_id}')
        self.assertEqual(response.status_code, 404)

        # Test deleting nonexistent book
        response = self.client.delete('/books/999')
        self.assertEqual(response.status_code, 404)

    def test_concurrent_operations(self):
        """Test multiple operations in sequence to verify data consistency."""
        # Create multiple books
        books = []
        for i in range(3):
            book = self.valid_book.copy()
            book['title'] = f'Book {i}'
            response = self.client.post('/books',
                                      data=json.dumps(book),
                                      content_type='application/json')
            books.append(response.get_json()['book_id'])

        # Update one book
        update_data = {"title": "Updated Book"}
        self.client.put(f'/books/{books[0]}',
                       data=json.dumps(update_data),
                       content_type='application/json')

        # Delete another book
        self.client.delete(f'/books/{books[1]}')

        # Verify final state
        response = self.client.get('/books')
        data = response.get_json()
        self.assertEqual(len(data), 2)  # One deleted
        titles = [book['title'] for book in data]
        self.assertIn('Updated Book', titles)
        self.assertIn('Book 2', titles)

    def test_edge_cases(self):
        """Test various edge cases and boundary conditions."""
        # Test with minimum valid year
        min_year_book = self.valid_book.copy()
        min_year_book['published_year'] = Config.MIN_YEAR
        response = self.client.post('/books',
                                  data=json.dumps(min_year_book),
                                  content_type='application/json')
        self.assertEqual(response.status_code, 201)

        # Test with maximum valid year
        max_year_book = self.valid_book.copy()
        max_year_book['published_year'] = Config.MAX_YEAR
        response = self.client.post('/books',
                                  data=json.dumps(max_year_book),
                                  content_type='application/json')
        self.assertEqual(response.status_code, 201)

        # Test with empty strings
        empty_book = self.valid_book.copy()
        empty_book['title'] = ""
        response = self.client.post('/books',
                                  data=json.dumps(empty_book),
                                  content_type='application/json')
        self.assertEqual(response.status_code, 400)

        # Test with very long strings
        long_book = self.valid_book.copy()
        long_book['title'] = "a" * 1000
        response = self.client.post('/books',
                                  data=json.dumps(long_book),
                                  content_type='application/json')
        self.assertEqual(response.status_code, 201)

if __name__ == '__main__':
    unittest.main()