import logging
from utils.compress import compress_pdf
from connections.kafkaConnections import KafkaConnections

topics = ['my-topic']


def main():
    connection = KafkaConnections()
    consumer = connection.get_consumer()
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
                logging.info("Consumed event from topic {topic}: key = {key} value = {value}".format(topic=msg.topic(), key=msg.key().decode(
                    'utf-8') if msg.key() is not None else None, value=msg.value().decode('utf-8') if msg.value() is not None else None))
                compress_pdf(msg.key().decode('utf-8') ,msg.value().decode('utf-8'))
    except KeyboardInterrupt:
        logging.basicConfig(level=logging.DEBUG)
        logging.debug("In except")
    finally:
        logging.basicConfig(level=logging.DEBUG)
        logging.debug("In finally")
        consumer.close()


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    logging.info("Running The Consumer Service")
    main()
