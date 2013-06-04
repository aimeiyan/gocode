__author__ = 'feng'

from math import log
from random import random


# Skip Lists: A Probabilistic Alternative to Balanced Trees

class Node(object):
    def __init__(self, key, value, nexts):
        self.key = key
        self.value = value
        self.forward = nexts


class End(object):
    def __cmp__(self, other):
        return 1  # always greater than any other object


class Head(object):
    def __cmp__(self, other):
        return -1


END = Node(End(), None, []) # the end


class SkipListMap(object):
    def __init__(self, expected_size, p=0.25):
        self.maxlevel = int(log(expected_size, 1 / p))
        self.header = Node(Head(), None, [END] * self.maxlevel)
        self.p = p
        self.size = 0

    def rand_level(self):
        lev = 1
        while random() < self.p and lev < self.maxlevel:
            lev += 1
        return lev

    def insert(self, key, value):
        update = [None] * self.maxlevel
        x = self.header
        for i in reversed(range(len(x.forward))):
            while x.forward[i].key < key:
                x = x.forward[i]
            update[i] = x
        x = x.forward[0]
        if x.key == key:
            x.value = value
        else:
            new = Node(key, value, [None] * self.rand_level())
            for i in range(len(new.forward)):
                update[i].forward[i], new.forward[i] = new, update[i].forward[i]
            self.size += 1

    def lookup(self, key):
        x = self.header
        for i in reversed(range(len(x.forward))):
            while x.forward[i].key < key:
                x = x.forward[i]
        x = x.forward[0]
        if x.key == key:
            return x.value
        else:
            return False

    def delete(self, key):
        update = [None] * self.maxlevel
        x = self.header
        for i in reversed(range(len(x.forward))):
            while x.forward[i].key < key:
                x = x.forward[i]
            update[i] = x
        x = x.forward[0]
        if x.key == key:
            self.size -= 1
            for i in range(len(x.forward)):
                update[i].forward[i] = x.forward[i]
        else:
            return False


import unittest


class SkipListMapTest(unittest.TestCase):
    def test_insert(self):
        m = SkipListMap(1000)
        for i in range(10):
            m.insert(i, i)
        for i in range(10):
            m.insert(i, i)
        self.assertEqual(10, m.size)
        for i in range(10):
            self.assertEqual(m.lookup(i), i)

    def test_delete(self):
        m = SkipListMap(1000)
        for i in range(3):
            m.insert(1, 1)
            self.assertEqual(m.lookup(1), 1)
            m.delete(1)
            self.assertEqual(m.lookup(1), False)












