__author__ = 'Daniil Popov'
# -*- coding: utf-8 -*-

from kafka import KafkaConsumer
from json import loads
from configparser import ConfigParser

config = ConfigParser()
config.read('env.ini')

consumer = KafkaConsumer(
        'tracker',
        bootstrap_servers=[config.get('kafka', 'port')],
        auto_offset_reset='latest',
        enable_auto_commit=False,
        group_id='tracker',
        value_deserializer=lambda x: loads(x.decode('utf-8')),
    )