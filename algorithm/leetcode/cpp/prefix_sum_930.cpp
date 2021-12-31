//
// Created by Pei Xingxin on 31/12/2021.
//
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
 public:
  int numSubarraysWithSum(vector<int>& nums, int goal) {
    int ans = 0;
    int prefix_sum = 0;
    unordered_map<int, int> cache;
    cache[0] = 1;
    for (int i = 0; i < nums.size(); i++) {
      prefix_sum += nums[i];
      if (cache.count(prefix_sum - goal)) ans += cache[prefix_sum - goal];
      cache[prefix_sum]++;  // why here? 因为如果prefix_sum - goal == prefix_sum 会导致上面多加一次，所以一定要放在后面
    }
    return ans;
  }
};