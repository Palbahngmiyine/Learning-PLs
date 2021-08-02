#include <iostream>

using namespace std;

int main() {

    int x = 1024;
    double pi = 3.141592;
    int y = 1;

    // Omit std
    cout << "Before your input, y was " << y << endl;

    cin >> y;

    cout << x << endl;

    // equals to cout << "end line test" << endl;
    cout << "end line test\n";

    cout << "y is " << y << " pi is " << pi << endl;

    // \t do spacing with smart indenting
    cout << "abc" << "\t" << "def" << endl;
    cout << "ab" << "\t" << "cdef" << endl;

    return 0;
}
