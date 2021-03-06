//
// Created by 裴星鑫 on 2022/3/25.
//
#include <vector>

using namespace std;

class Solution {
 public:
  static const int MOD = 1e9 + 7;

  int countRoutes(vector<int>& locations, int start, int finish, int fuel) {
    return DP(locations, start, finish, fuel);
  }

  int dfs(vector<int>& locations, int start, int finish, int fuel) {
    int size = locations.size();
    dp_ = vector<vector<int>> (size, vector<int>(fuel + 1, -1));
    return dfs(locations, start, finish, fuel);
  }

  int dfsHelper(vector<int>& locations, int start, int finish, int fuel) {
    if (dp_[start][fuel] != -1) return dp_[start][fuel];

    int need = abs(locations[start] - locations[finish]);
    if (need > fuel) {
      dp_[start][fuel] = 0;
      return 0;
    }

    int sum = start == finish ? 1 : 0;  //base condition
    for (int i = 0; i < locations.size(); i++) {
      if (i != start) {
        int need = abs(locations[start] - locations[i]);
        if (need > fuel) continue;
        sum += dfsHelper(locations, i, finish, fuel - need);
        sum %= MOD;
      }
    }
    dp_[start][fuel] = sum;
    return sum;
  }

  vector<vector<int>> dp_;


  int DP(vector<int>& locations, int start, int finish, int fuel) {
    int n = locations.size();
    vector<vector<int>> dp(n, vector<int>(fuel + 1, 0));

    for (int k = 0; k <= fuel; k++) dp[finish][k] = 1;

    for (int k = 0; k <= fuel; k++) {
      for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
          if (j != i) {
            int need = abs(locations[j] - locations[i]);
            if (need > k) continue;
            dp[i][k] += dp[j][k - need];
            dp[i][k] %= MOD;
          }
        }
      }
    }

    return dp[start][fuel];
  }
};