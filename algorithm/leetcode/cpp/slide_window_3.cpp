//
// Created by 裴星鑫 on 2022/3/30.
//
#include <string>
#include <unordered_map>

using namespace std;

class Solution {
 public:
  int lengthOfLongestSubstring(string s) {
    unordered_map<char, int> charset;
    int ans = 0;
    for (int left = 0, right = 0; right < s.size(); right++) {
      charset[s[right]]++;
      while (charset[s[right]] > 1) {
        charset[s[left]]--;
        left++;
      }
      // cout << right << "," << left << endl;
      ans = max(ans, right - left + 1);
    }
    return ans;
  }
};