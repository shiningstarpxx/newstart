//
// Created by 裴星鑫 on 2022/4/4.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int trap(vector<int>& height) {
    int n = height.size();
    vector<int> left(n);
    vector<int> right(n);

    int c_max = height[0];
    left[0] = c_max;
    for (int i = 1; i < n; i++) {
      c_max = max(height[i], c_max);
      left[i] = c_max;
    }

    c_max = height[n - 1];
    right[n - 1] = c_max;
    for (int i = n - 2; i >= 0; i--) {
      c_max = max(height[i], c_max);
      right[i] = c_max;
    }

    int ans = 0;
    for (int i = 0; i < n; i++) {
      ans += min(left[i], right[i]) - height[i];
    }
    return ans;
  }
};