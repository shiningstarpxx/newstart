//
// Created by 裴星鑫 on 2021/11/20.
//
#include <cmath>

using namespace std;

class Solution {
 public:
  int integerReplacement(int n) {
    return dfs(n);
  }

  int dfs(long n ) {
    if (n == 1) return 0;

    if (n % 2 == 0) return dfs(n / 2) + 1;
    int plus = dfs(n + 1) + 1;
    int minus = dfs(n - 1) + 1;
    return min(plus, minus);
  }
};