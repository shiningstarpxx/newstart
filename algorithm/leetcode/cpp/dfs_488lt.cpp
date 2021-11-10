//
// Created by 裴星鑫 on 2021/11/9.
//

#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>

using namespace std;
#include "gtest/gtest.h"

class Solution {
 public:
  int findMinStep(string board, string hand) {
    int n = hand.length();
    hand_ = hand;
    int ans = dfs(board, 1 << n);
    return ans < hand.length() ? ans : -1;
  }

  int dfs(string target, int current) {
    if (target.length() == 0) return 0;
    if (cache_.count(target)) return cache_[target];
    int count = INT_MAX;
    for (int i = 0; i < target.length(); i++) {
      if (i < target.length() - 1 && target[i + 1] == target[i]) continue;
      for (int j = 0; j < hand_.length(); j++) {
        if (hand_[j] != target[i] || (current & (1 << j)) != 0) continue;
        string tmp = target.substr(0, i) + hand_[j] + target.substr(i);
        tmp = compress(tmp);
        count = min(count, dfs(tmp, current | (1 << j)) + 1);
      }
    }
    cache_[target] = count;
    return count;
  }

  string compress(string s) {
    bool skip;
    do {
      skip = false;
      string res;
      for (int i = 0, j; i < s.length(); i = j) {
        for (j = i + 1; j < s.length() && s[j] == s[i]; j++) ;
        if (j - i < 3) {
          res += s.substr(i, j-i);
        }
        else skip = true;
      }
      s = res;
    } while(skip);
    return s;
  }

  string hand_;
  unordered_map<string, int> cache_;
};

TEST(Solution, compress) {
  Solution s;
  EXPECT_EQ("ww", s.compress("wwbbb"));
  EXPECT_EQ("wwxxww", s.compress("wwyyybbbxxww"));
  EXPECT_EQ("wwbbxxww", s.compress("wwyyybbxxww"));
}

TEST(Solution, findMinStep) {
  Solution s;
  EXPECT_EQ(2, s.findMinStep("WWRRBBWW", "WRBRW"));
  EXPECT_EQ(-1, s.findMinStep("WRRBBW", "RB"));
}