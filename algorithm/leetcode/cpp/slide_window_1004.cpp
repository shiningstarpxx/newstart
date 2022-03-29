//
// Created by 裴星鑫 on 2022/3/29.
//
#include <vector>

using namespace std;

class Solution {
public:
    int longestOnes(vector<int>& nums, int k) {
        int n = nums.size(), sum = 0, ans = 0;
        for (int left = 0, right = 0; right < n; right++) {
            if (nums[right] == 0) sum++;
            while (sum > k) {
                if (nums[left] == 0) sum--;
                left++;
            }
            ans = max(ans, right - left + 1);
        }
        return ans;
    }
};