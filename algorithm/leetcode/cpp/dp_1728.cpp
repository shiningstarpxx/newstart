//
// Created by 裴星鑫 on 2022/5/10.
//
#include <string>
#include <vector>
using namespace std;

class Solution {
  const static int N = 1000;
  const static int S = 8 * 8 * 8 * 8;
 public:
  bool canMouseWin(vector<string>& grid, int catJump, int mouseJump) {
    row_ = grid.size();
    column_ = grid[0].size();
    cat_jump_ = catJump;
    mouse_jump_ = mouseJump;
    grid_ = grid;

    int x, y, m, n;
    for (int i = 0; i < row_; i++) {
      for (int j = 0; j < column_; j++) {
        if (grid[i][j] == 'M') {
          x = i, y = j;
        } else if (grid[i][j] == 'C') {
          m = i, n = j;
        } else if (grid_[i][j] == 'F'){
          f_x_ = i, f_y_ = j;
        }
      }
    }

    return dfs(x, y, m, n, 0) == 0;
  }
  int CalcState(int x, int y, int m, int n) {
    return (x << 9) | (y << 6) | (m << 3) | n;
  }
  int row_, column_, f_x_, f_y_;
  int cat_jump_, mouse_jump_;
  vector<string> grid_;

  vector<vector<int>> dirs_ = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
  vector<vector<int>> dp_ = vector<vector<int>>(S, vector<int>(N, -1));
  int dfs(int x, int y, int m, int n, int step) {
    int state = CalcState(x, y, m, n);
    if (step == N - 1) return dp_[state][step] = 1;
    if (x == f_x_ && y == f_y_) return dp_[state][step] = 0;
    if (m == f_x_ && n == f_y_) return dp_[state][step] = 1;
    if (x == m && y == n) return dp_[state][step] = 1;
    if (dp_[state][step] != -1) return dp_[state][step];

    if (step % 2 == 0) {  // mouse
      for (int i = 0; i < dirs_.size(); i++) {
        for (int j = 0; j <= mouse_jump_; j++) {
          int a = x + dirs_[i][0] * j, b = y + dirs_[i][1] * j;
          if (a < 0 || a >= row_ || b < 0 || b >= column_) break;
          if (grid_[a][b] == '#') break;
          if (dfs(a, b, m, n, step + 1) == 0) return dp_[state][step] = 0;
        }
      }
      return dp_[state][step] = 1;
    } else {
      for (int i = 0; i < dirs_.size(); i++) {
        for (int j = 0; j <= cat_jump_; j++) {
          int a = m + dirs_[i][0] * j, b = n + dirs_[i][1] * j;
          if (a < 0 || a >= row_ || b < 0 || b >= column_) break;
          if (grid_[a][b] == '#') break;
          if (dfs(x, y, a, b, step + 1) == 1) return dp_[state][step] = 1;
        }
      }
      return dp_[state][step] = 0;
    }
  }
};