//
// Created by 裴星鑫 on 2021/11/17.
//
#include <cmath>
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  int maxProduct(vector<string>& words) {
    vector<int> cache;
    for (int i = 0; i < words.size(); i++) {
      int d = 0;
      for (int j = 0; j < words[i].size(); j++) {
        int swift = words[i][j] - 'a';
        d |= 1 << swift;
      }
      cache.push_back(d);
    }

    int ans = 0;
    for (int i = 0; i < words.size(); i++) {
      for (int j = i + 1; j < words.size(); j++) {
        if ((cache[i] & cache[j]) == 0) ans = max(ans, int(words[i].size() * words[j].size()));
      }
    }
    return ans;
  }
};
