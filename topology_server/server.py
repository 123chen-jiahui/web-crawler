import grpc
from concurrent import futures
import test_pb2
import test_pb2_grpc
import networkx as nx
import matplotlib.pyplot as plt
import string
import random

class Draw(test_pb2_grpc.DrawServicer):
    def Topology(self, request, context):
        G = nx.DiGraph()
        m = request.G
        print(m)
        urls = request.Contents
        print(urls)
        total = request.Total
        print(total)
        for i in range(total):
            print(i, urls[i])
        for i in range(total):
            G.add_node(i)
            m.setdefault(i)
            to = m[i].Nodes
            if len(to) == 0:
                continue
            for j in range(len(to)):
                G.add_edge(i, to[j])

        #  plt.savefig("../file_server/assert/tmp.png")
        nx.draw(G, with_labels=True)
        #  plt.savefig("/home/chen/tmp.png")
        filePath = ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(5))
        filePath += ".png"
        plt.savefig("../file_server/assert/" + filePath)
        #  plt.savefig("tmp.png")
        plt.close()

        print("hello world")
        filePath = "http://localhost:8081/static/" + filePath
        return test_pb2.TopologyReply(Ok=True, FilePath=filePath)

if __name__ == '__main__':
    # 1 实例化Server
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    # 2 注册逻辑到server中
    test_pb2_grpc.add_DrawServicer_to_server(Draw(), server)
    # 3 启动server
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()
