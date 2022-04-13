//
// Created by 裴星鑫 on 2022/4/13.
//
#include <numeric>
#include <vector>

using namespace std;

class Solution {
 public:
  bool canPartition(vector<int>& nums) {
    return OneDemision(nums);
  }

  // TLE
  bool Helper(vector<int>& nums, int index, int target, int sum) {
    if (sum > target) return false;
    if (sum == target) return true;
    if (index >= nums.size()) return false;

    return Helper(nums, index + 1, target, sum + nums[index]) || Helper(nums, index + 1, target, sum);
  }

  bool TwoDemisions(vector<int>& nums) {
    if (nums.size() < 2) return false;
    int sum = accumulate(nums.begin(), nums.end(), 0);
    if (sum % 2 != 0) return false;

    int target = sum / 2;
    // return Helper(nums, 0, target, 0);

    vector<vector<bool>> dp(nums.size() + 1, vector<bool>(target + 1));
    dp[0][0] = true;

    for (int i = 1; i <= nums.size(); i++) {
      for (int j = 0; j <= target; j++) {
        if (dp[i-1][j]) dp[i][j] = true;
        else if (j >= nums[i-1]) {
          dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]];
        }
      }
    }
    return dp[nums.size()].back();
  }

  bool OneDemision(vector<int>& nums) {
    if (nums.size() < 2) return false;
    int sum = accumulate(nums.begin(), nums.end(), 0);
    if (sum % 2 != 0) return false;

    int target = sum / 2;

    vector<bool> dp(target + 1);
    dp[0] = true;

    for (int i = 0; i < nums.size(); i++) {
      vector<bool> tmp = dp;
      for (int j = 1; j <= target; j++) {
        if (j >= nums[i])
          tmp[j] = dp[j] || dp[j - nums[i]];
      }
      dp = tmp;
    }
    return dp.back();
  }
};