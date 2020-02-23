__author__ = 'Daniil Popov'
# -*- coding: utf-8 -*-

import time
import json
import asyncpg
from configparser import ConfigParser

config = ConfigParser()
config.read('env.ini')

conn = None

async def init_conn():
    conn = await asyncpg.connect(dsn=config.get('database', 'url'))
    await conn.execute("""
    CREATE EXTENSION IF NOT EXISTS "pgcrypto";

    CREATE TABLE IF NOT EXISTS "public"."Event" (
        "id" uuid NOT NULL DEFAULT gen_random_uuid(),
        "event_data" json NOT NULL,
        "created_datetime" timestamp NOT NULL DEFAULT NOW(),
        PRIMARY KEY ("id")
    );
    """)


def close_conn():
    conn.close()


async def create(event: dict):
    conn = await asyncpg.connect(dsn=config.get('database', 'url'))
    await conn.execute("""
        INSERT INTO "Event" (event_data) VALUES('{0}')
    """.format(json.dumps(event)))
