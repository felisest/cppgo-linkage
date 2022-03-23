#include <vector>
#include <iostream>
#include "libgo/libgoex.h"


int main() {

    std::cout << "Running ..." << std::endl;

    std::string out_str = "Hello from GO!";
    GoString go_str{out_str.c_str(), (long)out_str.length()};
    PrintString(go_str);

    std::vector<char> out_vector{'A','B','C','D','E','F','G','H'};
    GoSlice go_vector{out_vector.data(), static_cast<GoInt>(out_vector.size()), static_cast<GoInt>(out_vector.size())};
    ReverseSlice(go_vector);

    for (auto i = 0; i < go_vector.len; i++) std::cout << ((char*)go_vector.data)[i];
    std::cout << "\n";

    return 1;
}
