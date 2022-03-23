#! /bin/bash

rm -R build_release
mkdir build_release && cd build_release

cmake ..
cmake --build . --target all --verbose

echo
echo

./cpp_go