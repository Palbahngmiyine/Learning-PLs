#include <iostream>

int main()
{
    int x = 1;
    x = x + 2;
    std::cout << x << std::endl;

    int y = x;
    std::cout << y << std::endl;

    // is (x + y) l-value or right value?
    std::cout << x + y << std::endl;

    std::cout << x << std::endl;

    // prints runtime error on cmake debug mode
    // int z;
    // std::cout << z << std::endl;

    return 0;
}
