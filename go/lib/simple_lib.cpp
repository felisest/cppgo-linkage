#include <iostream>

extern "C" void print_hello(char* in_str) {
    std::cout << in_str << "\n";
}

extern "C" void reverse_vector(void* in_buff, int size) {

    int r = size;
    for (int i = 0; i < size; i++) {

        std::swap(*((char*)in_buff + r), *((char*)in_buff + i));

		if (--r-i <= 1) break;
    }
}

