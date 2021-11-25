//
// Created by 裴星鑫 on 2021/11/20.
//
#include <algorithm>
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
 public:
  int findLHS(vector<int>& nums) {
    unordered_map<int, int> cache;

    for (int i = 0; i < nums.size(); i++) {
      cache[nums[i]]++;
    }

    int ans = 0;
    for (int i = 0; i < nums.size(); i++) {
      if (cache.count(nums[i]+1)) {
        ans = max(ans, cache[nums[i]] + cache[nums[i] + 1]);
      }
    }
    return ans;
  }
};
