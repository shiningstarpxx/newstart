//
// Created by 裴星鑫 on 2022/1/24.
//

#include <vector>

using namespace std;

class Solution {
 public:
  int networkDelayTime(vector<vector<int>>& times, int n, int k) {
    n_ = n;
    w_ = vector<vector<int>>(n + 1, vector<int>(n + 1, INT32_MAX / 2));
    for (int i = 1; i <= n; i++) w_[i][i] = 0;

    for (auto& v : times) {
      w_[v[0]][v[1]] = v[2];
    }

    floyd();

    int ans = 0;
    for (int i = 1; i <= n; i++) ans = max(ans, w_[k][i]);
    return ans >= INT32_MAX / 2 ? -1 : ans;
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

  vector<vector<int>> w_;
  int n_;
};