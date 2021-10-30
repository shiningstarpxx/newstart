//
// Created by 裴星鑫 on 2021/10/26.
//

class Solution {
 public:
  vector<int> nextGreaterElement(vector<int> &nums1, vector<int> &nums2) {
    reverse(nums2.begin(), nums2.end());

    unordered_map<int, int> cache;
    vector<int> stack;
    for (int i = 0; i < nums2.size(); i++) {
      while (!stack.empty() && stack.back() < nums2[i]) stack.pop_back();
      if (stack.empty()) cache[nums2[i]] = -1;
      else cache[nums2[i]] = stack.back();
      stack.push_back(nums2[i]);
    }

    vector<int> res;
    for (auto n : nums1) res.push_back(cache[n]);
    return res;
  }
};