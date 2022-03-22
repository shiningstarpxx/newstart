//
// Created by 裴星鑫 on 2022/3/22.
//
#include <string>

using namespace std;

class Solution {
 public:
  bool winnerOfGame(string colors) {
    int left = 0, right = 0;
    int a_count = 0;
    while (left < colors.size()) {
      if (colors[left] == 'A') {
        right = left + 1;
        while (colors[right] == 'A' && right < colors.size()) {
          right++;
        }
        a_count += right - left - 2 > 0 ? right - left - 2 : 0;
        left = right;
      } else left++;
    }
    left = 0, right = 0;
    int b_count = 0;
    while (left < colors.size()) {
      if (colors[left] == 'B') {
        right = left + 1;
        while (colors[right] == 'B' && right < colors.size()) {
          right++;
        }
        b_count += right - left - 2 > 0 ? right - left - 2 : 0;
        left = right;
      } else left++;
    }
    return a_count > b_count;
  }
};
