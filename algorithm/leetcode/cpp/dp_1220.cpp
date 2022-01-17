//
// Created by 裴星鑫 on 2022/1/17.
//

#include <ctype.h>
#include <vector>

using namespace std;

class Solution {
 public:
  int countVowelPermutation(int n) {
    vector<vector<int64_t>> dp(n + 1, vector<int64_t>(5));
    for (int i = 0; i < 5; i++) dp[1][i] = 1;

    int mod = 1e9 + 7;
    for (int i = 1; i < n; i++) {
      dp[i + 1][1] += dp[i][0];  // a e
      dp[i + 1][0] += dp[i][1];  // e a
      dp[i + 1][2] += dp[i][1];  // e i
      dp[i + 1][0] += dp[i][2];  // i a
      dp[i + 1][1] += dp[i][2];  // i e
      dp[i + 1][3] += dp[i][2];  // i o
      dp[i + 1][4] += dp[i][2];  // i u
      dp[i + 1][2] += dp[i][3];  // o i
      dp[i + 1][4] += dp[i][3];  // o u
      dp[i + 1][0] += dp[i][4];  // u a
      for (int j = 0; j < 5; j++) dp[i + 1][j] = dp[i + 1][j] % mod;
    }

    int64_t sum = 0;
    for (int i = 0; i < 5; i++) sum += dp[n][i];
    return sum %= mod;
  }
};