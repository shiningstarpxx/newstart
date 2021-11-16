//
// Created by 裴星鑫 on 2021/11/15.
//

#include <cmath>

class Solution {
 public:
  int bulbSwitch(int n) {
    int sum = 0;
    for (int i = 1; i <= sqrt(n); i++) sum++;
    return sum;
  }
};

