from datetime import datetime
from typing import Dict, Any, Tuple, Optional
from config import Config

def validate_book_data(data: Dict[Any, Any], update: bool = False) -> Tuple[bool, Optional[str]]:
    """Validate book data for creation or update."""
    required_fields = {'title', 'author', 'published_year', 'genre'}
    
    # For updates, we don't require all fields
    if not update:
        missing_fields = required_fields - set(data.keys())
        if missing_fields:
            return False, f"Missing required fields: {', '.join(missing_fields)}"
    
    if 'published_year' in data:
        try:
            year = int(data['published_year'])
            if not (Config.MIN_YEAR <= year <= Config.MAX_YEAR):
                return False, f"Published year must be between {Config.MIN_YEAR} and {Config.MAX_YEAR}"
        except ValueError:
            return False, "Published year must be a valid integer"
    
    if 'genre' in data and data['genre'] not in Config.VALID_GENRES:
        return False, f"Genre must be one of: {', '.join(Config.VALID_GENRES)}"
    
    return True, None