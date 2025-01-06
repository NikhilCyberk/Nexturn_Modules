from flask import Blueprint, request, jsonify
from models.book import Book
from services.book_service import BookService
from utils.validators import validate_book_data

book_routes = Blueprint('book_routes', __name__)
book_service = BookService()

@book_routes.route('/books', methods=['POST'])
def create_book():
    """Create a new book."""
    data = request.get_json()
    
    # Validate input data
    is_valid, error = validate_book_data(data)
    if not is_valid:
        return jsonify({'error': 'Invalid data', 'message': error}), 400
    
    # Create book object and add to database
    book = Book(**data)
    success, book_id, error = book_service.add_book(book)
    
    if not success:
        return jsonify({'error': 'Database error', 'message': error}), 500
    
    return jsonify({'message': 'Book added successfully', 'book_id': book_id}), 201

@book_routes.route('/books', methods=['GET'])
def get_books():
    """Get all books with optional filtering."""
    filters = {}
    if 'genre' in request.args:
        filters['genre'] = request.args['genre']
    if 'author' in request.args:
        filters['author'] = request.args['author']
    
    success, books, error = book_service.get_books(filters)
    
    if not success:
        return jsonify({'error': 'Database error', 'message': error}), 500
    
    return jsonify([vars(book) for book in books]), 200

@book_routes.route('/books/<int:book_id>', methods=['GET'])
def get_book(book_id):
    """Get a specific book by ID."""
    success, book, error = book_service.get_book_by_id(book_id)
    
    if not success:
        return jsonify({'error': 'Book not found', 'message': error}), 404
    
    return jsonify(vars(book)), 200

@book_routes.route('/books/<int:book_id>', methods=['PUT'])
def update_book(book_id):
    """Update an existing book."""
    data = request.get_json()
    
    # Validate input data
    is_valid, error = validate_book_data(data, update=True)
    if not is_valid:
        return jsonify({'error': 'Invalid data', 'message': error}), 400
    
    success, error = book_service.update_book(book_id, data)
    
    if not success:
        if error == "Book not found":
            return jsonify({'error': 'Book not found', 'message': error}), 404
        return jsonify({'error': 'Database error', 'message': error}), 500
    
    return jsonify({'message': 'Book updated successfully'}), 200

@book_routes.route('/books/<int:book_id>', methods=['DELETE'])
def delete_book(book_id):
    """Delete a book."""
    success, error = book_service.delete_book(book_id)
    
    if not success:
        if error == "Book not found":
            return jsonify({'error': 'Book not found', 'message': error}), 404
        return jsonify({'error': 'Database error', 'message': error}), 500
    
    return jsonify({'message': 'Book deleted successfully'}), 200
