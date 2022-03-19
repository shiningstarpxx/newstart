//
// Created by 裴星鑫 on 2022/3/19.
//
#include <string>
#include <unordered_set>
#include <vector>

using namespace std;

class Solution {
 public:
  int minimumLengthEncoding(vector<string>& words) {
    // return SubstrMatch(words);  // LTE
    // return TireHelper(words);
    return DProcess(words);
  }

  int DProcess(vector<string>& words) {
    unordered_set<string> unique_words(words.begin(), words.end());
    for (auto& s : words) {
      for (int i = 1; i < s.size(); i++) {
        unique_words.erase(s.substr(i));
      }
    }
    int res = 0;
    for (auto& v : unique_words) res += v.size() + 1;
    return res;
  }

  int SubstrMatch(vector<string>& words) {  // LTE
    unordered_set<string> unique_words(words.begin(), words.end());

    int res = 0;
    for (auto it = unique_words.begin(); it != unique_words.end(); it++) {
      bool match = false;
      for (auto& s : unique_words) {
        if (s == *it) continue;
        int index = 0, pos = 0;
        while ((index = s.find(*it, pos)) != string::npos && s.substr(index) != *it) {
          pos = index + 1;
        }
        if (index != string::npos && s.substr(index) == *it) {
          match = true;
          break;
        }
      }
      if (!match) res += (*it).size() + 1;
    }
    return res;
  }


  class Tire {
   public:
    Tire() {
      children_ = vector<Tire*> (26, nullptr);
      count_ = 0;
    }

    int Insert(string& s) {
      auto t = this;
      bool new_str = false;
      for (int i = 0; i < s.length(); i++) {
        if (t->children_[s[s.length() - 1 - i] - 'a'] == nullptr) {
          t->children_[s[s.length() - 1 - i] - 'a'] = new Tire;
          new_str = true;
        }
        t = t->children_[s[s.length() - 1 - i] - 'a'];
      }
      return new_str ? s.size() + 1 : 0;
    }

   private:
    vector<Tire*> children_;
    int count_;
  };

  int TireHelper(vector<string>& words) {
    Tire t;
    auto cmp = [](string& a, string& b) -> bool {
      return a.size() > b.size();
    };
    sort(words.begin(), words.end(), cmp);

    int res = 0;
    for (auto& s : words) {
      res += t.Insert(s);
    }
    return res;
  }

};