//
// Created by 裴星鑫 on 2022/3/20.
//
#include <queue>
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  class Tire {
   public:
    Tire() {
      children_ = vector<Tire*>(26, nullptr);
    }
    void Insert(string& s) {
      auto t = this;
      for (auto c : s) {
        if (t->children_[c - 'a'] == nullptr) t->children_[c - 'a'] = new Tire;
        t = t->children_[c - 'a'];
        t->words_.push(s);
        if (t->words_.size() > 3) t->words_.pop();
      }
    }
    priority_queue<string> Prefix(string& s) { // 这里效率不高，可以直接在下面实现，缺点是成员变量要public了
      auto t = this;
      for (auto c : s) {
        if (t->children_[c - 'a'] == nullptr) return {};
        t = t->children_[c - 'a'];
      }
      return t->words_;
    }
   private:
    vector<Tire*> children_;
    priority_queue<string> words_;
  };

  vector<vector<string>> suggestedProducts(vector<string>& products, string searchWord) {
    Tire t;
    for (auto& s : products) t.Insert(s);

    vector<vector<string>> res;
    string prefix = "";
    for (auto c : searchWord) {
      prefix += c;
      vector<string> tmp;
      auto q = t.Prefix(prefix);
      if (!q.empty()) {
        tmp.clear();
        while (!q.empty()) {
          tmp.push_back(q.top());
          q.pop();
        }
        reverse(tmp.begin(), tmp.end());
        res.push_back(tmp);
      } else {
        res.push_back(vector<string>{});
      }

    }
    return res;
  }
};