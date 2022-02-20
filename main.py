import email
from typing import Optional
from typing import Optional
from fastapi import FastAPI
from pydantic import BaseModel

from mongoengine import *

app = FastAPI()


class BookIn(BaseModel):
    title: str
    author: str
    listprice: str
    saleprice: str


class Book(Document):
    title = StringField(required=True)
    author = StringField(max_length=100)
    listprice = StringField(max_length=50)
    saleprice = StringField(max_length=50)


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
    book = get_books_from_mongo(book_title)
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


def connect_atlas():
    connect(
        host="mongodb+srv://is766:HTZIetTwHD4tkQjn@is766cluster0.dpa1z.mongodb.net/is766db?retryWrites=true&w=majority"
    )
    return


def get_books_from_mongo():
    books = []
    for rec in Book.objects:
        books.append(
            {
                "title": rec.title,
                "author": rec.author,
                "listprice": rec.listprice,
                "saleprice": rec.saleprice,
            }
        )
    return books


def get_books_from_mongo(book_title: str):
    book = Book.objects(title=book_title)[0]
    return book


def create_new_book_in_mongo(book: BookIn):
    bookDoc = Book(title=book.title)
    bookDoc.author = book.author
    bookDoc.listprice = book.listprice
    bookDoc.saleprice = book.saleprice
    bookDoc.save()
    return


def update_book_in_mongo(book_title: str, update_detail: BookIn):
    book = Book.objects(title=book_title)
    print(
        {
            "title": book[0].title,
            "author": book[0].author,
            "listprice": book[0].listprice,
            "saleprice": book[0].saleprice,
        }
    )
    book[0].delete()

    bookDoc = Book(title=update_detail.title)
    bookDoc.author = update_detail.author
    bookDoc.listprice = update_detail.listprice
    bookDoc.saleprice = update_detail.saleprice
    bookDoc.save()
    return


def delete_book_in_mongo(book_title: str):
    book = Book.objects(title=book_title)
    book[0].delete()
    return
