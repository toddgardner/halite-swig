#!/usr/bin/env python

"""
setup.py file for SWIG example
"""

from distutils.core import setup, Extension

halite_module = Extension(
  '_halite',
  sources=['halite_wrap.cxx'],
  extra_compile_args=['-std=c++11'])

setup (
  name = 'halite',
  version = '0.1',
  author = "tmgardner + sojumu",
  description = """halite bindings?""",
  ext_modules = [halite_module],
  py_modules = ["halite"])
