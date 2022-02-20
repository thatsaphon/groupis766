import email
from typing import Optional
from typing import Optional
from fastapi import FastAPI

from db import *


app = FastAPI()

connect_atlas()


@app.get("/book/")
def get_all_books():
    books = get_books_from_mongo()
    return {
        "books": books
    }


@app.get("/book/{book_title}")
def get_book_from_title(book_title: str):
    book = get_book_from_mongo(book_title)
    return book


@app.post("/book/", status_code=201)
def create_new_book(book: BookIn):
    create_new_book_in_mongo(book)
    return book


@app.put("/book/{book_title}")
def update_book(book_title: str, book: BookIn):
    update_book_in_mongo(book_title, book)
    return


@app.delete("/book/{book_title}")
def delete_book(book_title: str):
    delete_book_in_mongo(book_title)
    return
