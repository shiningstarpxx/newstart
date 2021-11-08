//
// Created by 裴星鑫 on 2021/11/8.
//

class Solution {
 public:
  string getHint(string secret, string guess) {
    int bull_count = 0;
    unordered_map<char, int> secret_char_count;
    unordered_map<char, int> guess_char_count;

    for (int i = 0; i < secret.size(); i++) {
      if (secret[i] == guess[i]) {
        bull_count++;
        continue;
      }
      secret_char_count[secret[i]]++;
      guess_char_count[guess[i]]++;
    }

    int cow_count = 0;
    for (auto it = secret_char_count.begin(); it != secret_char_count.end(); it++)
      cow_count += min(it->second, guess_char_count[it->first]);

    return to_string(bull_count) + "A" + to_string(cow_count) + "B";
  }
};