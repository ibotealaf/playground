import sqlite3

db = sqlite3.connect("bikes.db")

def distance_of_user(name):
    try:
        result = db.execute("""
                SELECT SUM(distance) 
                FROM Trips, Users 
                WHERE Trips.user_id = Users.id AND Users.name = ?
            """, [name]
        ).fetchone()
        
        return result[0]
    except:
        return None
    

def speed_of_user(name):
    try:
        result = db.execute("""
                SELECT SUM(distance), SUM(duration) 
                FROM Trips, Users 
                WHERE Trips.user_id = Users.id AND Users.name = ?
            """, [name]
        ).fetchone()
        km_per_hr = (result[0]/result[1]) * 0.06

        return round(km_per_hr, 2)
    except:
        return None
    
def duration_in_each_city(day):
    try:
        result = db.execute("""
                SELECT C.name, SUM(T.duration) 
                FROM Trips T, Bikes B, Cities C 
                WHERE T.bike_id = B.id AND B.city_id = C.id AND T.day = ?
                GROUP BY C.name
                ORDER BY C.name
            """, [day]
        ).fetchall()

        return result
    except:
        return None
    
def users_in_city(city):
    try:
        result = db.execute("""
                SELECT COUNT(DISTINCT T.user_id)
                FROM Trips T, Bikes B, Cities C
                WHERE T.bike_id = B.id AND B.city_id = C.id AND C.name = ?
                """, [city]
            ).fetchone()
        return result[0]
    except:
        return None


def trips_on_each_day(city):
    try:
        result = db.execute("""
                    SELECT T.day, COUNT(T.id)
                    FROM Trips T, Bikes B, Cities C
                    WHERE T.bike_id = B.id AND B.city_id = C.id AND C.name = ?
                    GROUP BY T.day
            """, [city]
            ).fetchall()
        return result
    except:
        return None
    

def most_popular_cities(city):
    try:
        result = db.execute("""
                    SELECT S.name, COUNT(T.from_id)
                    FROM stops S, trips T, cities C
                    WHERE T.from_id = S.id AND S.city_id = C.id AND C.name = ?
                    GROUP BY T.from_id
                    ORDER BY COUNT(T.from_id) DESC, C.name DESC
                    LIMIT 1 
        """, [city]).fetchone()
        return result
    except:
        return None