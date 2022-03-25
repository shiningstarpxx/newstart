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

  void UpdateFirstAndSecond(pair<int, int>& first, pair<int, int>& second, int index, int value) {
    if (value <= first.second) {
      second = first;
      first.first = index;
      first.second = value;
    } else if (value <= second.second) {
      second.first = index;
      second.second = value;
    }
  }

  int ImprovedDP(vector<vector<int>>& grid) {
    int r = grid.size();
    int c = grid[0].size();

    pair<int, int> first = {-1, INT_MAX};
    pair<int, int> second = {-1, INT_MAX};
    for (int j = 0; j < c; j++)  UpdateFirstAndSecond(first, second, j, grid[0][j]);

    for (int i = 1; i < r; i++) {
      pair<int, int> first_t = {-1, INT_MAX};
      pair<int, int> second_t = {-1, INT_MAX};
      for (int j = 0; j < c; j++) {
        if (j != first.first) {
          grid[i][j] += first.second;
        } else grid[i][j] += second.second;
        UpdateFirstAndSecond(first_t, second_t, j, grid[i][j]);
      }
      first = first_t;
      second = second_t;
    }

    int ans = INT_MAX;
    for (int i = 0; i < c; i++) ans = min(ans, grid[r - 1][i]);
    return ans;
  }
};