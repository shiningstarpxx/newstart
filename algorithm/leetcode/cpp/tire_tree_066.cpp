//
// Created by 裴星鑫 on 2022/3/23.
//
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>

using namespace std;

class MapSum {
 public:
  class Tire {
   public:
    Tire() {
      children_ = vector<Tire*> (26, nullptr);
      data_ = vector<vector<string>>(26, vector<string>{});
    }
    void Insert(string& s) {
      auto t = this;
      for (auto c : s) {
        if (t->children_[c - 'a'] == nullptr) t->children_[c - 'a'] = new Tire;
        t = t->children_[c - 'a'];
        t->data_[c - 'a'].push_back(s);
      }
    }

    vector<string> Prefix(string& s) {
      vector<string> res;
      auto t = this;
      for (auto c : s) {
        if (t->children_[c - 'a'] == nullptr) {
          res.clear();
          break;
        };
        t = t->children_[c - 'a'];
        res = t->data_[c - 'a'];
      }
      return res;
    }
   private:
    vector<Tire*> children_;
    vector<vector<string>> data_;
  };

  /** Initialize your data structure here. */
  MapSum() {

  }

  void insert(string key, int val) {
    if (!data_.count(key)) {
      t_.Insert(key);
    }
    data_[key] = val;
  }

  int sum(string prefix) {
    vector<string> keys = t_.Prefix(prefix);
    unordered_set<string> unique_keys(keys.begin(), keys.end());
    int res = 0;
    for (auto& k : unique_keys) {
      res += data_[k];
    }
    return res;
  }

  Tire t_;
  unordered_map<string, int> data_;
};

/**
 * Your MapSum object will be instantiated and called as such:
 * MapSum* obj = new MapSum();
 * obj->insert(key,val);
 * int param_2 = obj->sum(prefix);
 */