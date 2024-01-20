import logging
from confluent_kafka import Consumer

bootstrap_servers = "kafka-release.default.svc.cluster.local:9092"
topics = ['my-topic']

def main():
    conf = {
        'bootstrap.servers': bootstrap_servers,
        'group.id': 'foo',
        'auto.offset.reset': 'earliest',
    }
    consumer = Consumer(conf)
    consumer.subscribe(topics)
    try:
        while True:
            msg = consumer.poll(1.0)
            if msg is None:
                logging.basicConfig(level=logging.DEBUG)
                logging.debug("In continue")
                continue
            elif msg.error():
                logging.basicConfig(level=logging.ERROR)
                logging.error("Error: %s".format(msg.error()))
            else:
                logging.basicConfig(level=logging.INFO)
                logging.info("Consumed event from topic {topic}: key = {key} value = {value}".format(topic=msg.topic(),key=msg.key().decode('utf-8') if msg.key() is not None else None,value=msg.value().decode('utf-8') if msg.value() is not None else None))
    except KeyboardInterrupt:
        logging.basicConfig(level=logging.DEBUG)
        logging.debug("In except")
        pass
    finally:
        logging.basicConfig(level=logging.DEBUG)
        logging.debug("In finally")
        consumer.close()

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    logging.info("Running The Consumer Service")
    main()
