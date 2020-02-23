__author__ = 'Daniil Popov'
# -*- coding: utf-8 -*-

from kafka import KafkaConsumer
from json import loads
from configparser import ConfigParser

config = ConfigParser()
config.read('env.ini')

from worker.consumer.trackingConsumer import consumer
from worker.manager.trackingManager import conn, init_conn
from worker.manager.trackingManager import create


async def start_consuming():
    if not conn:
        await init_conn()

    for message in consumer:
        message = message.value

        if message is None:
            continue

        await create(message)