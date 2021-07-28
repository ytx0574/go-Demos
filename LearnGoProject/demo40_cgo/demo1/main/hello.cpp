


#include <iostream>

extern "C" {
    //todo 表示函数的链接符号遵循C的规则
    #include "hello.h"
}

//todo 实现C的SayHello函数
void SayHello(const char *s) {
    std::cout << "C++  begin\n";
	std::cout << s;
	std::cout << "C++  end\n";
}