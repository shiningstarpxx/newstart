//
// Created by 裴星鑫 on 2022/3/9.
//
#include <vector>

#include "gtest/gtest.h"

using namespace std;

class Solution {
 public:
  int bestRotation(vector<int>& nums) {
    int n = nums.size();
    vector<int> diffs(n + 1);

    auto add = [&](int low, int high) {
      diffs[low]++;
      diffs[high+1]--;
    };

    for (int i = 0; i < n; i++) {
      int low = (i + 1) % n;
      int high = (i - nums[i] + n) % n;
      if (low < high) {
        add(low, high);
      } else {
        add(0, high);
        add(low, n - 1);
      }
    }

    for (int i = 1; i <= n; i++) diffs[i] += diffs[i - 1];
    int score = diffs[0];
    int res = 0;
    for (int i = 1; i <= n; i++) {
      if (diffs[i] > score) {
        res = i;
        score = diffs[i];
      }
    }
    return res;
  }
};

TEST(Solution, bestRotation) {
  vector<int> nums = {2,3,1,4,0};
  Solution s;
  EXPECT_EQ(3, s.bestRotation(nums));
}