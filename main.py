from fastapi import FastAPI
from pydantic import BaseModel
from db import *


class BookIn(BaseModel):
    title: str
    author: str
    listprice: str
    saleprice: str


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
    bookDoc = Book(title=book.title)
    bookDoc.author = book.author
    bookDoc.listprice = book.listprice
    bookDoc.saleprice = book.saleprice
    bookDoc.save()
    create_new_book_in_mongo(bookDoc)
    return book


@app.put("/book/{book_title}")
def update_book(book_title: str, book: BookIn):
    bookDoc = Book(title=book.title)
    bookDoc.author = book.author
    bookDoc.listprice = book.listprice
    bookDoc.saleprice = book.saleprice
    bookDoc.save()
    update_book_in_mongo(book_title, bookDoc)
    return


@app.delete("/book/{book_title}")
def delete_book(book_title: str):
    delete_book_in_mongo(book_title)
    return
