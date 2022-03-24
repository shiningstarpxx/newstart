//
// Created by 裴星鑫 on 2022/3/24.
//
#include <vector>

#include "gtest/gtest.h"

using namespace std;

class Solution {
 public:
  int minimumTotal(vector<vector<int>>& triangle) {
    for (int i = triangle.size() - 2; i >= 0; i--) {
      for (int j = 0; j < triangle[i].size(); j++) {
        triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1]);
      }
    }
    return triangle[0][0];
  }
};

TEST(Solution, minimumTotal) {
  vector<vector<int>> d = {{2},{3,4},{6,5,7},{4,1,8,3}};
  Solution s;
  EXPECT_EQ(11, s.minimumTotal(d));
}