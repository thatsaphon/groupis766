import email
from typing import Optional
from fastapi import FastAPI
from pydantic import BaseModel

from mongoengine import *

app = FastAPI()


class BookIn(BaseModel):
    title = str
    author = str
    listprice = str
    saleprice = str


class Book(Document):
    title = StringField(required=True)
    author = StringField(max_length=100)
    listprice = StringField(max_length=50)
    saleprice = StringField(max_length=50)


@app.get("/")
def get_books():
    connect_atlas()
    books = get_books_from_mongo()
    return {
        "books": books
    }


@app.post("/", status_code=201)
def create_new_book(book: BookIn):
    print(book.title)
    connect_atlas()
    create_new_book_in_mongo(book)
    return {
        "message": "{} was created.".format(book.title)
    }


def connect_atlas():
    connect(
        host="mongodb+srv://is766:HTZIetTwHD4tkQjn@is766cluster0.dpa1z.mongodb.net/is766db?retryWrites=true&w=majority"
    )
    return


def get_books_from_mongo():
    books = []
    for rec in Book.objects:
        print()
        books.append(
            {
                "title": rec.title,
                "author": rec.author,
                "listprice": rec.listprice,
                "saleprice": rec.saleprice,
            }
        )
    return books


def create_new_book_in_mongo(book: BookIn):
    bookDoc = Book(title=book.title)
    bookDoc.author = book.author
    bookDoc.listprice = book.listprice
    bookDoc.saleprice = book.saleprice
    bookDoc.save()
    return
