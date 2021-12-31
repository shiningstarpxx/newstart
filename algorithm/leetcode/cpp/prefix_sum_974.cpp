//
// Created by Pei Xingxin on 31/12/2021.
//
#include <unordered_map>
#include <vector>

using namespace std;

// 同余定理
// (a - b) % k == 0  <==>  a % k == b % k
class Solution {
 public:
  int subarraysDivByK(vector<int>& nums, int k) {
    int ans = 0;
    int prefix_sum = 0;
    unordered_map<int, int> cache;
    cache[0] = 1;

    for (int i = 0; i < nums.size(); i++) {
      prefix_sum += nums[i];
      int d = (prefix_sum % k + k) % k;
      if (cache.count(d)) ans += cache[d];
      cache[d]++;  // 存储的是余数值
    }
    return ans;
  }
};
