//
// Created by Pei Xingxin on 31/12/2021.
//
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
 public:
  bool checkSubarraySum(vector<int>& nums, int k) {
    int ans = 0;
    int prefix_sum = 0;
    unordered_map<int, int> cache;
    cache[0] = -1;  // 确保了index 0 减去这个值，是一个数，依次类推；计算index时的技巧
    for (int i = 0; i < nums.size(); i++) {
      prefix_sum += nums[i];
      int d = (prefix_sum % k + long(k)) % k;
      if (cache.count(d) ) {
        if(i - cache[d] >= 2) return true;
        else continue;
      }
      cache[d] = i;
    }
    return false;
  }
};

