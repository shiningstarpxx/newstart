//
// Created by 裴星鑫 on 2022/5/14.
//
#include <vector>

using namespace std;

// 跟213， 打家劫舍2的思路类似，本质上都是对 "i"这个格子，取或者不取，带来解的迭代前行
// 唯一的不同是，取格子的个数的上界，213中没有限制-也就是最多可以取n/2，而1388限制为 n/3
// 列表循环+相领 问题下的通行解法： max(calc(0, n-2), calc(1, n-1))
class Solution {
public:
    int maxSizeSlices(vector<int>& slices) {
        n_ = slices.size();
        count_ =  n_ / 3;

        return max(Calc(slices, 0, n_-2), Calc(slices, 1, n_-1));
    }

    int Calc(vector<int>& slices, int begin, int end) {
        vector<vector<int>> dp(n_, vector<int>(count_+1));
        dp[begin][1] = slices[begin];
        for (int i = begin + 1; i <= end; i++) {
            for (int j = 1; j <= count_; j++) {
                // 第i个元素不取, 很显然
                // 第i个元素取，意味着 i-1 也不能取，那么需要在前 i-2个元素里取 j-1个
                dp[i][j] = max(dp[i - 1][j], i-2 >= 0? dp[i-2][j-1] + slices[i] : 0);
            }
        }
        return dp[end].back();
    }

    int n_, count_;
};