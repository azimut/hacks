#!/usr/bin/env mitmdump -s

from mitmproxy import ctx
from mitmproxy import http
from mitmproxy import addonmanager
from mitmproxy.script import concurrent

import pprint
import psycopg2

TABLE_DEFINITION = """
DROP TABLE http_entries;
CREATE TABLE IF NOT EXISTS http_entries(
    timestamp TIMESTAMP DEFAULT NOW(),
    method    VARCHAR(8),
    scheme    VARCHAR(8),
    host      VARCHAR(256),
    port      SMALLINT,
    path      VARCHAR(256),
    version   VARCHAR(16),
    status    SMALLINT
);
"""

pp = pprint.PrettyPrinter(indent=4)

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
        # 'request:http_version': flow.request.http_version,
        'request:host': flow.request.host,
        'request:port': flow.request.port,
        'request:path': flow.request.path,
        # 'request:headers': flow.request.headers,
        # 'request:content': flow.request.get_content(),
        # 'request:timestamp_start': flow.request.timestamp_start,
        # 'request:timestamp_end': flow.request.timestamp_end,
        'response:http_version': flow.response.http_version,
        'response:status_code': flow.response.status_code,
        'response:headers': flow.response.headers,
        # 'response:content': flow.response.get_content(),
        # 'response:timestamp_start': flow.response.timestamp_start,
        # 'response:timestamp_end': flow.response.timestamp_end,
    }
    with conn:
        with conn.cursor() as cur:
            cur.execute("""INSERT INTO http_entries(port, scheme, host, version, status, path, method)
                           VALUES (%s,%s,%s,%s,%s,%s,%s)""",
                        (flow.request.port,
                         flow.request.scheme,
                         flow.request.host,
                         flow.response.http_version,
                         flow.response.status_code,
                         flow.request.path,
                         flow.request.method,))
    pp.pprint(data)
