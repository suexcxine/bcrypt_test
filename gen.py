#!/usr/bin/env python
# -*- coding: utf-8 -*-

from passlib.hash import bcrypt
from random import Random
import sys

def random_str(randomlength=8):
    str = ''
    chars = 'AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789'
    length = len(chars) - 1
    random = Random()
    for i in range(randomlength):
        str += chars[random.randint(0, length)]
    return str

def gen():
    # 设定生成随机字符串的数量
    n = 20
    if len(sys.argv) > 1:
        n = int(sys.argv[1])

    # 生成指定数量的随机字符串并生成hash, 写入stdout
    for x in range(0, n):
        s = random_str()
        h = bcrypt.encrypt(s, rounds=4)
        sys.stdout.write(s + ' ' + h + '\n') 
    sys.stdout.flush()

if __name__ == '__main__':
    gen()

