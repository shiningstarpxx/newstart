//
// Created by 裴星鑫 on 2021/12/7.
//

#include <vector>

using namespace std;

// 跟flood fill类似（lt200）， 但是变成了阅读理解
// 上下左右都有相同链接的，叫连通分量，题目是找不满足这个条件的连通分量
class Solution {
 public:
  vector<vector<int>> colorBorder(vector<vector<int>>& grid, int row, int col, int color) {
    int rows = grid.size();
    int cols = grid[0].size();
    visited_ = vector<vector<bool>>(rows, vector<bool>(cols));

    if (grid[row][col] != color) {
        dfs(grid, rows, cols, row, col, grid[row][col]);
        for (auto c : cache_) grid[c.first][c.second] = color;
    }
    return grid;
  }

  void dfs(vector<vector<int>>& grid, int rows, int cols, int r, int c, int old_color) {
    if (r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != old_color || visited_[r][c]) return;
    if (checkDirections(grid, rows, cols, r, c)) cache_.push_back({r, c});
    visited_[r][c] = true;
    dfs(grid, rows, cols, r - 1, c, old_color);
    dfs(grid, rows, cols, r + 1, c, old_color);
    dfs(grid, rows, cols, r, c - 1, old_color);
    dfs(grid, rows, cols, r, c + 1, old_color);
  }

  bool checkDirections(vector<vector<int>>& grid, int rows, int cols, int r, int c) {
    int direc[4][2] = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
    for (int i = 0; i < 4; i++) {
      int rr = r + direc[i][0];
      int cc = c + direc[i][1];
      if (rr < 0 || rr >= rows || cc < 0 || cc >= cols || grid[rr][cc] != grid[r][c]) return true;
    }
    return false;
  }

  vector<pair<int, int>> cache_;
  vector<vector<bool>> visited_;
};