#!/usr/bin/env mitmdump -s

from mitmproxy import ctx
from mitmproxy import http
from mitmproxy import addonmanager
from mitmproxy.script import concurrent

import psycopg2

TABLE_DEFINITION = """
DROP TABLE http_entries;
CREATE TABLE IF NOT EXISTS http_entries(
    path VARCHAR(128)
);
"""

try:
    conn = psycopg2.connect("dbname=postgres user=postgres")
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
     data = {
         'request:url': flow.request.url,
         'request:method': flow.request.method,
         'request:scheme': flow.request.scheme,
         'request:http_version': flow.request.http_version,
         'request:host': flow.request.host,
         'request:port': flow.request.port,
         'request:path': flow.request.path,
         'request:headers': flow.request.headers,
         #        'request:content': flow.request.get_content(),
         'request:timestamp_start': flow.request.timestamp_start,
         'request:timestamp_end': flow.request.timestamp_end,
         'response:http_version': flow.response.http_version,
         'response:status_code': flow.response.status_code,
         'response:headers': flow.response.headers,
         #        'response:content': flow.response.get_content(),
         'response:timestamp_start': flow.response.timestamp_start,
         'response:timestamp_end': flow.response.timestamp_end,
     }
     with conn:
         with conn.cursor() as cur:
             cur.execute("INSERT INTO http_entries VALUES (%s)",
                         (flow.request.host + flow.request.path,))
     print(data)
