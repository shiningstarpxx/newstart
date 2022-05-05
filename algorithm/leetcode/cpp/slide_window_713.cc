//
// Created by 裴星鑫 on 2022/5/5.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int numSubarrayProductLessThanK(vector<int>& nums, int k) {
    UseFor(nums, k);
  }

  int UseWhile(vector<int>& nums, int k) {
    int l = 0, r = 0;
    int multi = 1;
    int ans = 0;
    while (r < nums.size()) {
      multi *= nums[r];
      while (l <= r && multi >= k) multi /= nums[l++];
      ans += r - l + 1;
      r++;  // 这个放在后面很丑
    }

    return ans;
  }

  int UseFor(vector<int>& nums, int k) {
    int ans = 0;
    for (int l = 0, r = 0, multi = 1; r < nums.size(); r++) {
      multi *= nums[r];
      while (l <= r && multi >= k) multi /= nums[l++];

      ans += r - l + 1;  // 核心算法，求连续数组，在左边界线固定的时候，到右边界数据的个数就是 (r - l + 1)
    }
  }
};