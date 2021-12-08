//
// Created by 裴星鑫 on 2021/12/8.
//

#include <algorithm>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<int> maxSumOfThreeSubarrays(vector<int>& nums, int k) {
    return dp(nums, k);
  }

  vector<int> dp(vector<int>& nums, int k) {
    int size = nums.size();
    vector<vector<int>> dp(size + 10, vector<int>(4));

    vector<int> pre_sum(nums.size() + 1);
    for (int i = 1; i <= nums.size(); i++) pre_sum[i] = nums[i - 1] + pre_sum[i - 1];

    for (int i = size - k + 1; i >= 1; i--) {
      for (int j = 1; j < 4; j++) {
        dp[i][j] = max(dp[i + 1][j], dp[i + k][j - 1] + pre_sum[i + k - 1] - pre_sum[i - 1]);
      }
    }

    vector<int> res(3);
    int i = 1, j = 3, idx = 0;
    while (j > 0) {
      if (dp[i + 1][j] > dp[i + k][j - 1] + pre_sum[i + k - 1] - pre_sum[i - 1]) {
        i++;
      } else {
        res[idx++] = i - 1;
        i += k; j--;
      }
    }
    return res;
  }

};

