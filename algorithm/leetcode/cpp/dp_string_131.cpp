//
// Created by 裴星鑫 on 2022/4/19.
//
#include <string>
#include <vector>

using namespace std;

class Solution {
public:
    vector<vector<string>> partition(string s) {
        int n = s.size();
        vector<vector<bool>> dp(n, vector<bool>(n));

        vector<vector<string>> res;

        // WA
        // for (int i = 0; i < n; i++) {
        //     for (int j = i; j >= 0; j--) {
        //         if (i == j) dp[i][j] = true;
        //         else if (i - j == 1) dp[i][j] = s[i] == s[j];
        //         else if (i - j >= 2) dp[i][j] = s[i] == s[j] && dp[i-1][j+1];
        //     }
        // }

        // ac
        // for (int j = 0; j < n; j++) {
        //     for (int i = j; i >= 0; i--) {
        //         if (i == j) dp[i][j] = true;
        //         else if (j - i == 1) dp[i][j] = s[i] == s[j];
        //         else if (j - i >= 2) dp[i][j] = s[i] == s[j] && dp[i + 1][j - 1];
        //     }
        // }

        for (int i = n - 1; i >= 0; i--) {
            for (int j = i; j < n; j++) {
                if (i == j) dp[i][j] = true;
                else if (j - i == 1) dp[i][j] = s[i] == s[j];
                else if (j - i >= 2) dp[i][j] = s[i] == s[j] && dp[i + 1][j - 1];
            }
        }

        vector<string> cur;
        dfs(res, 0, s, dp, cur);
        return res;
    }

    void dfs(vector<vector<string>>& res, int index, string& s, vector<vector<bool>>& dp, vector<string>& cur) {
        if (index >= s.size()) {
            if (!cur.empty()) res.push_back(cur);
            return;
        }

        for (int i = index; i < s.size(); i++) {
            if (dp[index][i]) {
                cur.push_back(s.substr(index, i - index + 1));
                dfs(res, i + 1, s, dp, cur);
                cur.pop_back();
            }
        }
    }
};