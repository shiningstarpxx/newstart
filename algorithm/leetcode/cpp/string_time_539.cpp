//
// Created by 裴星鑫 on 2022/1/18.
//
#include <vector>

using namespace std;
#include "gtest/gtest.h"

class Solution {
 public:
  int findMinDifference(vector<string>& timePoints) {
    if (timePoints.size() > DayMinute) return 0;

    sort(timePoints.begin(), timePoints.end());
    int preMinutes = calcMinitue(timePoints[0]);
    int ans = INT_MAX;
    for (int i = 1; i < timePoints.size(); i++) {
      ans = min(ans, calcMinitue(timePoints[i]) - preMinutes);
      preMinutes = calcMinitue(timePoints[i]);
    }
    ans = min(ans, calcMinitue(timePoints[0]) + DayMinute - preMinutes);
    return ans;
  }

 private:
  int calcMinitue(string& t) {
    return ((t[0] - '0') * 10 + (t[1] - '0')) * 60 + (t[3] - '0') * 10 + (t[4] - '0');
  }

  const int DayMinute = 1440;
};

TEST(Solution, findMinDifference) {
  vector<string> timePoints {"00:00", "00:01"};
  Solution s;
  EXPECT_EQ(1, s.findMinDifference(timePoints));
}

TEST(Solution, findMinDifference_Board) {
  vector<string> timePoints {"00:00", "23:59"};
  Solution s;
  EXPECT_EQ(1, s.findMinDifference(timePoints));
}

TEST(Solution, findMinDifference_Board2) {
  vector<string> timePoints {"23:59", "00:00"};
  Solution s;
  EXPECT_EQ(1, s.findMinDifference(timePoints));
}