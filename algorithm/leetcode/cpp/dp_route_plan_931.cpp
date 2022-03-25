//
// Created by 裴星鑫 on 2022/3/25.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int minFallingPathSum(vector<vector<int>>& matrix) {
    int r = matrix.size();
    int c = matrix[0].size();

    for (int i = 1; i < r; i++) {
      for (int j = 0; j < c; j++) {
        int from = INT_MAX;
        if (j >= 1 ) from = min(from, matrix[i - 1][j - 1]);
        if (j < c - 1) from = min(from, matrix[i - 1][j + 1]);
        matrix[i][j] += min(from, matrix[i - 1][j]);
      }
    }

    int ans = INT_MAX;
    for (int j = 0; j < c; j++) ans = min(ans, matrix[r - 1][j]);
    return ans;
  }
};