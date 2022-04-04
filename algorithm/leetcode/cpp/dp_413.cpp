//
// Created by 裴星鑫 on 2022/4/4.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int numberOfArithmeticSlices(vector<int>& nums) {
    int n = nums.size();

    int ans = 0;
    for (int i = 2; i < nums.size(); i++) {
      int diff = nums[i] - nums[i - 1];
      for (int j = i - 1; j >= 0; j--) {
        if (j - 1 >= 0 && diff == nums[j] - nums[j - 1]) ans++;
        else break;
      }
    }
    return ans;
  }
};