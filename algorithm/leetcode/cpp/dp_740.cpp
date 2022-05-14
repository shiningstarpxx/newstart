//
// Created by 裴星鑫 on 2022/5/14.
//
#include <vector>

using namespace std;

// 通过变形把一个未知的问题，转化成已知的问题
// 离散的数字变成了一条线上的邻居，变成了相邻取一个的问题上
class Solution {
public:
    int deleteAndEarn(vector<int>& nums) {
        if (nums.size() == 1) return nums[0];
        int max_value = *max_element(nums.begin(), nums.end());
        vector<int> data(max_value + 1);

        for (auto v : nums) data[v] += v;

        int first = data[1], second = max(data[1], data[2]);
        for (int i = 3; i <= max_value; i++) {
            int t = second;
            second = max(first + data[i], second);
            first = t;
        }
        return max(second, first);
    }
};