//
// Created by 裴星鑫 on 2022/7/2.
//

#include <queue>
#include <vector>

using namespace std;
class Solution {
 public:
  int minRefuelStops(int target, int startFuel, vector<vector<int>>& stations) {
    priority_queue<int, vector<int>, less<>> q;

    int start = startFuel;
    int ans = 0;
    int index = 0;
    while (start < target) {
      if (index < stations.size() && start >= stations[index][0]) {
        q.push(stations[index][1]);
        index++;
      } else {
        if (q.empty()) return -1;
        start += q.top();
        q.pop();
        ans++;
      }
    }
    return ans;
  }

  int dp(int target, int startFuel, vector<vector<int>>& stations) {
    int n = stations.size();

    vector<long> dp(n + 1, 0);
    dp[0] = startFuel;

    for (int i = 0; i < n; i++) {
      for (int j = i; j >= 0; j--) {
        if (dp[j] >= stations[i][0])
          dp[j + 1] = max(dp[j + 1], dp[j] + stations[i][1]);
      }
    }

    for (int i = 0; i <= n; i++) {
      if (dp[i] >= target) return i;
    }
    return -1;
  }
};