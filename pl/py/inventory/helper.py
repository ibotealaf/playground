import sqlite3

db = sqlite3.connect("products.db")
db.isolation_level = None

def createTable():
    return db.execute(
        """
        CREATE TABLE IF NOT EXISTS Products
        (id INTEGER PRIMARY KEY, 
        name VARCHAR(255) NOT NULL, 
        price FLOAT NOT NULL)
        """)

def instructions():
    print("""
0 -> All Products
1 -> Find product 
2 -> Create new product
3 -> Exit
    """)

def allProducts():
    try:
        result = db.execute("SELECT * FROM Products").fetchall()
        return result
        
    except:
        print("An error occurred while processing request..")

def getProductPrice(name):
    try:
        price = db.execute("SELECT price FROM Products WHERE name = ?", [name.lower()]).fetchone()
        return price
    except:
        print("An error occurred while processing request..")
        return None


def createProduct(name, price):
    try:
        result = db.execute("INSERT INTO Products(name, price) VALUES(?, ?)", [name, round(price, 2)])
        return result.lastrowid
    except:
        print("An error occrred while creating product...")
        return None
