__author__ = 'Daniil Popov'
# -*- coding: utf-8 -*-

import asyncio
from worker import start_consuming

if __name__ == "__main__":
    asyncio.get_event_loop().run_until_complete(start_consuming())
