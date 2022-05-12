//
// Created by 裴星鑫 on 2022/5/12.
//
#include <numeric>
#include <vector>
using namespace std;

// 最核心的思想，是反过来想
// 求最大就可以变成求最小问题
class Solution {
public:
    int maxSubarraySumCircular(vector<int>& nums) {
        int sum = accumulate(nums.begin(), nums.end(), 0);

        int max_ans = nums[0];
        int current = 0;
        for (auto v : nums) {
            current = max(current + v, v);
            max_ans = max(max_ans, current);
        }

        int min_ans = nums[0];
        current = 0;
        for (auto v : nums) {
            current = min(current + v, v);
            min_ans = min(min_ans, current);
        }

        // cout << max_ans << "," << min_ans << endl;
        return sum - min_ans == 0 ? max_ans : max(max_ans, sum - min_ans);
    }
};