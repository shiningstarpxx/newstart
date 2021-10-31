//
// Created by 裴星鑫 on 2021/10/31.
//

class Solution {
 public:
  vector<string> findWords(vector<string>& words) {
    vector<string> res;

    for (auto& s : words) {
      int last_pos = 0;
      int i = 0;
      for (; i < s.size(); i++) {
        if (last_pos == 0 || cache_[tolower(s[i])] == last_pos) {
          last_pos = cache_[tolower(s[i])];
        } else break;
      }
      if (i == s.size()) res.emplace_back(s);
    }

    return res;
  }

  Solution() {
    string s1 = "qwertyuiop";
    insertCache(s1, 1);
    string s2 = "asdfghjkl";
    insertCache(s2, 2);
    string s3 = "zxcvbnm";
    insertCache(s3, 3);
  }

  void insertCache(string& word, int pos) {
    for (auto c : word) {
      cache_[c] = pos;
    }
  }

  unordered_map<char, int> cache_;
};

