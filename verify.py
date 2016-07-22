#!/usr/bin/env python
# -*- coding: utf-8 -*-

from passlib.hash import bcrypt
import sys

def verify():
	# 逐行verify并打印是否成功
    for line in sys.stdin.readlines():
        tokens = line.split()
        s = tokens[0]
        h = tokens[1]
        if bcrypt.verify(s, h):
            print s, h, "verify ok"
        else:
            print s, h, "verify fail"

if __name__ == '__main__':
    verify()

