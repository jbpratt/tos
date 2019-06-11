#!/bin/bash
sqlite3 mookies.db .dump | gzip -c > mookies.db.dump.gz
