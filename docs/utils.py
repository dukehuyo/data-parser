# utils.py
import os
import json
from datetime import datetime

def load_config(file_path: str) -> dict:
    """Load configuration from a JSON file."""
    with open(file_path, 'r') as f:
        return json.load(f)

def get_timestamp() -> str:
    """Return current timestamp in ISO 8601 format."""
    return datetime.now().isoformat()

def ensure_dir(file_path: str):
    """Create directory if it does not exist."""
    if not os.path.exists(file_path):
        os.makedirs(file_path)

def get_file_extension(file_name: str) -> str:
    """Return file extension from file name."""
    return os.path.splitext(file_name)[1]