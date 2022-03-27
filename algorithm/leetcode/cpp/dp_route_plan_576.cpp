//
// Created by 裴星鑫 on 2022/3/27.
//
#include <vector>

using namespace std;

class Solution {
public:
    int findPaths(int m, int n, int maxMove, int startRow, int startColumn) {
        return ThreeDimensions(m, n, maxMove, startRow, startColumn);
    }

    static const int MOD = 1e9 + 7;

    int ThreeDimensions(int m, int n, int maxMove, int startRow, int startColumn) {
        vector<vector<vector<int>>> dp(m, vector<vector<int>>(n, vector<int>(maxMove + 1)));

        dp[startRow][startColumn][0] = 1;
        int ans = 0;

        int dirs[][2] = {{-1, 0}, {1, 0}, {0, 1}, {0, -1}};

        for (int k = 0; k < maxMove; k++) {
            for (int i = 0; i < m; i++) {
                for (int j = 0; j < n; j++) {
                    if (dp[i][j][k] <= 0) continue;  // for acccelration
                    for (int t = 0; t < 4; t++) {
                        int x = i + dirs[t][0];
                        int y = j + dirs[t][1];
                        if (x >= 0 && x < m && y >= 0 && y < n) {
                            dp[x][y][k + 1] = (dp[x][y][k + 1] + dp[i][j][k]) % MOD;
                        } else {
                            ans = (ans + dp[i][j][k]) % MOD;
                        }
                    }
                }
            }
        }
        return ans;
    }
};