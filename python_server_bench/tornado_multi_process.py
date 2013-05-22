import tornado.ioloop
import tornado.web
import pymysql
import threading


# conn = pymysql.connect(host='192.168.2.243', port=3306, user='q3boy', passwd='123', db='dealrank')

# conn = pymysql.connect(host='192.168.1.101', port=3306, user="feng", passwd='', db='rssminer')
conns = threading.local()
conns.con = None

class MainHandler(tornado.web.RequestHandler):
    # def __init__(self, *args, **argv):
    #     super(MainHandler, self).__init__(*args, **argv)
    #
    #     # tornado.web.RequestHandler.__init__(*args, **argv)
    #     self.conn = None

    def get(self):
        if conns.con is None:
            conns.con = pymysql.connect(host='192.168.1.101', port=3306, user="feng", passwd='', db='rssminer')

        cur = conns.con.cursor()
        cur.execute("select * from users limit 10")

        # cur.execute("select * from dealrank_ctr limit 10")

        a = []
        for r in cur.fetchall():
            a.append(str(r))
        cur.close()
        self.write("".join(a))
        # self.write("hello world")


application = tornado.web.Application([
    (r"/", MainHandler),
])

"""
python, Mon May 20 08:09:59 CST 2013
 Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.65ms  270.46us   5.83ms   86.94%
    Req/Sec     0.00      0.00     0.00    100.00%
  21771 requests in 10.00s, 39.51MB read
Requests/sec:   2176.85
Transfer/sec:      3.95MB

pypy, Mon May 20 08:10:32 CST 2013
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.66ms    7.04ms 116.16ms   99.63%
    Req/Sec     1.00k     0.00     1.00k   100.00%
  66058 requests in 10.00s, 119.88MB read
Requests/sec:   6605.56
Transfer/sec:     11.99MB
"""

if __name__ == "__main__":
    import tornado.httpserver
    server = tornado.httpserver.HTTPServer(application)
    server.bind(8888)
    server.start(0)  # autodetect number of cores and fork a process for each
    tornado.ioloop.IOLoop.instance().start()

    # application.listen(8888)
    # tornado.ioloop.IOLoop.instance().start()
