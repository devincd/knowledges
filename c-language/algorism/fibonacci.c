#include <stdio.h>

// index 表示求数列中第 index 个位置上的数的值
int fibonacci(int index) {
    // 设置结束递归的限制条件
    if (index == 1 || index == 2) {
        return 1;
    }

    return fibonacci(index-1) + fibonacci(index -2);
}

// 以非递归方式实现 fibonacci 函数，调用一次就可以生成长度为 n 的斐波那契数列
int normalFibonacci(int n) {
    int num1 = 1;
    int num2 = 1;

    int i, nextNum;
    for (i = 1; i <= n; i++) {
        printf("%d ", num1);
        nextNum = num1 + num2;
        num2 = nextNum;
        num1 = num2;
    }
    return 0;
}

int main() {
    int i;
    for (i = 1; i <= 10; i++) {
        printf("%d ", fibonacci(i)); 
    }
    print("\n");
    normalFibonacci(20);
    return 0;
}