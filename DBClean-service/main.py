from pymongo import MongoClient
from datetime import datetime, timedelta
import os
import base64


def delete_old_files():
    username_b64 = os.environ.get("MONGO_USERNAME")
    password_b64 = os.environ.get("MONGO_PASS")
    username = base64.b64decode(str(username_b64)).decode("utf-8")
    password = base64.b64decode(str(password_b64)).decode("utf-8")
    client = MongoClient(
        f"mongodb://{username}:{password}@mongodb.default.svc.cluster.local:27017/myFiles?authSource=admin&authMechanism=SCRAM-SHA-1")
    db = client["myFiles"]
    fs = db["fs.files"]
    threshold_time = datetime.utcnow() - timedelta(minutes=2)
    fs.delete_many({"uploadDate": {"$lt": threshold_time}})
    client.close()


delete_old_files()
