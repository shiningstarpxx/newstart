//
// Created by 裴星鑫 on 2021/10/28.
//

class Solution {
 public:
  string countNum(int n) {
    string s(10, 0);
    while (n) {
      s[n % 10]++;
      n /= 10;
    }
    return s;
  }

  Solution() {
    for (int i = 1; i <= 1e9; i <<= 1) {
      cache_.insert(countNum(i));
    }
  }

  bool reorderedPowerOf2(int n) {
    return cache_.count(countNum(n));
  }

  unordered_set<string> cache_;
};