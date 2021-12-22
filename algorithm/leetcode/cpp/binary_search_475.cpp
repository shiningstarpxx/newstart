//
// Created by 裴星鑫 on 2021/12/20.
//
#include <algorithm>
#include <vector>

#include <iostream>

#include "gtest/gtest.h"
using namespace std;

class Solution {
 public:
  int findRadius(vector<int>& houses, vector<int>& heaters) {
    sort(heaters.begin(), heaters.end());
    int res = 0;
    for (auto h : houses) {
      int d = upper_bound(heaters.begin(), heaters.end(), h) - heaters.begin();
      int right_distance = d < heaters.size() ? heaters[d] - h : INT_MAX;
      int left_distance = (d - 1) >= 0 ? h - heaters[d - 1] : INT_MAX;
      cout << right_distance <<  ", " << left_distance << endl;
      res = max(res, min(right_distance, left_distance));
    }
    return res;
  }
};

TEST(Solutin, findRadius) {
  Solution s;
  vector<int> houses = {1, 2, 3};
  vector<int> heaters = {2};
  ASSERT_EQ(1, s.findRadius(houses, heaters));
}