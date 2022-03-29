//
// Created by 裴星鑫 on 2022/3/29.
//
#include <vector>

using namespace std;

class Solution {
 public:
  bool canJump(vector<int>& nums) {
//        return DP(nums);
    return Greedy(nums);
  }

  bool DP(vector<int>& nums) {
    int n = nums.size();
    vector<int> dp(n, 0);
    dp[0] = 1;
    for (int i = 1; i < n; i++) {
      for (int j = i - 1; j >= 0; j--) {
        if (dp[j] && nums[j] + j >= i) {
          dp[i] = 1;
          break;
        }
      }
    }
    return dp.back();
  }

  bool Greedy(vector<int>& nums) {
    int n = nums.size();
    int right = 0;
    for (int i = 0; i < n; i++) {
      if (i > right) continue;
      right = max(right, nums[i] + i);
      if (right >= n - 1) return true;
    }
    return false;
  }
};
