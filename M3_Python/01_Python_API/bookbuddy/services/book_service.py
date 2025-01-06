import sqlite3
from typing import List, Optional, Tuple, Dict, Any
from models.book import Book
from config import Config

class BookService:
    def __init__(self, db_path: str = Config.SQLITE_DB_PATH):
        self.db_path = db_path
        self._initialize_db()

    def _get_connection(self):
        """Create and return a new database connection."""
        conn = sqlite3.connect(self.db_path)
        conn.row_factory = sqlite3.Row
        return conn

    def _initialize_db(self) -> None:
        """Initialize the database and create the books table if it doesn't exist."""
        conn = self._get_connection()
        try:
            conn.execute('''
                CREATE TABLE IF NOT EXISTS books (
                    id INTEGER PRIMARY KEY AUTOINCREMENT,
                    title TEXT NOT NULL,
                    author TEXT NOT NULL,
                    published_year INTEGER NOT NULL,
                    genre TEXT NOT NULL
                )
            ''')
            conn.commit()
        finally:
            conn.close()
    def _execute_query(self, query: str, params: tuple = ()) -> Tuple[bool, Any, Optional[str]]:
        """Execute a database query with proper connection handling."""
        conn = self._get_connection()
        try:
            cursor = conn.cursor()
            cursor.execute(query, params)
            conn.commit()
            return True, cursor, None
        except sqlite3.Error as e:
            return False, None, str(e)
        finally:
            conn.close()

    def add_book(self, book: Book) -> Tuple[bool, int, Optional[str]]:
        """Add a new book to the database."""
        try:
            with sqlite3.connect(self.db_path) as conn:
                cursor = conn.cursor()
                cursor.execute(
                    'INSERT INTO books (title, author, published_year, genre) VALUES (?, ?, ?, ?)',
                    (book.title, book.author, book.published_year, book.genre)
                )
                return True, cursor.lastrowid, None
        except sqlite3.Error as e:
            return False, -1, str(e)

    def get_books(self, filters: Optional[Dict[str, str]] = None) -> Tuple[bool, List[Book], Optional[str]]:
        """Get all books, optionally filtered by genre or author."""
        try:
            with sqlite3.connect(self.db_path) as conn:
                conn.row_factory = sqlite3.Row
                cursor = conn.cursor()
                
                query = 'SELECT * FROM books'
                params = []
                
                if filters:
                    conditions = []
                    if 'genre' in filters:
                        conditions.append('genre = ?')
                        params.append(filters['genre'])
                    if 'author' in filters:
                        conditions.append('author = ?')
                        params.append(filters['author'])
                    
                    if conditions:
                        query += ' WHERE ' + ' AND '.join(conditions)
                
                cursor.execute(query, params)
                rows = cursor.fetchall()
                
                books = [Book(
                    id=row['id'],
                    title=row['title'],
                    author=row['author'],
                    published_year=row['published_year'],
                    genre=row['genre']
                ) for row in rows]
                
                return True, books, None
        except sqlite3.Error as e:
            return False, [], str(e)

    def get_book_by_id(self, book_id: int) -> Tuple[bool, Optional[Book], Optional[str]]:
        """Get a book by its ID."""
        try:
            with sqlite3.connect(self.db_path) as conn:
                conn.row_factory = sqlite3.Row
                cursor = conn.cursor()
                cursor.execute('SELECT * FROM books WHERE id = ?', (book_id,))
                row = cursor.fetchone()
                
                if not row:
                    return False, None, "No book exists with the provided ID"
                
                book = Book(
                    id=row['id'],
                    title=row['title'],
                    author=row['author'],
                    published_year=row['published_year'],
                    genre=row['genre']
                )
                return True, book, None
        except sqlite3.Error as e:
            return False, None, str(e)

    def update_book(self, book_id: int, data: Dict[str, Any]) -> Tuple[bool, Optional[str]]:
        """Update an existing book."""
        try:
            with sqlite3.connect(self.db_path) as conn:
                cursor = conn.cursor()
                
                # Check if book exists
                cursor.execute('SELECT 1 FROM books WHERE id = ?', (book_id,))
                if not cursor.fetchone():
                    return False, "No book exists with the provided ID"
                
                # Build update query dynamically based on provided fields
                update_fields = []
                params = []
                for key, value in data.items():
                    if key in {'title', 'author', 'published_year', 'genre'}:
                        update_fields.append(f'{key} = ?')
                        params.append(value)
                
                if not update_fields:
                    return False, "No valid fields to update"
                
                params.append(book_id)
                query = f'UPDATE books SET {", ".join(update_fields)} WHERE id = ?'
                cursor.execute(query, params)
                
                return True, None
        except sqlite3.Error as e:
            return False, str(e)

    def delete_book(self, book_id: int) -> Tuple[bool, Optional[str]]:
        """Delete a book by its ID."""
        try:
            with sqlite3.connect(self.db_path) as conn:
                cursor = conn.cursor()
                cursor.execute('DELETE FROM books WHERE id = ?', (book_id,))
                
                if cursor.rowcount == 0:
                    return False, "No book exists with the provided ID"
                
                return True, None
        except sqlite3.Error as e:
            return False, str(e)