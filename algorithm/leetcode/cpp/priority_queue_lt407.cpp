//
// Created by 裴星鑫 on 2021/11/3.
//
class Solution {
 public:
  int trapRainWater(vector<vector<int>>& heightMap) {
    int m = heightMap.size();
    int n = heightMap[0].size();
    priority_queue<pair<int, int>, vector<pair<int, int>>, greater<pair<int, int>>> pq;
    vector<vector<bool>> visit(m, vector<bool>(n));
    for (int i = 0; i < m; i++) {
      for (int j = 0; j < n; j++) {
        if (i == 0 || i == m - 1 || j == 0 || j == n - 1) {
          pq.push({heightMap[i][j], i * n + j});
          visit[i][j] = true;
        }
      }
    }

    int res = 0;
    int dir[4][2] = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
    while (!pq.empty()) {
      auto pi = pq.top();
      pq.pop();
      for (int i = 0; i < 4; i++) {
        int x = pi.second / n + dir[i][0];
        int y = pi.second % n + dir[i][1];
        if (x >= 0 && x < m && y >= 0 && y < n && !visit[x][y]) {
          if (pi.first > heightMap[x][y]) res += pi.first - heightMap[x][y];
          pq.push({max(pi.first, heightMap[x][y]), x * n + y});
          visit[x][y] = true;
        }
      }
    }
    return res;
  }
};
