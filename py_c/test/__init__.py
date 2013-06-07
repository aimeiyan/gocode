import sys, glob, os
from unittest import TestCase

version = sys.version.split(" ")[0]
majorminor = version[0:3]

path = glob.glob("build/lib*-%s/*.so" % majorminor)[0]
sys.path.insert(0, os.path.dirname(path))


from unittest import *
import example

class ExampleTest(TestCase):

    def test_gcd(self):
        self.assertEquals(example.gcd(2, 8), 2)

    def test_replace(self):
        self.assertEquals(example.replace("abc", 'a', 'b')[1], "bbc")
