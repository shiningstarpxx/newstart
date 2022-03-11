//
// Created by 裴星鑫 on 2022/3/11.
//
#include <functional>
#include <vector>
#include <unordered_map>

#include "gtest/gtest.h"

using namespace std;

class Solution {
 public:
  int countHighestScoreNodes(vector<int>& parents) {
    int n = parents.size();
    long max_score = 0;
    int res = 0;

    unordered_map<int, vector<int>> children;

    function<int(int)> dfs;
    dfs = [&](int node) -> int {
      long score = 1;
      int size = n - 1;
      for (auto v : children[node]) {
        int r = dfs(v);
        score *= r;
        size -= r;
      }
      if (node != 0) score *= size;
      if (score == max_score) res++;
      else if (score >= max_score) {
        res = 1;
        max_score = score;
      }
      return n - size;
    };

    for (int i = 0; i < parents.size(); i++) {
      if (parents[i] != -1)
        children[parents[i]].push_back(i);
    }
    dfs(0);
    return res;
  }
};

TEST(Solution, countHighestScoreNoeds) {
  Solution s;
  vector<int> scores = {-1,2,0,2,0};
  EXPECT_EQ(3, s.countHighestScoreNodes(scores));
}
