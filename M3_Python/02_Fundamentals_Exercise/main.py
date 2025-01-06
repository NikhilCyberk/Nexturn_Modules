
from book_management import BookManager  
from customer_management import CustomerManager  
from sales_management import SalesManager

def main():
    """Main program entry point with menu-driven interface."""
    book_manager = BookManager()
    customer_manager = CustomerManager()
    sales_manager = SalesManager(book_manager, customer_manager)

    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")

        try:
            choice = int(input("Enter your choice: "))
            
            if choice == 1:
                handle_book_management(book_manager)
            elif choice == 2:
                handle_customer_management(customer_manager)
            elif choice == 3:
                handle_sales_management(sales_manager)
            elif choice == 4:
                print("Thank you for using BookMart!")
                break
            else:
                print("Invalid choice. Please try again.")
        
        except ValueError as e:
            print(f"Error: {e}")
        except Exception as e:
            print(f"An unexpected error occurred: {e}")

def handle_book_management(book_manager: BookManager):
    """Handles book management menu options."""
    print("\nBook Management")
    print("1. Add Book")
    print("2. View All Books")
    print("3. Search Books")
    
    try:
        choice = int(input("Enter your choice: "))
        
        if choice == 1:
            title = input("Enter book title: ")
            author = input("Enter author name: ")
            price = float(input("Enter price: "))
            quantity = int(input("Enter quantity: "))
            book_manager.add_book(title, author, price, quantity)
            print("Book added successfully!")
            
        elif choice == 2:
            if not book_manager.books:
                print("No books in inventory.")
            else:
                for book in book_manager.books.values():
                    print("\n" + book.display_details())
                    print("-" * 30)
                    
        elif choice == 3:
            query = input("Enter search term (title or author): ")
            results = book_manager.search_book(query)
            if results:
                for book in results:
                    print("\n" + book.display_details())
                    print("-" * 30)
            else:
                print("No matching books found.")
                
    except ValueError as e:
        print(f"Error: {e}")

def handle_customer_management(customer_manager: CustomerManager):
    """Handles customer management menu options."""
    print("\nCustomer Management")
    print("1. Add Customer")
    print("2. View All Customers")
    
    try:
        choice = int(input("Enter your choice: "))
        
        if choice == 1:
            name = input("Enter customer name: ")
            email = input("Enter email: ")
            phone = input("Enter phone number: ")
            customer_manager.add_customer(name, email, phone)
            print("Customer added successfully!")
            
        elif choice == 2:
            if not customer_manager.customers:
                print("No customers registered.")
            else:
                for customer in customer_manager.customers:
                    print("\n" + customer.display_details())
                    print("-" * 30)
                    
    except ValueError as e:
        print(f"Error: {e}")

def handle_sales_management(sales_manager: SalesManager):
    """Handles sales management menu options."""
    print("\nSales Management")
    print("1. Create Sale")
    print("2. View All Sales")
    
    try:
        choice = int(input("Enter your choice: "))
        
        if choice == 1:
            email = input("Enter customer email: ")
            book_title = input("Enter book title: ")
            quantity = int(input("Enter quantity: "))
            transaction = sales_manager.create_sale(email, book_title, quantity)
            print("\nSale completed successfully!")
            print(transaction.display_transaction())
            
        elif choice == 2:
            if not sales_manager.transactions:
                print("No sales records found.")
            else:
                for transaction in sales_manager.transactions:
                    print("\n" + transaction.display_transaction())
                    print("-" * 30)
                    
    except ValueError as e:
        print(f"Error: {e}")

if __name__ == "__main__":
    main()