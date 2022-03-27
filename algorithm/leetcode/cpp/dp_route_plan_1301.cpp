//
// Created by 裴星鑫 on 2022/3/27.
//

#include <vector>

using namespace std;

class Solution {
 public:
  static const int MOD = 1e9 + 7;
  vector<int> pathsWithMaxScore(vector<string>& board) {
    int n = board.size();
    int m = board[0].size();

    vector<vector<int>> dp(n, vector<int>(m));
    vector<vector<int>> count(n, vector<int>(m));

    for (int i = n - 1; i >= 0; i--) {
      for (int j = m - 1; j >= 0; j--) {
        if (i == n - 1 && j == m - 1) {
          count[i][j] = 1;
          continue;
        }
        if (board[i][j] == 'X') {
          dp[i][j] = INT_MIN;
          continue;
        }

        int val = (i == 0 && j == 0) ? 0 : board[i][j] - '0';
        int score = INT_MIN, plan_number = 0;

        if (i + 1 < n) {
          int current_score = dp[i + 1][j] + val;
          int current_plan = count[i + 1][j];
          Update(current_score, current_plan, score, plan_number);
        }
        if (j + 1 < m) {
          int current_score = dp[i][j + 1] + val;
          int current_plan = count[i][j + 1];
          Update(current_score, current_plan, score, plan_number);
        }
        if (i + 1 < n && j + 1 < m) {
          int current_score = dp[i + 1][j + 1] + val;
          int current_plan = count[i + 1][j + 1];
          Update(current_score, current_plan, score, plan_number);
        }

        dp[i][j] = score < 0 ? INT_MIN : score;
        count[i][j] = plan_number % MOD;
      }
    }

    vector<int> res(2);
    res[0] = dp[0][0] == INT_MIN ? 0 : dp[0][0];
    res[1] = count[0][0] == INT_MIN ? 0 : count[0][0];
    return res;
  }

  void Update(int current_score, int current_plan, int& score, int& plan_number) {
    if (current_score > score) {
      score = current_score;
      plan_number = current_plan;
      return;
    }
    if (current_score == score && current_score != INT_MIN) {
      plan_number += current_plan;
      return;
    }
  }
};