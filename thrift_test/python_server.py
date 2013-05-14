import sys
sys.path.append('./gen-py')

from tutorial.ttypes import UserProfile
from tutorial.UserStorage import Processor

from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol
from thrift.server import TServer

class UserStorageHandler:

    def __init__(self):
        self.users = []

    def store(self, user):
        self.users.append(user)

    def retrieve(self, id):
        return self.users[id]


processor = Processor(UserStorageHandler())
transport = TSocket.TServerSocket(port=9090)

tfactory = TTransport.TBufferedTransportFactory()
pfactory = TBinaryProtocol.TBinaryProtocolFactory()

server = TServer.TSimpleServer(processor, transport, tfactory, pfactory)

server.serve()
print 'done.'
