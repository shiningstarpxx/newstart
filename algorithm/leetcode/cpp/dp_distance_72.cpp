//
// Created by 裴星鑫 on 2022/4/13.
//
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  int minDistance(string word1, string word2) {
    int m = word1.size();
    int n = word2.size();

    vector<vector<int>> dp(m + 1, vector<int>(n + 1,INT_MAX/2));
    // 两个空字符无需操作，是0
    for (int i = 0; i <= m; i++) dp[i][0] = i;
    for (int j = 0; j <= n; j++) dp[0][j] = j;

    for (int i = 0; i < m; i++) {
      for (int j = 0; j < n; j++) {
        // dp[i + 1][j + 1]
        if (word1[i] != word2[j]) {
          // dp[i+1][j] + 1 表示添加
          // dp[i][j+1] 表示删除
          // dp[i][j] 表示替换
          dp[i+1][j+1] = min(dp[i+1][j] + 1, min(dp[i][j+1]+1, dp[i][j] + 1));
        } else {
          dp[i+1][j+1] = min(dp[i+1][j] + 1, min(dp[i][j+1]+1, dp[i][j]));
        }

      }
    }
    return dp[m][n];
  }
};