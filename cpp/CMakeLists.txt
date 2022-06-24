cmake_minimum_required(VERSION 3.6)

project(cpp_go)

include(CheckCXXCompilerFlag)

CHECK_CXX_COMPILER_FLAG("-std=c++11" COMPILER_SUPPORTS_CXX11)

if(COMPILER_SUPPORTS_CXX11)
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11")
else()
    message(FATAL_ERROR "The compiler ${CMAKE_CXX_COMPILER} doesn't have C++11 support.")
endif()

message("CMAKE_CXX_FLAGS of ${PROJECT_NAME}: ${CMAKE_CXX_FLAGS}")

message(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}")
message(CMAKE_LINKER "${CMAKE_LINKER}")
message(CMAKE_CXX_LINK_EXECUTABLE " ${CMAKE_CXX_LINK_EXECUTABLE}")

set(SRC ${SRC} "get_libgo.cpp")
set(HEADERS ${HEADERS} "libgo/libgoex.h")

set(LIBS ${LIBS} goex)

add_executable(cpp_go ${SRC} ${HEADERS})
target_link_directories(cpp_go PUBLIC "./libgo")
target_link_libraries(cpp_go ${LIBS})
