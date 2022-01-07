//
// Created by 裴星鑫 on 2022/1/7.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int minPathSum(vector<vector<int>>& grid) {
    int m = grid.size();
    int n = grid[0].size();
    vector<vector<long>> dp(m, vector<long>(n));
    for (int i = 0; i < m; i++) {
      for (int j = 0; j < n; j++) {
        if (i == 0 && j == 0) dp[i][j] = grid[i][j];
        else {
          dp[i][j] = min(i > 0? dp[i - 1][j]:INT_MAX, j > 0? dp[i][j-1] : INT_MAX) + grid[i][j];
        }
      }
    }
    return dp[m-1][n-1];
  }
};