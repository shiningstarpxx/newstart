//
// Created by 裴星鑫 on 2022/1/6.
//

#include <vector>

using namespace std;

class Solution {
 public:
  int uniquePathsWithObstacles(vector<vector<int>> &obstacleGrid) {
    int m = obstacleGrid.size();
    int n = obstacleGrid[0].size();
    if (obstacleGrid[0][0] == 1) return 0;

    vector<vector<int>> dp(m, vector<int>(n, 0));
    dp[0][0] = 1;
    for (int i = 0; i < m; i++) {
      for (int j = 0; j < n; j++) {
        if (obstacleGrid[i][j] != 1) {
          if (i > 0 && j > 0) {
            dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
          } else if (i > 0) {
            dp[i][j] = dp[i - 1][j];
          } else if (j > 0) {
            dp[i][j] = dp[i][j - 1];
          }
        }
      }
    }
    return dp[m - 1][n - 1];
  }
};
