//
// Created by 裴星鑫 on 2021/11/12.
//
#include <vector>

#include "gtest/gtest.h"

using namespace std;

// 算法思路： dp[i][j] =>  表示 【i, j]区间内的最小代价
// if i <= j 这个代价需要计算
// i > j 这个代价显然就是0, 用默认初始化值即可
class Solution {
 public:
  int getMoneyAmount(int n) {
    vector<vector<int>> dp(n + 1, vector<int>(n + 1));
    for (int i = n; i >= 1; i--) {
      for (int j = i + 1; j <= n; j++) {
        int tmp = INT_MAX;
        for (int k = i; k < j; k++) {
          tmp = min(tmp, k + max(dp[i][k-1], dp[k+1][j]));
        }
        dp[i][j] = tmp;
      }
    }
    return dp[1][n];
  }
};

TEST(Solution, getMoneyAmount) {
  Solution s;
  ASSERT_EQ(0, s.getMoneyAmount(1));
  ASSERT_EQ(1, s.getMoneyAmount(2));
  ASSERT_EQ(16, s.getMoneyAmount(10));
}