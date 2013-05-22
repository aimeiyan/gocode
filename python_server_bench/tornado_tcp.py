__author__ = 'feng'

from tornado.tcpserver import TCPServer
from tornado.ioloop import IOLoop

from socket_test import RESPONSE

class TestTCPServer(TCPServer):
    def __init__(self, io_loop, **kwargs):
        TCPServer.__init__(self, io_loop=io_loop, ssl_options=None, **kwargs)

    # def _on_header

    def handle_stream(self, stream, address):
        def on_header(data):
            # print len(data)
            stream.write_to_fd(RESPONSE)
            stream.read_until(b'\r\n\r\n', on_header)

            # pass

        # print "handle_stream", address
        # data = stream.read_from_fd()
        stream.read_until(b'\r\n\r\n', on_header)
        # while True:

        # print stream


"""
pypy
unning 10s test @ http://127.0.0.1:9090/
  1 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   117.45us   18.32us 172.00us   66.67%
    Req/Sec    73.47k     0.96k   74.00k    83.33%
  746136 requests in 10.00s, 51.23MB read
Requests/sec:  74613.35
Transfer/sec:      5.12MB

python:
wrk -d10 -c10 -t1 'http://127.0.0.1:9090/'
Running 10s test @ http://127.0.0.1:9090/
  1 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   483.22us   15.30us 543.00us   79.10%
    Req/Sec    19.99k   122.17    20.00k    98.51%
  203276 requests in 10.00s, 13.96MB read
Requests/sec:  20327.67
Transfer/sec:      1.40MB

"""

if __name__ == "__main__":
    server = TestTCPServer(IOLoop.instance())
    server.listen(9090)
    IOLoop.instance().start()

# server.bind(9090)
# server.start()