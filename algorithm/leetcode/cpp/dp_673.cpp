//
// Created by 裴星鑫 on 2022/4/4.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int findNumberOfLIS(vector<int>& nums) {
    int n = nums.size();
    vector<int> dp(n, 1);
    vector<int> g(n, 1);
    int ans = 0;
    int max_count = 1;
    for (int i = 1; i < n; i++) {
      for (int j = i - 1; j >= 0; j--) {
        if (nums[i] > nums[j]) {
          if (dp[i] == dp[j] + 1) {
            g[i] += g[j];
          } else if (dp[i] < dp[j] + 1) {
            dp[i] = dp[j] + 1;
            g[i] = g[j];
          }
        }
      }
      max_count = max(max_count, dp[i]);
      // cout << endl;
    }
    for (int i = 0; i < n; i++) {
      if (dp[i] == max_count) ans += g[i];
    }
    return ans;
  }
};