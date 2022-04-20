//
// Created by 裴星鑫 on 2022/4/20.
//
#include <vector>

#include "gtest/gtest.h"

using namespace std;

class Solution {
 public:
  int maxValue(int c, vector<int>& values, vector<int>& weights) {

  }

  int wholeArray(int c, vector<int>& values, vector<int>& weights) {
    int n = values.size();

    vector<vector<int>> dp(n, vector<int>(c+1));
    for (int i = 0; i <= c; i++) {
      dp[0][i] = i >= weights[0] ? values[0] : 0;
    }

    for (int i = 1; i < n; i++) {
      for (int j = 0; j <= c; j++) {
        int a = dp[i - 1][j];
        int b = j >= weights[i] ? dp[i - 1][j - weights[i]] + values[i] : 0;
        dp[i][j] = max(a, b);
      }
    }

    return dp[n - 1][c];
  }

  int twoRowArray(int c, vector<int>& values, vector<int>& weights) {
    int n = values.size();
    vector<vector<int>> dp(2, vector<int>(c + 1));
    for (int i = 0; i <= c; i++) dp[0][i] = i >= weights[0] ? values[0] : 0;

    for (int i = 1; i < n; i++) {
      for (int j = 0; j <= c; j++) {
        int a = dp[(i - 1) & 1][j];
        int b = j >= weights[i] ? dp[(i - 1) & 1][j - weights[i]] + values[i] : 0;
        dp[i & 1][j] = max(a, b);
      }
    }

    return dp[(n - 1) & 1][c];
  }

  int oneRowArray(int c, vector<int>& values, vector<int>& weights) {
    int n = values.size();
    vector<int> dp(c + 1);
    for (int i = 0; i <= c; i++) dp[i] = i >= weights[0] ? values[0] : 0;

    for (int i = 1; i < n; i++) {
      for (int j = c; j >= weights[i]; j--) {
        int a = dp[j];
        int b = j >= weights[i] ? dp[j - weights[i]] + values[i] : 0;
        dp[j] = max(a, b);
      }
    }

    return dp[c];
  }
};