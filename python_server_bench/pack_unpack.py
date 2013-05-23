__author__ = 'feng'

import cPickle
import pickle
import time
import json
import logging

FORMAT = '%(asctime)-15s  %(message)s'
logging.basicConfig(format=FORMAT, level=logging.INFO)


def f(n):
    if n < 2:
        return 1
    return f(n - 1) + f(n - 2)

if __name__ == "__main__":
    a = {"asdfsd": 11212.11, "b": range(100, 200), "c": {"a": "---------"}, "dictkey": [1.1, 1.4334, 1.2121]}
    logging.info("data: %s", a)

    loops = 4000
    st = time.time()

    for i in range(1, loops):
        pickle.loads(pickle.dumps(a))

    # 2.13ms
    logging.info("pickle takes time %s ms", time.time() - st)

    st = time.time()
    for i in range(1, loops):
        # s = cStringIO.StringIO()
        cPickle.loads(cPickle.dumps(a))

    # 0.23ms
    logging.info("cpickle takes time %s ms", time.time() - st)
    # print "cpickle", time.time() - st

    st = time.time()
    for i in range(1, loops):
        json.loads(json.dumps(a))

    # 0.172ms
    logging.info("json takes time %s ms", time.time() - st)
    # print "json", time.time() - st

    # data length, json: 586, cpickle: 743, pickle: 727
    logging.info("data length, json: %s, cpickle: %s, pickle: %s" % (
        len(json.dumps(a)), len(cPickle.dumps(a)), len(pickle.dumps(a))
    ))
