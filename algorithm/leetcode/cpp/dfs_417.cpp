//
// Created by 裴星鑫 on 2022/4/27.
//
#include <vector>

using namespace std;

class Solution {
public:
    vector<vector<int>> pacificAtlantic(vector<vector<int>>& heights) {
      int m = heights.size();
      int n = heights[0].size();
      heights_ = heights;

      vector<vector<bool>> pacific(m, vector<bool>(n));
      vector<vector<bool>> atlantic(m, vector<bool>(n));

      for (int i = 0; i < m; i++) dfs(i, 0, m, n, pacific);
      for (int j = 0; j < n; j++) dfs(0, j, m, n, pacific);
      for (int i = 0; i < m; i++) dfs(i, n - 1, m, n, atlantic);
      for (int j = 0; j < n; j++) dfs(m - 1, j, m, n, atlantic);

      vector<vector<int>> res;
      for (int i = 0; i < m; i++) {
        for (int j = 0; j < n; j++) {
          if (pacific[i][j] && atlantic[i][j])
            res.push_back({i, j});
        }
      }
      return res;
    }

    void dfs(int x, int y, int m, int n, vector<vector<bool>>& visited) {
      if (visited[x][y]) return;
      visited[x][y] = true;
      static vector<vector<int>> dir = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
      for (int i = 0; i < dir.size(); i++) {
        int xx = x + dir[i][0];
        int yy = y + dir[i][1];
        if (xx >= 0 && xx < m && yy >= 0 && yy < n && heights_[xx][yy] >= heights_[x][y])
          dfs(xx, yy, m, n, visited);
      }
    }

    vector<vector<int>> heights_;

    vector<vector<int>> StackOverFlow(vector<vector<int>>& heights) {
      int m = heights.size();
      int n = heights[0].size();

      vector<vector<int>> ans;
      for (int i = 0; i < m; i++) {
        for (int j = 0; j < n; j++) {
          if (dfs(heights, i, j, m, n, heights[i][j]) == 3) {
            ans.push_back({i, j});
          }
        }
      }
      return ans;
    }

    int dfs(vector<vector<int>>& data, int x, int y, int m, int n, int last_data) {
        if (y < 0 || x < 0) {
           return 1;
        } else if (y >= n || x >= m) {
           return 2;
        }

        // cout << x << " " << y << " " << m << " " << n << endl;
        int d = data[x][y];
        if (d > last_data) return 0;

        int rr = 0;
        rr |= dfs(data, x - 1, y, m, n, d);
        rr |= dfs(data, x + 1, y, m, n, d);
        rr |= dfs(data, x, y - 1, m, n, d);
        rr |= dfs(data, x, y + 1, m, n, d);
        return rr;
    }
};