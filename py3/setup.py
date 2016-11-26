#!/usr/bin/env python

"""
Halite engine warpped in SWIG
"""

from distutils.core import setup, Extension

halite_module = Extension(
  '_halite',
  sources=[
      'halite_wrap.cxx',
      'wrapped/core/Halite.cpp',
      'wrapped/halite-core.cpp',
      'wrapped/networking/Networking.cpp',
  ],
  extra_compile_args=['-std=c++11'])

setup (
  name = 'halite',
  version = '0.1',
  author = "tmgardner + sojumu",
  description = """halite bindings?""",
  ext_modules = [halite_module],
  py_modules = ["halite"])
