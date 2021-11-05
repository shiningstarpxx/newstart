//
// Created by 裴星鑫 on 2021/11/5.
//

class Solution {
 public:
  int longestSubsequence(vector<int>& arr, int difference) {
    unordered_map<int, int> dp;
    int ans = 0;
    for (auto d : arr) {
      dp[d] = dp[d - difference] + 1;
      ans = max(ans, dp[d]);
    }
    return ans;
  }
};

