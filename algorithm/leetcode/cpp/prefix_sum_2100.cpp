//
// Created by 裴星鑫 on 2022/3/6.
//
#include <vector>

#include "gtest/gtest.h"

using namespace std;

class Solution {
 public:
  vector<int> goodDaysToRobBank2(vector<int> &security, int time) {
    int n = security.size();
    vector<int> global(n);
    for (int i = 1; i < n; i++) {
      if (security[i] == security[i - 1]) continue;
      global[i] = security[i] > security[i - 1] ? 1 : -1;
    }

    vector<int> big(n + 1);
    vector<int> small(n + 1);
    for (int i = 1; i <= n; i++) big[i] = big[i - 1] + (global[i - 1] == 1 ? 1 : 0);
    for (int i = 1; i <= n; i++) small[i] = small[i - 1] + (global[i - 1] == -1 ? 1 : 0);

    vector<int> res;
    for (int i = time; i < n - time; i++) {
      int l = big[i + 1] - big[i - time + 1];
      int r = small[i + time + 1] - small[i + 1];
      if (l == 0 && r == 0) res.push_back(i);
    }
    return res;
  }
};

TEST(Solution, goodDaysToRobBank_Prefix) {
  vector<int> security = {5,3,3,3,5,6,2};
  Solution s;
  auto r = s.goodDaysToRobBank2(security, 2);
  ASSERT_EQ(2, r.size());
  ASSERT_EQ(2, r[0]);
  ASSERT_EQ(3, r[1]);
}