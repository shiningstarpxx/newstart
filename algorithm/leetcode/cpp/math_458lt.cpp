//
// Created by 裴星鑫 on 2021/11/25.
//

#include <math.h>

#include "gtest/gtest.h"

class Solution {
 public:
  int poorPigs(int buckets, int minutesToDie, int minutesToTest) {
    int k = minutesToTest / minutesToDie + 1;
    int ans = ceil(log(buckets) / log(k));
    return ans;
  }
};

TEST(Solution, poorPigs) {
  Solution s;
  ASSERT_EQ(5, s.poorPigs(1000, 15, 60));
}

