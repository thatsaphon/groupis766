import email
from typing import Optional
from typing import Optional
from fastapi import FastAPI

from db import *


app = FastAPI()


@app.get("/")
def get_all_books():
    connect_atlas()
    books = get_books_from_mongo()
    return {
        "books": books
    }


@app.get("/{book_title}")
def get_book_from_title(book_title: str):
    connect_atlas()
    book = get_book_from_mongo(book_title)
    return book


@app.post("/", status_code=201)
def create_new_book(book: BookIn):
    connect_atlas()
    create_new_book_in_mongo(book)
    return book


@app.put("/{book_title}")
def update_book(book_title: str, book: BookIn):
    connect_atlas()
    update_book_in_mongo(book_title, book)
    return


@app.delete("/{book_title}")
def delete_book(book_title: str):
    connect_atlas()
    delete_book_in_mongo(book_title)
    return
