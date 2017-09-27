#!/usr/bin/env python

# dpw@alameda.local
# 2017.09.27

import socket

HOST = "localhost"
PORT = 3001

class Unique:
    def __init__(self, host, port):
        self.host = host
        self.port = port
        self.sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    def open(self):
        self.sock.connect((self.host, self.port))

    def send(self, cmd):
        self.sock.send(cmd)
        data = self.sock.recv(64)
        print 'received', repr(data)

    def close(self):
        self.sock.close()


if __name__ == '__main__':
    u = Unique(HOST, PORT)
    u.open()
    u.send("uuid")
    u.send("ulid")
    u.send("guid")
    u.send("tsid")
    u.send("txid")
    u.send("bytes")
    u.send("version")
    u.close()
