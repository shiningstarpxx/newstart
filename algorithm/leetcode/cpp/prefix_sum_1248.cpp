//
// Created by Pei Xingxin on 31/12/2021.
//
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
 public:
  int numberOfSubarrays(vector<int>& nums, int k) {
    int ans = 0;
    int prefix_sum = 0;
    unordered_map<int, int> d;
    d[0] = 1;
    for (int i = 0; i < nums.size(); i++) {
      prefix_sum += nums[i] & 0x1;
      if (d.count(prefix_sum - k)) ans += d[prefix_sum - k];
      d[prefix_sum]++;
    }
    return ans;
  }
};

