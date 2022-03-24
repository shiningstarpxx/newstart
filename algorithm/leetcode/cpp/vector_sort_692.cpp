//
// Created by 裴星鑫 on 2022/3/18.
//
#include <string>
#include <vector>
#include <unordered_map>
using namespace std;

class Solution {
 public:
  vector<string> topKFrequent(vector<string>& words, int k) {
    unordered_map<string, int> cache_;
    for (auto& s : words) cache_[s]++;

    vector<string> res;
    vector<pair<string, int>> data;
    for(auto [k, v] : cache_) data.push_back({k, v});

    auto cmp = [](pair<string, int>& a, pair<string, int>& b) {
      if (a.second == b.second) return a.first < b.first;
      return a.second > b.second;
    };
    sort(data.begin(), data.end(), cmp);

    for (int i = 0; i < k; i++) res.push_back(data[i].first);
    return res;
  }
};