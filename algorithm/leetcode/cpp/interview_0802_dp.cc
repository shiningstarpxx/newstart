// dp && construct route

class Solution {
 public:
  vector<vector<int>> pathWithObstacles(vector<vector<int>> &obstacleGrid) {
    vector<vector<int>> res;
    int r = obstacleGrid.size();
    int c = obstacleGrid[0].size();

    vector<vector<bool>> dp(r, vector<bool>(c));
    dp[0][0] = obstacleGrid[0][0] == 0;
    dp[r - 1][c - 1] = obstacleGrid[r - 1][c - 1] == 0;

    for (int i = 0; i < r; i++) {
      for (int j = 0; j < c; j++) {
        if (i == 0 && j == 0) continue;
        dp[i][j] = (obstacleGrid[i][j] == 0);
        bool flag = false;
        if (i > 0)
          flag |= dp[i - 1][j];
        if (j > 0)
          flag |= dp[i][j - 1];
        dp[i][j] = dp[i][j] && flag;
      }
    }

    if (!dp[r - 1][c - 1]) return res;
    res.push_back({r - 1, c - 1});

    for (int i = r - 1, j = c - 1; i > 0 || j > 0;) {
      if (i > 0 && dp[i - 1][j]) i--;
      else j--;
      res.push_back({i, j});
    }
    reverse(res.begin(), res.end());
    return res;
  }
};