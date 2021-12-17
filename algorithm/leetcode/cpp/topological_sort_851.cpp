//
// Created by 裴星鑫 on 2021/12/17.
//
#include <numeric>
#include <vector>
#include <queue>

#include "gtest/gtest.h"

using namespace std;

class Solution {
 public:
  vector<int> loudAndRich(vector<vector<int>>& richer, vector<int>& quiet) {
    int size = quiet.size();
    vector<int> ans(size);
    iota(ans.begin(), ans.end(), 0);

    vector<vector<int>> richer_list(size);
    vector<int> in_degree(size);
    for (auto& v : richer) {
      richer_list[v[0]].emplace_back(v[1]);
      in_degree[v[1]]++;
    }

    deque<int> q;
    for (int i = 0; i < size; i++) {
      if (in_degree[i] == 0) q.push_back(i);
    }

    while (!q.empty()) {
      int t = q.front();
      q.pop_front();
      for (auto d : richer_list[t]) {
        if (quiet[ans[d]] > quiet[ans[t]]) ans[d] = ans[t];
        in_degree[d]--;
        if (in_degree[d] == 0) {
          q.push_back(d);
        }
      }
    }
    return ans;
  }
};

TEST(Solution, loudAndRich) {
  Solution s;
  // [[1,0],[2,1],[3,1],[3,7],[4,3],[5,3],[6,3]]
  // [3,2,5,4,6,1,7,0]
  // [5,5,2,5,4,5,6,7]
  vector<vector<int>> richer {
    {1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}
  };
  vector<int> quiet{3, 2, 5, 4, 6, 1, 7, 8};
  auto t = s.loudAndRich(richer, quiet);
  ASSERT_EQ(8, t.size());
  ASSERT_EQ(5, t[0]);
  ASSERT_EQ(5, t[1]);

}