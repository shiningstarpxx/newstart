//
// Created by 裴星鑫 on 2022/4/18.
//
#include <math.h>
#include <string>
#include <vector>

using namespace std;

// 是数位DP，这里最难的部分是比如  12345, 计算的时候实际是 1->1, 2->2 3->3, 4->4, 5->5 , 然后dp也这么算
class Solution {
 public:
  int atMostNGivenDigitSet(vector<string>& digits, int n) {
    string s = to_string(n);
    int k = s.size();
    vector<int> dp(k + 1);
    dp[k] = 1;

    // 情况1: 位数相等，有多少个？
    for (int i = k - 1; i >= 0; i--) {
      int sd = s[i] - '0';
      for (auto &d : digits) {
        int t = d[0] - '0';
        if (t < sd) {
          dp[i] += pow(digits.size(), k - i -1);
        } else if (t == sd)
          dp[i] += dp[i + 1];
      }
    }

    // 情况2: 位数少于n
    for (int i = 1; i < k; i++) {
      dp[0] += pow(digits.size(), i);
    }
    return dp[0];
  }
};