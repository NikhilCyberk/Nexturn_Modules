from flask import Flask
from routes.book_routes import book_routes

def create_app():
    app = Flask(__name__)
    app.register_blueprint(book_routes)
    return app

if __name__ == '__main__':
    app = create_app()
    app.run(debug=True)