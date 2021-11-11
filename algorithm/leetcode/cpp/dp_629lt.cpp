//
// Created by 裴星鑫 on 2021/11/11.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int kInversePairs(int n, int k) {
    vector<vector<long>> dp(n + 1, vector<long>(k + 1));
    const int mod = 1e9 + 7;
    dp[0][0] = 1; dp[1][0] = 1;
    for (int i = 1; i <= n; i++) {
      for (int j = 0; j <= k; j++) {
        dp[i][j] = ((j - 1 >= 0 ? dp[i][j-1] : 0) + dp[i-1][j] - (j-i >= 0 ? dp[i-1][j-i] : 0)) % mod;
        if (dp[i][j] < 0) dp[i][j] += mod;
      }
    }

    return int(dp[n][k]);
  }
};