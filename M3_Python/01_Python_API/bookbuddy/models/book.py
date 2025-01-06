from dataclasses import dataclass
from typing import Optional

@dataclass
class Book:
    title: str
    author: str
    published_year: int
    genre: str
    id: Optional[int] = None