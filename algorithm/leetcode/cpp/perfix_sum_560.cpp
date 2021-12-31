//
// Created by Pei Xingxin on 31/12/2021.
//
#include <vector>
#include <unordered_map>

using namespace std;


// why prefix -k ?
// ---------------|----------|
// prefix - k     |     k    |
//            prefix

class Solution {
 public:
  int subarraySum(vector<int>& nums, int k) {
    int ans = 0;
    unordered_map<int, int> cache;
    cache[0] = 1;  // why? 因为prefix - k == 0的时候，需要加1, 表示找到了一个解
    int prefix_sum = 0;
    for (int i = 0; i < nums.size(); i++) {
      prefix_sum += nums[i];
      if (cache.count(prefix_sum - k)) ans += cache[prefix_sum -k];
      cache[prefix_sum]++;
    }
    return ans;
  }
};