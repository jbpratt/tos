#!/bin/bash
zcat mookies.db.dump.gz | sqlite3 mookies.db
