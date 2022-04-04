//
// Created by 裴星鑫 on 2022/4/4.
//
#include <unordered_map>
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  bool wordBreak(string s, vector<string>& wordDict) {
    int n = s.size() + 1;
    vector<bool> dp(n);
    dp[0] = true;

    unordered_map<string, bool> cache;
    for (auto& s : wordDict) cache[s] = true;

    for (int i = 1; i < n; i++) {
      for (int j = 0; j < i; j++) {
        if (dp[j] && cache.count(s.substr(j, i - j))) {
          dp[i] = true;
          break;
        }
      }
    }
    return dp.back();
  }
};