//
// Created by 裴星鑫 on 2022/3/18.
//

#include <string>
#include <vector>

using namespace std;

class Solution {
 public:

  class Tire {
   public:
    Tire() {
      children_ = vector<Tire*>(26, nullptr);
      is_end_ = false;
    }

    void Insert(string& s) {
      auto t = this;
      for (auto c : s) {
        if (t->children_[c - 'a'] == nullptr) {
          t->children_[c - 'a'] = new Tire;
        }
        t = t->children_[c - 'a'];
      }
      t->is_end_ = true;
    }
    bool Prefix(string& s) {
      auto t = this;
      for (auto c : s) {
        if (t->children_[c - 'a'] == nullptr) return false;
        t = t->children_[c - 'a'];
      }
      return true;
    }
   private:
    vector<Tire*> children_;
    bool is_end_;
  };

  void SplitString(string& s, char c, vector<string>& res) {
    int index = 0;
    for (int i = 1; i < s.size(); i++) {
      if (s[i] == c) {
        res.push_back(s.substr(index, i - index));
        index = i + 1;
      }
    }
    if (index < s.size()) res.push_back(s.substr(index));
  }
  string replaceWords(vector<string>& dictionary, string sentence) {
    vector<string> words;
    SplitString(sentence, ' ', words);
    auto cmp = [](string& a, string& b) ->bool {
      return a.length() < b.length();
    };
    sort(dictionary.begin(), dictionary.end(), cmp);

    for (auto& s: words) {
      Tire t;
      t.Insert(s);
      for (auto& d : dictionary) {
        if (t.Prefix(d)) {
          s = d;
          break;
        }
      }
    }

    string res;
    for (int i = 0; i < words.size(); i++) {
      res += words[i];
      if (i < words.size() - 1) res += ' ';
    }
    return res;
  }
};