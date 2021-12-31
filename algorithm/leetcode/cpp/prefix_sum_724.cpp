//
// Created by Pei Xingxin on 31/12/2021.
//
#include <algorithm>
#include <vector>

using namespace std;

class Solution {
 public:
  int pivotIndex(vector<int>& nums) {
    int sum = accumulate(nums.begin(), nums.end(), 0);
    int prefix_sum = 0;
    for (int i = 0; i < nums.size(); i++) {
      if (prefix_sum == sum - prefix_sum - nums[i]) return i;
      prefix_sum += nums[i];
    }
    return -1;
  }
};