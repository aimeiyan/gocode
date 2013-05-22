__author__ = 'feng'


import socket
import threading


RESPONSE = "HTTP/1.1 200 OK\r\nConnection: Keep-Alive\r\nContent-Length:10\r\n\r\n1234567890"


s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind(("127.0.0.1", 9090))
s.listen(1024)

"""
python:
Running 10s test @ http://127.0.0.1:9090/
  1 threads and 1 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    21.00us    2.43us  38.00us   96.97%
    Req/Sec    35.95k   273.54    36.00k    96.97%
  368413 requests in 10.00s, 16.86MB read
Requests/sec:  36841.36
Transfer/sec:      1.69MB
wrk -d10 -c1 -t1 'http://127.0.0.1:9090/'  0.81s user 3.29s system 41% cpu 10.004 total

pypy:
1 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.60us   22.36us 181.00us   98.53%
    Req/Sec    42.82k   486.88    43.00k    85.29%
  431064 requests in 10.00s, 19.73MB read
  Socket errors: connect 0, read 0, write 0, timeout 36
Requests/sec:  43106.50
Transfer/sec:      1.97MB
"""
def single_thread():
    while True:
        conn, addr = s.accept()
        print 'Connected by', addr
        while True:
            data = conn.recv(1024)
            if not data:
                conn.close()
                print "close", addr
                break
            conn.sendall(RESPONSE)

"""
Running 10s test @ http://127.0.0.1:9090/
  1 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   155.57us   91.07us 404.00us   66.18%
    Req/Sec    53.99k   440.24    55.00k    92.65%
  546533 requests in 10.00s, 37.53MB read
Requests/sec:  54653.44
Transfer/sec:      3.75MB
wrk -d10 -c10 -t1 'http://127.0.0.1:9090/'  1.54s user 6.10s system 76% cpu 10.004 total

pypy:
wrk -d10 -c10 -t1 'http://127.0.0.1:9090/'
Running 10s test @ http://127.0.0.1:9090/
  1 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   127.56us   13.48us 160.00us   82.35%
    Req/Sec    71.49k     1.79k   73.00k    97.06%
  731014 requests in 10.00s, 50.19MB read
Requests/sec:  73101.61
Transfer/sec:      5.02MB
wrk -d10 -c10 -t1 'http://127.0.0.1:9090/'  2.05s user 7.76s system 98% cpu 10.005 total
"""

def multi_thread():
    while True:
        conn, addr = s.accept()
        print 'Connected by', addr

        def worker(conn):
            while True:
                data = conn.recv(1024)
                if not data:
                    conn.close()
                    print "close", addr
                    break
                conn.sendall(RESPONSE)

        th = threading.Thread(target = worker, args = (conn,))
        th.start()


if __name__ == "__main__":
    multi_thread()
    # single_thread()






