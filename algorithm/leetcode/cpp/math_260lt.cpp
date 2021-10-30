//
// Created by 裴星鑫 on 2021/10/30.
//
#include <vector>

#include "gtest/gtest.h"

using namespace std;

class Solution {
 public:
  vector<int> singleNumber(vector<int> &nums) {
    int xor_data = 0;
    for (int i = 0; i < nums.size(); i++) {
      xor_data ^= nums[i];
    }

    int l = (xor_data == INT_MIN) ? xor_data : (xor_data & -xor_data);

    int first = 0;
    int second = 0;
    for (int i = 0; i < nums.size(); i++) {
      if (nums[i] & l) {
        first ^= nums[i];
      } else {
        second ^= nums[i];
      }
    }
    return {first, second};
  }
};

TEST(Solution, singleNumber) {
  Solution s;
  vector<int> d{1, 2, 1, 3, 2, 5};
  auto r = s.singleNumber(d);
  ASSERT_EQ(2, r.size());
  ASSERT_EQ(3, r[0]);
  ASSERT_EQ(5, r[1]);
}
