//
// Created by 裴星鑫 on 2022/3/21.
//
#include <ctype.h>
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<bool> camelMatch(vector<string>& queries, string pattern) {
//        return BruteFroce(queries, pattern);
    return TireMatch(queries, pattern);
  }

  vector<bool> BruteFroce(vector<string>& queries, string pattern) {
    vector<bool> res;
    for (auto& s : queries) {
      int index = 0;
      int i = 0;
      for (; i < s.size(); i++) {
        if (islower(s[i])) {
          if (s[i] == pattern[index]) index++;
          else continue;
        } else {
          if (s[i] == pattern[index]) index++;
          else break;
        }
      }
      res.push_back(i == s.size() && index == pattern.size());
    }
    return res;
  }

  class Tire {
   public:
    Tire() {
      children_ = vector<Tire* > (52, nullptr);
      is_end_ = false;
    }
    void Insert(string& s) {
      auto t = this;
      for (auto c : s) {
        if (islower(c)) {
          if (t->children_[c - 'a'] == nullptr) t->children_[c - 'a'] = new Tire;
          t = t->children_[c - 'a'];
        } else {
          if (t->children_[c - 'A' + 26] == nullptr) t->children_[c - 'A' + 26] = new Tire;
          t = t->children_[c - 'A' + 26];
        }
      }
      t->is_end_ = true;
    }
    bool Match(string& s) {
      auto t = this;
      for (auto c : s) {
        if (islower(c)) {
          if (t->children_[c - 'a'] == nullptr) continue;
          t = t->children_[c - 'a'];
        } else {
          if (t->children_[c - 'A' + 26] == nullptr) return false;
          t = t->children_[c - 'A' + 26];
        }
      }
      return t->is_end_;
    }
   private:
    vector<Tire*> children_;
    bool is_end_;
  };

  vector<bool> TireMatch(vector<string>& queries, string pattern) {
    Tire t;
    t.Insert(pattern);
    vector<bool>res;
    for (auto& s : queries) {
      if (t.Match(s)) res.push_back(true);
      else res.push_back(false);
    }
    return res;
  }
};