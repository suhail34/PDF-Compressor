from typing import Optional
from confluent_kafka import Consumer, KafkaException


class KafkaConnections:
    _instance: Optional["KafkaConnections"] = None
    conf = {
        'bootstrap.servers': "kafka-release.default.svc.cluster.local:9092",
        'group.id': 'foo',
        'auto.offset.reset': 'earliest',
    }
    consumer: Consumer

    def __new__(cls) -> "KafkaConnections":
        if not cls._instance:
            cls._instance = super(KafkaConnections, cls).__new__(cls)
            cls._instance.consumer = Consumer(cls.conf)
        return cls._instance

    def get_consumer(self):
        return self.consumer
