//
// Created by 裴星鑫 on 2022/3/23.
//

#include <iostream>

#include "gtest/gtest.h"

using namespace std;

class Solution {
public:
    int findKthNumber(int n, int k) {
      int res = 1;
      k--;
      while (k > 0) {
        int steps = GetSteps(res, n);
        if (steps <= k) {
          k -= steps;
          res++;
        } else {
          k--;
          res *= 10;
        }
      }
      return res;
    }

    int GetSteps(int curr, long n) {
      int steps = 0;
      long first = curr;
      long end = curr;
      while (first <= n) {  // 字典树构造，一个节点下面会是下一个区间的节点
        steps += min(end, n) - first + 1;
        first = first * 10;
        end = end * 10 + 9;
      }
      // 比如【10， 19】 =》 【100， 199】
      return steps;
    }

    int min(int a, int b) {
      return a <= b ? a : b;
    }
};

TEST(Solution, findKthNumber) {
  Solution s;
  EXPECT_EQ(1, s.findKthNumber(5, 1));
}