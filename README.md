# travis-qemu-example [![Build Status](https://travis-ci.org/jdub/travis-qemu-example.svg?branch=master)](https://travis-ci.org/jdub/travis-qemu-example)

Example use of modern qemu on Travis CI

- `ci` is a submodule pointing to [`travis-qemu`](https://github.com/jdub/travis-qemu)
- `.travis.yml` builds a simple Go program for arm and i386, and runs it with qemu
