
import re
from dataclasses import dataclass
from typing import List, Optional

@dataclass
class Customer:
    """Represents a customer with basic information."""
    name: str
    email: str
    phone: str

    def __post_init__(self):
        """Validates customer data after initialization."""
        if not self.name or not self.email or not self.phone:
            raise ValueError("All customer fields are required")
        if not re.match(r"[^@]+@[^@]+\.[^@]+", self.email):
            raise ValueError("Invalid email format")
        if not re.match(r"^\+?[\d\s-]{10,}$", self.phone):
            raise ValueError("Invalid phone number format")

    def display_details(self) -> str:
        """Returns a formatted string of customer details."""
        return f"Name: {self.name}\nEmail: {self.email}\nPhone: {self.phone}"

class CustomerManager:
    """Manages customer records and related operations."""
    def __init__(self):
        self.customers: List[Customer] = []

    def add_customer(self, name: str, email: str, phone: str) -> None:
        """Adds a new customer to the system."""
        customer = Customer(name, email, phone)
        self.customers.append(customer)

    def find_customer(self, email: str) -> Optional[Customer]:
        """Finds a customer by their email."""
        return next((c for c in self.customers if c.email == email), None)