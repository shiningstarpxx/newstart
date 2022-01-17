//
// Created by 裴星鑫 on 2022/1/14.
//
#include <queue>
#include <vector>

using namespace std;

// 多路归并的典型应用
class Solution {
 public:
  int kthSmallest(vector<vector<int>>& matrix, int k) {
    auto cmp = [&](pair<int, int>& a, pair<int, int>& b) {
      return matrix[a.first][a.second] > matrix[b.first][b.second];
    };
    priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(cmp)> q(cmp);

    for (int i = 0; i < k && i < matrix.size(); i++) {
      q.emplace(i, 0);
    }

    while (k > 1) {
      auto p = q.top();
      q.pop();
      if (p.second + 1 < matrix.size()) q.emplace(p.first, p.second + 1);
      k--;
    }
    auto p = q.top();
    return matrix[p.first][p.second];
  }
};
