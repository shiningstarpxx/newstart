//
// Created by 裴星鑫 on 2022/4/6.
//
#include <vector>

using namespace

class Solution {
 public:
  vector<int> findMinHeightTrees(int n, vector<vector<int>>& edges) {
    if (n == 1) return {0};

    vector<vector<int>> graph(n);
    vector<int> degree(n);

    for (auto &v : edges) {
      graph[v[0]].push_back(v[1]);
      graph[v[1]].push_back(v[0]);
      degree[v[0]]++;
      degree[v[1]]++;
    }

    vector<int> q;
    for (int i = 0; i < n; i++)
      if (degree[i] == 1) q.push_back(i);

    vector<int> ans;
    while (!q.empty()) {
      ans.clear();
      vector<int> tmp;
      for (int i = 0; i < q.size(); i++) {
        ans.emplace_back(q[i]);
        vector<int> &v = graph[q[i]];
        for (auto vv : v) {
          --degree[vv];
          if (degree[vv] == 1) tmp.push_back(vv);
        }
      }
      q = tmp;
    }

    return ans;
  }
};