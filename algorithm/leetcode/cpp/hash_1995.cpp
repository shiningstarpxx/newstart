//
// Created by 裴星鑫 on 2021/12/29.
//
#include <unordered_map>
#include <vector>

using namespace std;
class Solution {
 public:
  int countQuadruplets(vector<int>& nums) {
    // return bruteForce(nums);
    // return simpleHash(nums);
    return improvedHash(nums);
  }

  int bruteForce(vector<int>& nums) {
    int ans = 0;

    for (int i = 0; i < nums.size(); i++) {
      for (int j = i + 1; j < nums.size(); j++) {
        for (int m = j + 1; m < nums.size(); m++) {
          for (int n = m +1; n < nums.size(); n++) {
            if (nums[n] == nums[m] + nums[i] + nums[j]) ans++;
          }
        }
      }
    }
    return ans;
  }

  int simpleHash(vector<int>& nums) {  // n^3
    int ans = 0;
    unordered_map<int, int> cache;

    for (int c = nums.size() - 2; c >= 2; c--) {
      cache[nums[c+1]]++;
      for (int i = 0; i < c; i++) {
        for (int j = i + 1; j < c; j++) {
          int d = nums[c] + nums[i] + nums[j];
          if (cache.count(d)) ans += cache[d];
        }
      }
    }
    return ans;
  }

  int improvedHash(vector<int>& nums) { // n^2
  int ans = 0;
    unordered_map<int, int> cache;
    for (int b = nums.size() - 3; b >= 1; b--) {
      for (int d = b+2; d < nums.size(); d++) {
        cache[nums[d] - nums[b+1]]++;
      }
      for (int a = 0; a < b; a++) {
        if (cache.count(nums[a] + nums[b])) ans += cache[nums[a] + nums[b]];
      }
    }
    return ans;
  }
};