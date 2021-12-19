//
// Created by 裴星鑫 on 2021/12/19.
//

class Solution {
 public:
  int findJudge(int n, vector<vector<int>>& trust) {
    vector<int> in_degree(n + 1);
    vector<int> out_degree(n + 1);
    for (auto& v : trust) {
      in_degree[v[1]]++;
      out_degree[v[0]]++;
    }

    for (int i = 1; i < in_degree.size(); i++) {
      if (in_degree[i] == n - 1 && out_degree[i] == 0) return i;
    }
    return -1;
  }

  int set_solution(int n, vector<vector<int>>& trust) {
    if (n == 1 && trust.empty()) return 1;

    unordered_map<int, unordered_set<int>> trusted;
    unordered_map<int, int> trust_count;

    for (auto& v : trust) {
      trusted[v[1]].emplace(v[0]);
      trust_count[v[0]]++;
    }

    int index = 0;
    for (auto it = trusted.begin(); it != trusted.end(); it++) {
      if (it->second.size() == n - 1) {
        if (index != 0) return -1;
        index = it->first;
      }
    }

    if (trust_count[index] == 0 && index != 0) return index;
    return -1;
  }
};