from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()


class LogIn(BaseModel):
    username: str
    password: str


class Register(BaseModel):
    username: str
    password: str
    email: str


class Job(BaseModel):
    title: str
    salary: str
    company: str
    url: str


@app.post("/login/")
def login(login: LogIn):
    return {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
    }


@app.post("/register/")
def register(register: Register):
    return {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
    }


@app.get("/search")
def searchJob(keyword: str):
    return {
        "jobs": [
            {
                "title": "Accountant",
                "salary": "25000-30000",
                "company": "PTT",
                "url": "https://www.jobsdb.com/test/ptt"
            },
            {
                "title": "Accountant",
                "salary": "30000-40000",
                "company": "SCG",
                "url": "https://www.jobsdb.com/test/scg"
            }
        ]
    }


@app.post("/like")
def addToFavourite(job: Job):
    return {
        "message": "job is added to favourite."
    }


@app.post("/apply")
def markAsApplied(job: Job):
    return {
        "message": "job is added applied."
    }


@app.get("/job")  # header userId
def getAllSavedJob():
    return{
        "jobs": [
            {
                "title": "Accountant",
                "salary": "25000-30000",
                "company": "PTT",
                "url": "https://www.jobsdb.com/test/ptt"
            },
            {
                "title": "Accountant",
                "salary": "30000-40000",
                "company": "SCG",
                "url": "https://www.jobsdb.com/test/scg"
            }
        ]
    }
