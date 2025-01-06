from datetime import datetime
from dataclasses import dataclass
from typing import List
from customer_management import Customer  
from book_management import BookManager  
from customer_management import CustomerManager  

@dataclass
class Transaction(Customer):
    """
    Represents a sales transaction, inheriting customer information.
    Adds transaction-specific details.
    """
    book_title: str
    quantity_sold: int
    transaction_date: datetime = datetime.now()
    total_amount: float = 0

    def display_transaction(self) -> str:
        """Returns a formatted string of transaction details."""
        return (f"Transaction Date: {self.transaction_date:%Y-%m-%d %H:%M}\n"
                f"Customer: {self.name}\n"
                f"Book: {self.book_title}\n"
                f"Quantity: {self.quantity_sold}\n"
                f"Total Amount: ${self.total_amount:.2f}")

class SalesManager:
    """Manages sales transactions and related operations."""
    def __init__(self, book_manager: BookManager, customer_manager: CustomerManager):
        self.transactions: List[Transaction] = []
        self.book_manager = book_manager
        self.customer_manager = customer_manager

    def create_sale(self, customer_email: str, book_title: str, quantity: int) -> Transaction:
        """Creates a new sale transaction."""
        customer = self.customer_manager.find_customer(customer_email)
        if not customer:
            raise ValueError("Customer not found")

        book = self.book_manager.get_book(book_title)
        if not book:
            raise ValueError("Book not found")

        if quantity <= 0:
            raise ValueError("Quantity must be positive")

        if quantity > book.quantity:
            raise ValueError(f"Error: Only {book.quantity} copies available")

        # Update book quantity
        book.update_quantity(quantity)

        # Create transaction
        transaction = Transaction(
            name=customer.name,
            email=customer.email,
            phone=customer.phone,
            book_title=book_title,
            quantity_sold=quantity,
            total_amount=book.price * quantity
        )
        self.transactions.append(transaction)
        return transaction