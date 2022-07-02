//
// Created by 裴星鑫 on 2022/6/10.
//

#include <vector>

using namespace std;

class Solution {
 public:
  double largestSumOfAverages(vector<int>& nums, int k) {
    int n = nums.size();
    vector<int> pre_sum(n + 1);
    for (int i = 0 ; i < n; i++) pre_sum[i + 1] = pre_sum[i] + nums[i];

    vector<vector<double>> dp(n + 1, vector<double>(k+1));

    for (int i = 1; i <= n; i++) dp[i][1] = double (pre_sum[i]) / (i);
    for (int m = 2; m <= k; m++) {
      for (int i = 1; i <= n; i++) {
        for (int j = 1; j < i; j++) {
          dp[i][m] = max(dp[i][m], dp[j][m-1] + double (pre_sum[j] - pre_sum[i]) / (j-i));
        }
      }
    }

    return dp[n][k];
  }
};