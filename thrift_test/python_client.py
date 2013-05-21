import sys

sys.path.append('./gen-py')

from tutorial.ttypes import UserProfile
from tutorial.UserStorage import Client

from thrift.transport import TSocket
from thrift.protocol import TBinaryProtocol
from time import time


transport = TSocket.TSocket("localhost", 9090)
transport.open()
protocol = TBinaryProtocol.TBinaryProtocol(transport)

service = Client(protocol)

up = UserProfile(uid=1,
                 ids=range(1, 100),
                 name="Test User",
                 blurb="Thrift is great")
service.store(up)

print "warm up pypy, 10k loops"
for i in range(1, 10000):
    service.retrieve(0)

print "warm up pypy done"

start = time()
for i in range(1, 100):
    service.retrieve(0)

t = (time() - start) * 1000
print t, t / 100

""" pypy is much faster than python
server python:
python client: 69.1139698029 0.691139698029 => 0.69ms per request
pypy client:   33.2369804382 0.332369804382

server pypy
python client: 51.5530109406 0.515530109406
pypy client:   17.6401138306 0.176401138306
"""
