from distutils.core import setup, Extension
import glob

setup(name="example",
      version="1.0",
      packages=["example"],
      # py_modules = ['example'],
      ext_modules = [
          Extension("_example",
                    sources=glob.glob("src/*.c"),
                    # ["pyexample.c","example.c"]
                )
      ]
)
