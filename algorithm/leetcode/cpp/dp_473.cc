//
// Created by 裴星鑫 on 2022/6/1.
//

#include <vector>
#include <numeric>

using namespace std;

class Solution {  // 状态压缩
public:
    bool makesquare(vector<int>& matchsticks) {
        int sum = accumulate(matchsticks.begin(), matchsticks.end(), 0);
        if (sum % 4 != 0) return false;
        m_ = matchsticks;
        average_ = sum / 4;
        int n = matchsticks.size();
        size_ = n;
        dp_ = vector<int>(1<<n + 1, -1);
        return dfs(average_, 4, 0);
    }

    vector<int> dp_;
    int average_;
    vector<int> m_;
    int size_;

    bool dfs(int left, int count, int mask) {
        if (dp_[mask] != -1) return dp_[mask] == 1;

        if (left < 0) {dp_[mask] = 0; return false;}
        if (left == 0) {
            count--;
            left = average_;
            if (count == 0) {dp_[mask] = 1; return true;}
        }
        for (int i = 0; i < size_; i++) {
            int p = mask & (1 << i);
            if (p == 0) {
                bool r = dfs(left - m_[i], count, mask | (1 << i));
                if (r) {dp_[mask] = 1; return true;}
            }
        }
        dp_[mask] = 0;
        return false;
    }
};