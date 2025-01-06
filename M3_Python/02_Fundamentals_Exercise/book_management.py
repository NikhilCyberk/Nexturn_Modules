from typing import List, Optional, Dict
from dataclasses import dataclass
import re

@dataclass
class Book:
    """
    Represents a book in the inventory with essential attributes and methods.
    """
    title: str
    author: str
    price: float
    quantity: int

    def __post_init__(self):
        """Validates the book data after initialization."""
        if self.price <= 0:
            raise ValueError("Price must be a positive number")
        if self.quantity < 0:
            raise ValueError("Quantity cannot be negative")
        if not self.title or not self.author:
            raise ValueError("Title and author cannot be empty")

    def display_details(self) -> str:
        """Returns a formatted string of book details."""
        return f"Title: {self.title}\nAuthor: {self.author}\nPrice: ${self.price:.2f}\nQuantity: {self.quantity}"

    def update_quantity(self, sold_quantity: int) -> None:
        """Updates book quantity after a sale."""
        if sold_quantity > self.quantity:
            raise ValueError(f"Error: Only {self.quantity} copies available")
        self.quantity -= sold_quantity

class BookManager:
    """Manages the book inventory and related operations."""
    def __init__(self):
        self.books: Dict[str, Book] = {}

    def add_book(self, title: str, author: str, price: float, quantity: int) -> None:
        """Adds a new book to the inventory."""
        title = title.strip()
        if title in self.books:
            raise ValueError("Book already exists in inventory")
        self.books[title] = Book(title, author, price, quantity)

    def search_book(self, query: str) -> List[Book]:
        """Searches for books by title or author."""
        query = query.lower().strip()
        return [book for book in self.books.values() 
                if query in book.title.lower() or query in book.author.lower()]

    def get_book(self, title: str) -> Optional[Book]:
        """Retrieves a book by its exact title."""
        return self.books.get(title.strip())