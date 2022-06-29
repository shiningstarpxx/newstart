//
// Created by 裴星鑫 on 2022/6/11.
//

#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  int minFlipsMonoIncr(string s) {
//      return dp(s);
//      return dp2(s);
    return presum(s);
  }

  int presum(string& s) {
    int n = s.size();
    vector<int> presum(n + 1);
    for (int i = 1; i <= n; i++) presum[i] = presum[i-1] + (s[i-1] - '0');

    int ans = n;
    for (int i = 1; i <= n; i++) {
      int left = presum[i-1], right = presum[n] - presum[i];
      ans = min(ans, left + (n - i - right));
    }
    return ans;
  }

  int dp(string& s) {
    int n = s.size();
    vector<vector<int>> dp(n, vector<int>(2));
    if (s[0] == '0') dp[0][1] = 1;
    if (s[0] == '1') dp[0][0] = 1;

    for (int i = 1; i < n; i++) {
      if (s[i] == '0') {
        dp[i][0] = dp[i-1][0];
        dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1;
      } else {
        dp[i][0] = dp[i-1][0] + 1;
        dp[i][1] = min(dp[i-1][0],dp[i-1][1]);
      }
    }
    return min(dp[n-1][0], dp[n-1][1]);
  }

  int dp2(string& s) {
    int n = s.size();
    int zero = 0, one = 0;
    if (s[0] == '0') one = 1;
    if (s[0] == '1') zero = 1;
    for (int i = 1; i < n; i++) {
      int a = zero, b = one;
      if (s[i] == '0') {
        zero = a;
        one = min(a, b) + 1;
      } else {
        zero = a + 1;
        one = min(a, b);
      }
    }
    return min(zero, one);
  }
};