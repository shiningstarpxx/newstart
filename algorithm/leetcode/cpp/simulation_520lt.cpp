//
// Created by 裴星鑫 on 2021/11/13.
//

#include <string>

using namespace std;

class Solution {
 public:
  bool detectCapitalUse(string word) {
    int lower_count = 0;
    int upper_count = 0;
    int first_count = 0;

    for (int i = 0; i < word.size(); i++) {
      if (word[i] >= 'A' && word[i] <= 'Z') {
        upper_count++;
        if (i == 0) first_count++;
      } else lower_count++;
    }

    return upper_count == word.size() || upper_count == first_count;
  }
};
