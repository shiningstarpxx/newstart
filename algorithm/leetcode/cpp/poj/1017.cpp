//
// Created by 裴星鑫 on 2022/10/6.
//
#include <cstdio>
#include <iostream>

using namespace std;

int main() {
  int a, b, c, d, e, f;
  int box2[4] = {0, 5, 3, 1};
  int box1[4] = {0, 7, 6, 5};
  while (true) {
    scanf("%d%d%d%d%d%d", &a, &b, &c, &d, &e, &f);
    if (a+b+c+d+e+f == 0) break;
    int res = d + e + f + (c + 3) / 4;
    int b2 = 5 * d + box2[c % 4];
    if (b2 < b) res += (b - b2 + 8) / 9;
    int a1 = 36 * res - 36 * f - 25 * e - 16 * d - 9 * c - 4 * b;
    if (a > a1) res += (a - a1 + 35) / 36;
//    printf("%d\n", res);
    cout << res << endl;
  }

  return 0;
}
