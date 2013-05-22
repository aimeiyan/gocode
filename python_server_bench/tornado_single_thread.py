import tornado.ioloop
import tornado.web
import pymysql
import threading

con = pymysql.connect(host='192.168.2.243', port=3306, user='q3boy', passwd='123', db='dealrank')

# con = pymysql.connect(host='192.168.1.101', port=3306, user="feng", passwd='', db='rssminer')

class MainHandler(tornado.web.RequestHandler):
    # def __init__(self, *args, **argv):
    #     super(MainHandler, self).__init__(*args, **argv)
    #
    #     # tornado.web.RequestHandler.__init__(*args, **argv)
    #     self.conn = None

    def get(self):

        # print self.request.connection.address

        # cur = con.cursor()
        # # cur.execute("select * from users limit 10")
        # cur.execute("select * from dealrank_ctr limit 10")
        #
        # a = []
        # for r in cur.fetchall():
        #     a.append(str(r))
        # cur.close()
        # self.write("".join(a))
        self.write("hello world")


application = tornado.web.Application([
    (r"/", MainHandler),
    ])

"""
python: Mon May 20 08:15:09 CST 2013
 Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    20.52ms    1.51ms  26.88ms   69.81%
    Req/Sec     0.00      0.00     0.00    100.00%
  3920 requests in 10.00s, 7.11MB read
Requests/sec:    391.87
Transfer/sec:    728.26KB

pypy: Mon May 20 08:14:10 CST 2013
 Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.58ms    3.73ms  66.10ms   99.26%
    Req/Sec     0.00      0.00     0.00    100.00%
  14448 requests in 10.00s, 26.22MB read
Requests/sec:   1444.56
Transfer/sec:      2.62MB
"""

if __name__ == "__main__":
    # import tornado.httpserver
    # server = tornado.httpserver.HTTPServer(application)
    # server.bind(8888)
    # server.start(0)  # autodetect number of cores and fork a process for each
    # tornado.ioloop.IOLoop.instance().start()

    application.listen(8888)
    tornado.ioloop.IOLoop.instance().start()
