//
// Created by 裴星鑫 on 2022/1/5.
//
#include <vector>

#include "gtest/gtest.h"
using namespace std;

class Solution {
 public:
  int uniquePaths(int m, int n) {
    return ThreeLeaves(m, n);
  }

  int DP(int m, int n) {
    vector<vector<int>> dp(m, vector<int>(n, 0));
    for (int i = 0; i < m; i++) dp[i][0] = 1;
    for (int i = 0; i < n; i++) dp[0][i] = 1;
    for (int i = 1; i < m; i++) {
      for (int j = 1; j < n; j++) {
        dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
      }
    }
    return dp[m - 1][n - 1];
  }

  int ThreeLeaves(int m, int n) {
    vector<vector<int>> dp(m, vector<int>(n));
    dp[0][0] = 1;
    for (int i = 0; i < m; i++) {
      for (int j = 0; j < n; j++) {
        if (i > 0 && j > 0) dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
        else if (i > 0) dp[i][j] = dp[i - 1][j];
        else if (j > 0) dp[i][j] = dp[i][j - 1];
      }
    }
    return dp[m - 1][n - 1];
  }
};

TEST(Solution62, uniquePaths62) {
  Solution s;
  EXPECT_EQ(28, s.uniquePaths(3, 7));
  EXPECT_EQ(3, s.uniquePaths(3, 2));
}