//
// Created by 裴星鑫 on 2022/1/21.
//
#include <queue>
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
 public:
  int minJumps(vector<int>& arr) {
    unordered_map<int, vector<int>> sameValues;
    for (int i = 0; i < arr.size(); i++) {
      sameValues[arr[i]].push_back(i);
    }

    deque<pair<int, int>> q;
    unordered_map<int, int> visited;
    visited[0] = 1;
    q.push_back({0, 0});
    while (!q.empty()) {
      auto [idx, step] = q.front();
      q.pop_front();
      if (idx == arr.size() - 1) return step;

      int val = arr[idx];
      if (sameValues.count(val)) {
        for (auto v : sameValues[val]) {
          if (!visited.count(v)) {
            q.push_back({v, step + 1});
            visited[v] = 1;
          }
        }
        sameValues.erase(val);
      }
      if (idx + 1 < arr.size() && !visited.count(idx + 1)) {
        q.push_back({idx + 1, step + 1});
        visited[idx + 1] = 1;
      }
      if (idx - 1 >= 0 && !visited.count(idx - 1)) {
        q.push_back({idx - 1, step + 1});
        visited[idx - 1] = 1;
      }
    }
    return -1;
  }
};