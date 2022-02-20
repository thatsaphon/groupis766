from pydantic import BaseModel
from mongoengine import *


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


def get_book_from_mongo(book_title: str):
    book = Book.objects(title=book_title)[0]
    return {
        "title": book.title,
        "author": book.author,
        "listprice": book.listprice,
        "saleprice": book.saleprice,
    }


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
