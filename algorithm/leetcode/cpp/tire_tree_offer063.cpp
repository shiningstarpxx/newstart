//
// Created by 裴星鑫 on 2022/3/22.
//
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<string> SplitString(string& s) {
    vector<string> res;
    int left = 0;
    for (int i = 0; i < s.size(); i++) {
      if (s[i] == ' ') {
        res.push_back(s.substr(left, i - left));
        left = i + 1;
      }
    }
    if (left < s.size()) res.push_back(s.substr(left));
    return res;
  }

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
    bool FetchPrefix(string& s, string& res) {
      auto t = this;
      for (auto c : s) {
        if (t->children_[c - 'a'] == nullptr) return false;
        res += c;
        t = t->children_[c - 'a'];
        if (t->is_end_) return true;
      }
      return false;
    }
   private:
    vector<Tire*> children_;
    bool is_end_;
  };

  string replaceWords(vector<string>& dictionary, string sentence) {
    vector<string> v = SplitString(sentence);
    Tire t;
    for (auto& s : dictionary) t.Insert(s);
    for (auto& s : v) {
      string tmp = "";
      if (t.FetchPrefix(s, tmp)) s = tmp;
    }
    string res = "";
    for (int i = 0; i < v.size(); i++) {
      res += v[i];
      if (i < v.size() - 1) res += " ";
    }
    return res;
  }
};
