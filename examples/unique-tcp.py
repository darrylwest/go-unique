#!/usr/bin/env python

# dpw@alameda.local
# 2017.09.27

import socket
import time

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
        print(cmd, repr(data.rstrip()))

    def close(self):
        self.sock.close()


if __name__ == '__main__':
    u = Unique(HOST, PORT)
    u.open()
    while True:
        cmds = [ b"uuid", b"ulid", b"guid", b"tsid", b"txid", b"cuid", b"xuid", b"bytes", b"ping", b"version" ]
        for cmd in cmds:
            u.send(cmd)

        time.sleep(2)

    u.close()

