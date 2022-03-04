//
// Created by 裴星鑫 on 2022/2/28.
//

#include <vector>

#include "gtest/gtest.h"

using namespace std;

class Solution {
 public:
  int maximumRequests(int n, vector<vector<int>>& requests) {
    vector<int> delta(n);
    int ans = 0;
    int m = requests.size();
    for (int mask = 1; mask < (1 << m); mask++) {
      int cnt = __builtin_popcount(mask);
      if (cnt <= ans) continue;
      fill(delta.begin(), delta.end(), 0);
      for (int i = 0; i < m; i++) {
        if (mask & (1 << i)) {
          delta[requests[i][0]]++;
          delta[requests[i][1]]--;
        }
      }
      if (all_of(delta.begin(), delta.end(), [](int x) {
        return x == 0;
      })) {
        ans = cnt;
      }
    }
    return ans;
  }
};


TEST(Solution, maximumRequests) {
  Solution s;
  vector<vector<int>> request = {{0, 0}, {1, 2}, {2, 1}};
  EXPECT_EQ(3, s.maximumRequests(3, request));

  vector<vector<int>> requests1 = {{0,3},{3,1},{1,2},{2,0}};
  EXPECT_EQ(4, s.maximumRequests(4, requests1));
}
