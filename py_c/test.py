import test
from unittest import *

if __name__ == "__main__":
    suite = TestSuite()
    suite.addTest(makeSuite(test.ExampleTest))
    TextTestRunner().run(suite)
