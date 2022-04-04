//
// Created by 裴星鑫 on 2022/3/30.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int maxSubarraySumCircular(vector<int>& nums) {
    int sum = 0, cur_max = 0, cur_min = 0, max_sum = nums[0], min_sum = nums[0];

    for (auto& v : nums) {
      cur_max = v + max(cur_max, 0);
      max_sum = max(cur_max, max_sum);
      cur_min = min(cur_min + v, v);
      min_sum = min(cur_min, min_sum);
      sum += v;
    }

    return max_sum > 0 ? max(max_sum, sum - min_sum): max_sum;
  }

};