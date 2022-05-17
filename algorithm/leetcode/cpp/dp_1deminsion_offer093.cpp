//
// Created by 裴星鑫 on 2022/5/17.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int lenLongestFibSubseq(vector<int>& arr) {
    int n = arr.size();
    vector<vector<int>> dp(n, vector<int>(n));
    int ans = 0;
    for (int i = 0; i < n; i++) {
      for (int j = 0; j < i; j++) {
        int index = GetIndex(arr, j-1, arr[i] - arr[j]);
        dp[i][j] = index == - 1? 2 : dp[j][index] + 1;
        ans = max(ans, dp[i][j]);
      }
    }
    return ans > 2 ? ans : 0;
  }

  int GetIndex(vector<int>& arr, int right, int target) {
    int l = 0, r = right;
    while (l <= r) {
      int mid = l + (r - l) / 2;
      if (target == arr[mid]) return mid;
      else if (target > arr[mid]) l = mid + 1;
      else r = mid - 1;
    }
    return -1;
  }
};