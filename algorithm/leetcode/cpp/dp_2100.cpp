//
// Created by 裴星鑫 on 2022/3/6.
//
#include <vector>

#include "gtest/gtest.h"

using namespace std;

class Solution {
 public:
  vector<int> goodDaysToRobBank(vector<int>& security, int time) {
    int n = security.size();
    vector<int> left(n);
    vector<int> right(n);

    for (int i = 1; i < n; i++) {
      if (security[i] <= security[i - 1]) left[i] = left[i-1] + 1;
      if (security[n - i - 1] <= security[n - i]) right[n - i - 1] = right[n - i] + 1;
    }

    vector<int> res;
    for (int i = time; i < n - time; i++) {
      if (left[i] >= time && right[i] >= time) res.push_back(i);
    }
    return res;
  }
};

TEST(Solution, goodDaysToRobBank_DP) {
  vector<int> security = {5,3,3,3,5,6,2};
  Solution s;
  auto r = s.goodDaysToRobBank(security, 2);
  ASSERT_EQ(2, r.size());
  ASSERT_EQ(2, r[0]);
  ASSERT_EQ(3, r[1]);
}