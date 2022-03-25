//
// Created by 裴星鑫 on 2022/3/25.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int minFallingPathSum(vector<vector<int>>& grid) {
    return BruteForceDP(grid);
  }

  int BruteForceDP(vector<vector<int>>& grid) {
    int r = grid.size();
    int c = grid[0].size();

    for (int i = 1; i < r; i++) {
      for (int j = 0; j < c; j++) {
        int from = INT_MAX;
        for (int k = 0; k < c; k++) {
          if (k != j) from = min(from, grid[i - 1][k]);
        }
        grid[i][j] += from;

      }
    }

    int ans = INT_MAX;
    for (int i = 0; i < c; i++) ans = min(ans, grid[r - 1][i]);
    return ans;
  }
};