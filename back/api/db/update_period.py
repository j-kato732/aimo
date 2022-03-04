import sqlite3
from tkinter import W

dbname = 'test.db'
conn = sqlite3.connect(dbname)

cur = conn.cursor()

conn.close()