//
// Created by 裴星鑫 on 2022/4/10.
//
#include <vector>
using namespace std;
class Solution {
public:
    int maximalSquare(vector<vector<char>>& matrix) {
        int m = matrix.size();
        int n = matrix[0].size();

        vector<vector<int>> dp(m, vector<int>(n, 0));

        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int d = matrix[i][j] == '0' ? 0 : 1;
                if (d == 0) continue;
                if (i == 0 || j == 0) {
                    dp[i][j] = d;
                } else {
                    dp[i][j] = min(dp[i - 1][j], min(dp[i][j - 1], dp[i - 1][j - 1])) + d;
                }
                ans = max(ans, dp[i][j]);
            }
        }

        return ans * ans;
    }
};