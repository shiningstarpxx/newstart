//
// Created by 裴星鑫 on 2022/5/13.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int rob(vector<int>& nums) {
    return LessSpace(nums);
  }

  int MoreSpace(vector<int>& nums) {
    if (nums.size() <= 1) return nums[0];
    int n = nums.size();
    vector<vector<int>> dp(n, vector<int>(2));  // 0 是不取，1是取
    dp[0][0] = 0; dp[0][1] = nums[0];
    for (int i = 1; i < n; i++) {
      dp[i][0] = max(dp[i-1][0], dp[i-1][1]);
      dp[i][1] = dp[i-1][0] + nums[i];
    }
    return max(dp[n-1][0], dp[n-1][1]);
  }

  int LessSpace(vector<int>& nums) {
    if (nums.size() <= 1) return nums[0];
    int n = nums.size();
    int not_use = 0, use = nums[0];
    for (int i = 1; i < n; i++) {
      int a = max(not_use, use);
      int b = not_use + nums[i];
      not_use = a;
      use = b;
    }
    return max(not_use, use);
  }
};