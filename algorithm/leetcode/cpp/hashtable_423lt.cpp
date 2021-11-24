//
// Created by 裴星鑫 on 2021/11/24.
//
#include <string>
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
 public:
  string originalDigits(string s) {
    unordered_map<char, int> cache;
    for (auto c : s) cache[c]++;

    vector<int> count(10, 0);
    count[0] = cache['z'];
    count[2] = cache['w'];
    count[4] = cache['u'];
    count[6] = cache['x'];
    count[8] = cache['g'];

    count[3] = cache['h'] - count[8];
    count[5] = cache['f'] - count[4];
    count[7] = cache['s'] - count[6];

    count[1] = cache['o'] - count[0] - count[2] - count[4];
    count[9] = cache['i'] - count[5] - count[6] - count[8];

    string ans;
    for (int i = 0; i <= 9; i++) {
      for (int j = 0; j < count[i]; j++) {
        ans += (i + '0');
      }
    }
    return ans;
  }
};