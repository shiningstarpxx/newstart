//
// Created by 裴星鑫 on 2022/3/27.
//
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  string longestPalindrome(string s) {
    string str = ManachStr(s);
    vector<int> radius(str.size());
    int r = 0, c = 0;

    int max_length = 0;
    int max_index = 0;
    // L  i/  C  i  R
    for (int i = 0; i < str.size(); i++) {
      int mirror_index = 2 * c - i;
      if (i < r) radius[i] = min(r - i, radius[mirror_index]);
      int right = i + (1 + radius[i]);
      int left = i - (1 + radius[i]);
      while (left >= 0 && right < str.size() && str[left] == str[right]) {
        radius[i]++;
        left--;
        right++;
      }
      if (i + radius[i] > r) {
        r = i + radius[i];
        c = i;
        if (radius[i] > max_length) {
          max_index = c;
          max_length = radius[i];
        }
      }
    }

    // cout << max_length << "," << max_index << endl;
    string res;
    for (int i = max_index - max_length; i <= max_index + max_length; i++) {
      if (str[i] != '#') res += str[i];
    }
    return res;
  }

  string ManachStr(string& s) {
    string res = "#";
    for (int i = 0; i < s.size(); i++) {
      res = res + s[i] + '#';
    }
    return res;
  }
};