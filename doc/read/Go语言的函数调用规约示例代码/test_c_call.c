// test_c_call.c
// 定义一个接受 6 个以上参数的函数，确保存入寄存器和栈
int my_c_function(int a, int b, int c, int d, int e, int f, int g) {
    return a + b + c + d + e + f + g;
}

int main() {
    // 调用函数，传入特定数值方便在汇编中查找
    // 0x11 -> RDI
    // 0x22 -> RSI
    // 0x33 -> RDX
    // 0x44 -> RCX
    // 0x55 -> R8
    // 0x66 -> R9
    // 0x77 -> Stack
    my_c_function(0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77);
    return 0;
}
