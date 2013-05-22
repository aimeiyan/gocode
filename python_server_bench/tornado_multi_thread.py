import threading

import tornado.ioloop
import tornado.web
import pymysql
from concurrent.futures import ThreadPoolExecutor
import traceback


conns = threading.local()
pool = ThreadPoolExecutor(max_workers=2)

def task(handler):
    try:
        if not hasattr(conns, 'con'):
            conns.con = pymysql.connect(host='192.168.2.243', port=3306, user='q3boy', passwd='123', db='dealrank')
        cur = conns.con.cursor()
        # cur.execute("select * from users limit 10")

        cur.execute("select * from dealrank_ctr limit 10")

        a = []
        for r in cur.fetchall():
            a.append(str(r))
        cur.close()

        def write_back(h, d):
            h.write(d)
            h.finish()

        tornado.ioloop.IOLoop.instance().add_callback(write_back, handler, "".join(a))

        # handler.finish()
    except Exception as e:
        print e
        traceback.print_stack()

# conn = pymysql.connect(host='192.168.2.243', port=3306, user='q3boy', passwd='123', db='dealrank')

# conn = pymysql.connect(host='192.168.1.101', port=3306, user="feng", passwd='', db='rssminer')
# con = pymysql.connect(host='192.168.1.101', port=3306, user="feng", passwd='', db='rssminer')


class MainHandler(tornado.web.RequestHandler):
    # def __init__(self, *args, **argv):
    #     super(MainHandler, self).__init__(*args, **argv)
    #
    #     # tornado.web.RequestHandler.__init__(*args, **argv)
    #     self.conn = None

    def get(self):
        self._auto_finish = False
        pool.submit(task, self)
        # cur = con.cursor()
        # cur.execute("select * from users limit 10")
        #
        # # cur.execute("select * from dealrank_ctr limit 10")
        #
        # a = []
        # for r in cur.fetchall():
        #     a.append(str(r))
        # cur.close()
        # self.write("".join(a))
        # self.write("hello world")


application = tornado.web.Application([
    (r"/", MainHandler),
])

"""
single:
Running 10s test @ http://127.0.0.1:8888/
  4 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    23.24ms   36.76ms 337.99ms   97.73%
    Req/Sec     0.00      0.00     0.00    100.00%
  3366 requests in 10.00s, 2.65MB read
Requests/sec:    336.48
Transfer/sec:    271.42KB

threaded:
  4 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    18.35ms   23.87ms 347.52ms   99.27%
    Req/Sec     0.00      0.00     0.00    100.00%
  4698 requests in 10.00s, 3.70MB read
Requests/sec:    469.67
Transfer/sec:    378.86KB

pypy, single:
Running 10s test @ http://127.0.0.1:8888/
  4 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    15.90ms   42.72ms 263.67ms   95.85%
    Req/Sec     0.00      0.00     0.00    100.00%
  6995 requests in 10.00s, 5.51MB read
Requests/sec:    699.43
Transfer/sec:    564.19KB

pypy, threaded:
 4 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.95ms    1.22ms  10.38ms   72.41%
    Req/Sec     0.00      0.00     0.00    100.00%
  14270 requests in 10.00s, 11.24MB read
Requests/sec:   1426.77
Transfer/sec:      1.12MB
"""

if __name__ == "__main__":
    # import tornado.httpserver
    # server = tornado.httpserver.HTTPServer(application)
    # server.bind(8888)
    # server.start(0)  # autodetect number of cores and fork a process for each
    # tornado.ioloop.IOLoop.instance().start()
    # print "------------begine"
    application.listen(8888)
    tornado.ioloop.IOLoop.instance().start()
