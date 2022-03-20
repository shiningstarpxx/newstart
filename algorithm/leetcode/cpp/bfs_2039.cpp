//
// Created by 裴星鑫 on 2022/3/20.
//
#include <vector>
#include <deque>
using namespace std;

class Solution {
 public:
  int networkBecomesIdle(vector<vector<int>>& edges, vector<int>& patience) {
    int n = patience.size();
    vector<vector<int>> graph(n);

    for (int i = 0; i < edges.size(); i++) {
      graph[edges[i][0]].push_back(edges[i][1]);
      graph[edges[i][1]].push_back(edges[i][0]);
    }

    vector<bool> visited(n);

    deque<int> q;
    q.push_back(0);
    visited[0] = true;

    int dist = 1;
    int res = 0;
    while (!q.empty()) {
      int size = q.size();
      for (int i = 0; i < size; i++) {
        int c = q.front();
        q.pop_front();

        auto& v = graph[c];
        for (int j = 0; j < v.size(); j++) {
          if (!visited[v[j]]) {
            visited[v[j]] = true;
            int t = (2 * dist - 1) / patience[v[j]] * patience[v[j]] + 2 * dist + 1;
            res = max(res, t);
            q.push_back(v[j]);
          }
        }
      }
      dist++;
    }

    return res;
  }
};
