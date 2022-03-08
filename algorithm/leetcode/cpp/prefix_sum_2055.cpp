//
// Created by 裴星鑫 on 2022/3/8.
//
#include <vector>

#include "gtest/gtest.h"
using namespace std;

class Solution {
public:
    vector<int> platesBetweenCandles(string s, vector<vector<int>>& queries) {
        int n = s.size();
        vector<int> pre_sum(n);

        int sum = 0;
        for (int i = 0; i < n; i++) {
            if (s[i] == '*') sum++;
            pre_sum[i] = sum;
        }

        vector<int> left(n);
        int l = -1;
        for (int i = 0; i < n; i++) {
            if (s[i] == '|') l = i;
            left[i] = l;
        }

        vector<int> right(n);
        int r = -1;
        for (int i = n - 1; i >= 0; i--) {
            if (s[i] == '|') r = i;
            right[i] = r;
        }

        vector<int> res;
        for (int i = 0; i < queries.size(); i++) {
            int ll = right[queries[i][0]], rr = left[queries[i][1]];
            if (ll == - 1 || rr == -1 || rr <= ll) res.push_back(0);
            else res.push_back(pre_sum[rr] - pre_sum[ll]);
        }
        return res;
    }
};

TEST(Solution, platesBetweenCandles) {
  string s = "**|**|***|";
  vector<vector<int>> qs = {{2, 5}, {5, 9}};
  Solution sl;
  auto r = sl.platesBetweenCandles(s, qs);
  EXPECT_EQ(2, r.size());
  EXPECT_EQ(2, r[0]);
  EXPECT_EQ(3, r[1]);
}