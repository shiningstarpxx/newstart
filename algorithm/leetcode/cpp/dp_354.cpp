//
// Created by 裴星鑫 on 2022/5/11.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int maxEnvelopes(vector<vector<int>>& envelopes) {
    sort(envelopes.begin(), envelopes.end(), [](vector<int>& a, vector<int>& b) {
      if (a[0] == b[0]) return a[1] > b[1];
      return a[0] < b[0];
    });

    // vector<int> ans;
    // for (int i = 0; i < envelopes.size(); i++) {
    //     if (ans.empty() || envelopes[i][1] > ans.back()) ans.push_back(envelopes[i][1]);
    //     else {
    //         auto it = lower_bound(ans.begin(), ans.end(), envelopes[i][1]);
    //         *it = envelopes[i][1];
    //     }
    // }
    // return ans.size();
    vector<vector<int>> ans;

    for (int i = 0; i < envelopes.size(); i++) {
      if (ans.empty() || envelopes[i][1] > ans.back()[1]) ans.push_back(envelopes[i]);
      else {
        int l = 0, r = ans.size() - 1;
        while (l < r) {
          int mid = l + (r - l) / 2;
          if (ans[mid][1] >= envelopes[i][1]) r = mid;
          else l = mid + 1;
        }
        ans[l] = envelopes[i];
      }
    }
    return ans.size();
  }

  int DP(vector<vector<int>>& envelopes) {  // TLE
    sort(envelopes.begin(), envelopes.end(), [](vector<int>& a, vector<int>& b) {
      if (a[0] == b[0]) return a[1] < b[1];
      return a[0] < b[0];
    });

    int ans = 0;
    int n = envelopes.size();
    vector<int> dp(n, 1);
    for (int i = 0; i < n; i++) {
      for (int j = 0; j < i; j++) {
        if (envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1]) {
          dp[i] = max(dp[i], dp[j] + 1);
        }
      }
      ans = max(dp[i], ans);
    }
    return ans;
  }
};