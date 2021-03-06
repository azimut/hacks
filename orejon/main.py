#!/usr/bin/env mitmdump --ssl-insecure -s

from mitmproxy import ctx
from mitmproxy import http
from mitmproxy import addonmanager
from mitmproxy.script import concurrent

import json
import psycopg2
import os

from psycopg2.extras import Json

TABLE_DEFINITION = """
CREATE TABLE IF NOT EXISTS http_entries(
    timestamp TIMESTAMP DEFAULT NOW(),
    method    VARCHAR(8),
    scheme    VARCHAR(8),
    host      VARCHAR(256),
    port      INTEGER,
    path      VARCHAR(256),
    version   VARCHAR(16),
    status    SMALLINT,
    length    INTEGER,
    qheaders  JSONB,
    rheaders  JSONB
);
"""

dbname = os.getenv('DB',default='postgres')

class MyJson(Json):
    def dumps(self, obj):
        d = {}
        for x,y in obj.fields:
            val=y.decode()
            key=x.decode()
            d[key]=val
        return json.dumps(d)

try:
    conn = psycopg2.connect(dbname=dbname, user="postgres")
except:
    ctx.log.error("unable to connect")

def load(entry: addonmanager.Loader):
     cur = conn.cursor()
     cur.execute(TABLE_DEFINITION)
     conn.commit()

def done():
     conn.commit()
     conn.close()

@concurrent
def response(flow: http.HTTPFlow) -> None:
    with conn:
        with conn.cursor() as cur:
            cur.execute("""INSERT INTO http_entries(port, scheme, host, version, status, path, method, length, qheaders, rheaders)
                           VALUES (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)""",
                        (flow.request.port,
                         flow.request.scheme,
                         flow.request.host,
                         flow.response.http_version,
                         flow.response.status_code,
                         flow.request.path,
                         flow.request.method,
                         len(flow.response.get_content()),
                         MyJson(flow.request.headers),
                         MyJson(flow.response.headers)))
