//
// Created by 裴星鑫 on 2021/11/14.
//

#include <string>
#include <unordered_map>
using namespace std;

struct TireNode677 {
  int val;
  TireNode677* next[26];
  TireNode677() {
    val = 0;
    for (int i = 0; i < 26; i++) {
      next[i] = nullptr;
    }
  }
};

class MapSum {
 public:
  /** Initialize your data structure here. */
  MapSum() {
    root_ = new TireNode677();
  }

  void insert(string key, int val) {
    int delta = val;
    if (cache_.count(key)) delta -= cache_[key];

    TireNode677* t = root_;
    for (auto c : key) {
      if (t->next[c - 'a'] == nullptr) {
        t->next[c - 'a'] = new TireNode677();
      }
      t = t->next[c - 'a'];
      t->val += delta;
    }
    cache_[key] = val;
  }

  int sum(string prefix) {
    TireNode677* t = root_;
    for (auto c : prefix) {
      if (t->next[c - 'a'] == nullptr) return 0;
      t = t->next[c - 'a'];
    }
    return t->val;
  }

  int naviesSum(string prefix) {
    int sum = 0;
    for (auto it = cache_.begin(); it != cache_.end(); it++) {
      if (it->first.substr(0, prefix.size()) == prefix) sum += it->second;
    }
    return sum;
  }

  unordered_map<string, int> cache_;

  TireNode677* root_;
};

/**
 * Your MapSum object will be instantiated and called as such:
 * MapSum* obj = new MapSum();
 * obj->insert(key,val);
 * int param_2 = obj->sum(prefix);
 */
