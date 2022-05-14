//
// Created by 裴星鑫 on 2022/5/14.
//
#include <string>
#include <vector>

using namespace std;

// 深度搜索，可以解决很多问题，这里深度搜索+memorization实现了最优化解的问题
class Solution {
public:
    int minStickers(vector<string>& stickers, string target) {
        int m = target.size();
        vector<int> dp(1 << m, -1);
        dp[0] = 0;
        function<int(int)> dfs = [&](int state) {
            if (dp[state] != -1) return dp[state];
            // 0 的情况前面定义了
            dp[state] = m + 1;
            for (auto &s : stickers) {
                int cnt[26] = {0};  // 数组的初始化可能是random
                int next_state = state;
                for (auto c : s) {
                    cnt[c - 'a']++;
                }
                for (int i = 0; i < m; i++) {
                    if ((next_state >> i & 1) && cnt[target[i] - 'a'] > 0) {
                        cnt[target[i] - 'a']--;
                        next_state ^= 1 << i;
                    }
                }
                if (next_state < state) {
                    dp[state] = min(dp[state], dfs(next_state) + 1);
                }
            }
            return dp[state];
        };

        int res = dfs((1<<m) - 1);
        return res <= m ? res : -1;
    }
};