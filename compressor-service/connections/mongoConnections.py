import os
from typing import Optional
from pymongo import MongoClient
import base64


class MongoConnection:
    _instance: Optional["MongoConnection"] = None
    client: MongoClient

    def __new__(cls) -> "MongoConnection":
        if not cls._instance:
            username_b64 = os.environ.get("MONGO_USERNAME")
            password_b64 = os.environ.get("MONGO_PASS")
            username = base64.b64decode(str(username_b64)).decode("utf-8")
            password = base64.b64decode(str(password_b64)).decode("utf-8")
            cls._instance = super(MongoConnection, cls).__new__(cls)
            cls._instance.client = MongoClient(
                f"mongodb://{username}:{password}@mongodb.default.svc.cluster.local:27017/myFiles?authSource=admin&authMechanism=SCRAM-SHA-1")
        return cls._instance

    def get_client(self):
        return self.client
