//
// Created by 裴星鑫 on 2022/3/29.
//
#include <string>

using namespace std;

class Solution {
 public:
  int Match(string& answerKey, int k, char t) {
    int sum = 0;
    int ans = 0;
    for (int left = 0, right = 0; right < answerKey.size(); right++) {
      if (answerKey[right] != t) sum++;
      while (sum > k && left < right) {
        if (answerKey[left] != t) sum--;
        left++;
      }
      ans = max(ans, right - left + 1);
    }
    return ans;
  }

  int maxConsecutiveAnswers(string answerKey, int k) {
    return max(Match(answerKey, k, 'T'), Match(answerKey, k, 'F'));
  }
};
