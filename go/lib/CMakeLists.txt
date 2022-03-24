cmake_minimum_required(VERSION 3.6)

project(cpp_go)

set(SRC ${SRC} simple_lib.cpp )

add_library(${PROJECT_NAME}_static STATIC ${SRC})
target_include_directories(${PROJECT_NAME}_static PRIVATE ".")

add_library(${PROJECT_NAME}_shared SHARED ${SRC})
target_include_directories(${PROJECT_NAME}_shared PRIVATE ".")
