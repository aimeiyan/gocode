import sys
sys.path.append('./gen-py')

from tutorial.ttypes import UserProfile
from tutorial.UserStorage import Client

from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol
from thrift.server import TServer


transport = TSocket.TSocket("localhost", 9090)
transport.open()
protocol = TBinaryProtocol.TBinaryProtocol(transport)

service = Client(protocol)

up = UserProfile(uid=1,
                 name="Test User",
                 blurb="Thrift is great")
service.store(up)

print service.retrieve(0)
