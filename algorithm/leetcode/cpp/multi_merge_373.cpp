//
// Created by 裴星鑫 on 2022/1/14.
//
#include <queue>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<int>> kSmallestPairs(vector<int>& nums1, vector<int>& nums2, int k) {
    auto cmp = [&](pair<int, int>& a, pair<int, int>& b) {
      return nums1[a.first] + nums2[a.second] > nums1[b.first] + nums2[b.second];
    };

    priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(cmp)> q(cmp);
    int n = nums1.size();
    int m = nums2.size();

    for (int i = 0; i < n && i < k; i++) q.emplace(i, 0);

    vector<vector<int>> res;
    int count = 0;
    while (count < k && !q.empty()) {
      auto p = q.top();
      q.pop();
      res.push_back({nums1[p.first], nums2[p.second]});
      count++;
      if (p.second + 1 < m) q.emplace(p.first, p.second + 1);
    }
    return res;
  }
};

