//
// Created by 裴星鑫 on 2022/3/29.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int jump(vector<int>& nums) {
    return Greed(nums);
  }

  int Greed(vector<int>& nums) {
    int n = nums.size(), right = 0, end = 0, ans = 0;
    for (int i = 0; i < n - 1; i++) {  // trick
      right = max(right, nums[i] + i);
      if (i == end) {
        end = right;
        ans++;
      }
    }
    return ans;
  }

  int DP(vector<int>& nums) {  // TLE
    int n = nums.size();
    vector<int> dp(n, INT_MAX);
    dp[0] = 0;

    for (int i = 1; i < n; i++) {
      for (int j = i - 1; j >= 0; j--) {
        if (dp[j] != INT_MAX && nums[j] + j >= i)
          dp[i] = min(dp[i], dp[j] + 1);
      }
    }
    return dp.back();
  }
};

