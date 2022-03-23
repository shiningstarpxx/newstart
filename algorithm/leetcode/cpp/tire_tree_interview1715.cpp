//
// Created by 裴星鑫 on 2022/3/23.
//
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:

  class Tire {
   public:
    Tire() {
      children_ = vector<Tire*> (26, nullptr);
      is_end_ = false;
    }
    void Insert(string& s) {
      auto t = this;
      for (auto c : s) {
        if (t->children_[c - 'a'] == nullptr) t->children_[c - 'a'] = new Tire;
        t = t->children_[c - 'a'];
      }
      t->is_end_ = true;
    }
    bool Match(string s) {
      auto t = this;
      return MatchHelper(s, 0, t);
    }
    bool MatchHelper(string s, int index, Tire* t) {
      if (index >= s.size()) {
        return t->is_end_;
      }

      if (t->children_[s[index] - 'a'] == nullptr) {
        return false;
      }

      t = t->children_[s[index] - 'a'];
      bool flag = false;

      if (t->is_end_ && index + 1 < s.size()) {
        Tire* tmp = this;
        flag = MatchHelper(s.substr(index + 1), 0, tmp);
      }
      flag |= MatchHelper(s, index + 1, t);
      return flag;
    }

   private:
    vector<Tire*> children_;
    bool is_end_;
  };

  string longestWord(vector<string>& words) {
    auto cmp = [](string& a, string& b) {
      return a.size() < b.size();
    };
    sort(words.begin(), words.end(), cmp);

    Tire t;
    int count = 0;
    string res;
    for (auto& s : words) {
      if (t.Match(s)) {
        if (s.size() > count) {
          res = s;
          count = s.size();
        } else if (s.size() == count) res = res < s ? res : s;
      } else {
        t.Insert(s);
      }
    }
    return res;
  }
};
