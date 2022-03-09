//
// Created by 裴星鑫 on 2022/3/9.
//
#include <vector>

#include "gtest/gtest.h"

using namespace std;

class Solution {
public:
    vector<int> corpFlightBookings(vector<vector<int>>& bookings, int n) {
        vector<int> diffs(n + 1);
        for (auto& v : bookings) {
            int l = v[0] - 1, r = v[1] - 1;
            diffs[l] += v[2];
            diffs[r + 1] -= v[2];
        }

        vector<int> ans(n);
        ans[0] = diffs[0];
        for (int i = 1; i < n; i++) ans[i] = ans[i - 1] + diffs[i];
        return ans;
    }
};

TEST(Solution, corpFlightBookings) {
  vector<vector<int>> bookings = {{1,2,10},{2,3,20},{2,5,25}};
  Solution s;
  auto r = s.corpFlightBookings(bookings, 5);
  EXPECT_EQ(5, r.size());
  EXPECT_EQ(10, r[0]);
  EXPECT_EQ(55, r[1]);
  EXPECT_EQ(45, r[2]);
  EXPECT_EQ(25, r[3]);
  EXPECT_EQ(25, r[4]);
}