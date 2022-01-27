//
// Created by 裴星鑫 on 2022/1/24.
//
#include <stdio.h>
#include <vector>

#include "gtest/gtest.h"

using namespace std;

class Solution {
 public:
  int networkDelayTime(vector<vector<int>>& times, int n, int k) {
    n_ = n;
    k_ = k;

    return calcMatrix(times);
  }

  // 邻接矩阵
  int calcMatrix(vector<vector<int>>& times) {
    w_ = vector<vector<int>>(n_ + 1, vector<int>(n_ + 1, INT32_MAX / 2));
    for (int i = 1; i <= n_; i++) w_[i][i] = 0;

    for (auto& v : times) {
      w_[v[0]][v[1]] = v[2];
      w_[v[1]][v[0]] = v[2];
    }

//    return calcFloyd();
    return calcDijkstra();
  }

  void floyd() {
    for (int k = 1; k <= n_; k++) {
      for (int i = 1; i <= n_; i++) {
        for (int j = 1; j <= n_; j++) {
          w_[i][j] = min(w_[i][j], w_[i][k] + w_[k][j]);
        }
      }
    }
  }

  int calcFloyd() {
    floyd();
    int ans = 0;
    for (int i = 1; i <= n_; i++) {
      ans = max(ans, w_[k_][i]);
    }
    return ans >= INT32_MAX / 2 ? -1 : ans;
  }

  void dijkstra() {
    printf("***\n");
    visited_ = vector<bool>(n_ + 1, false);
    dist_ = vector<int>(n_ + 1, INT32_MAX / 2);
    dist_[k_] = 0;

    for (int p = 1; p <= n_; p++) {
      int t = -1;
      for (int i = 1; i <= n_; i++) {
        if (!visited_[i] && (t == -1 || dist_[i] < dist_[t])) t = i;
      }
      visited_[t] = true;
      printf("%d\n", t);
      for (int i = 1; i <= n_; i++) {
        dist_[i] = min(dist_[i], dist_[t] + w_[t][i]);
        printf("%d ", dist_[i]);
      }
      printf("\n");
    }
  }

  int calcDijkstra() {
    dijkstra();

    int ans = 0;
    for (int i = 1; i <= n_; i++) ans = max(ans, dist_[i]);
    return ans >= INT32_MAX / 2 ? -1 : ans;
  }

  vector<vector<int>> w_;
  vector<bool> visited_;
  vector<int> dist_;
  int n_;
  int k_;
};

TEST(Solution, dijkstra) {
  Solution s;
  vector<vector<int>> times = {{2, 1, 1}, {2, 3, 1}, {3, 4, 1}};
  int r = s.networkDelayTime(times, 4, 2);
  ASSERT_EQ(2, r);
  ASSERT_EQ(3, r);
}