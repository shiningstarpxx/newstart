//
// Created by 裴星鑫 on 2022/5/11.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int lengthOfLIS(vector<int>& nums) {
    // [1, 2, 3, 4, 5]
    // [1] -> [1, 2] -> [1, 2, 3] .... nums[i] > list.back() ,push
    // [1, 3, 4, 2, 5, 3] + [4, 5]
    // [1] -> [1, 3] -> [1, 3, 4]  尴尬的时刻， 2来了怎么办？ 如果直接弹掉 [1, 2, 5] 最长的 [1, 2, 4, 5]
    // [1, 3, 4] -> [1, 2, 4]  =
    // 分两种情况证明  在最终解的
    // 1. back 大，直接push，不影响最终解. 5 [1, 2, 4, 5]
    // 2. back 小 , 【1， 2， 3， 5】 -> [1 2, 3,4 ,5]

    vector<int> ans;  // 有序
    for (int i = 0; i < nums.size(); i++) {
      if (ans.empty() || nums[i] > ans.back()) {
        ans.push_back(nums[i]);
      } else {
        // auto it = lower_bound(ans.begin(), ans.end(), nums[i]);
        // *it = nums[i];
        int l = 0, r = ans.size();
        while (l < r) {
          int mid = l + (r - l) / 2;
          if (ans[mid] < nums[i]) {
            l = mid + 1;
          } else r = mid;
        }
        ans[l] = nums[i];
      }
    }

    return ans.size();
  }

  int dp(vector<int>& nums) {
    // dp[i] = max(dp[i], dp[j] + 1) if (j >= 0 && j < i && nums[i] > nums[j])
    int n = nums.size();
    int ans = 0;
    vector<int> dp(n, 1);
    for (int i = 0; i < n; i++) {
      for (int j = 0; j < i; j++)  // 无序，我如果变成有序，那么就可以二分了？？？？？？？？？
        if (nums[i] > nums[j]) {
          dp[i] = max(dp[i], dp[j] + 1);
        }
      ans = max(ans, dp[i]);
    }
    return ans;
  }
};