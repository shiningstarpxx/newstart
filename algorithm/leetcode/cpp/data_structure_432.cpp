//
// Created by 裴星鑫 on 2022/3/16.
//
#include <string>
#include <list>
#include <unordered_map>
#include <unordered_set>

using namespace std;

class AllOne {
  list<pair<unordered_set<string>, int>> list_;
  unordered_map<string, list<pair<unordered_set<string>, int>>::iterator> cache_;
 public:
  AllOne() {

  }

  void inc(string key) {
    if (cache_.count(key)) {
      auto it = cache_[key];
      int count = it->second; count++;
      auto nx = next(it);
      if (nx == list_.end() || nx->second > count) {
        cache_[key] = list_.emplace(nx, unordered_set<string>{key}, count);
      } else {
        cache_[key] = nx;
        nx->first.emplace(key);
      }
      it->first.erase(key);
      if (it->first.empty()) list_.erase(it);
    } else {
      if (list_.empty() || list_.begin()->second > 1) {
        list_.emplace_front(unordered_set<string>{key}, 1);
      } else {
        list_.begin()->first.emplace(key);
      }
      cache_[key] = list_.begin();
    }
  }

  void dec(string key) {
    auto it = cache_[key];
    int count = it->second; count--;
    if (count == 0) {
      cache_.erase(key);
    } else {
      auto nx = prev(it);
      if (it == list_.begin() || nx->second < count) {
        cache_[key] = list_.emplace(it, unordered_set<string>{key}, count);;
      } else {
        cache_[key] = nx;
        nx->first.emplace(key);
      }
    }

    it->first.erase(key);
    if (it->first.empty()) list_.erase(it);
  }

  string getMaxKey() {
    return  list_.empty() ? "" : *list_.rbegin()->first.begin();
  }

  string getMinKey() {
    return list_.empty() ? "" : *list_.begin()->first.begin();
  }
};

/**
 * Your AllOne object will be instantiated and called as such:
 * AllOne* obj = new AllOne();
 * obj->inc(key);
 * obj->dec(key);
 * string param_3 = obj->getMaxKey();
 * string param_4 = obj->getMinKey();
 */