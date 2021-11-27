//
// Created by 裴星鑫 on 2021/11/27.
//
#include <random>
#include <unordered_map>
#include <vector>

using namespace std;

class RandomizedSet {
 public:
  RandomizedSet() {

  }

  bool insert(int val) {
    if (cache_.count(val)) return false;
    cache_[val] = data_.size();
    data_.push_back(val);
    return true;
  }

  bool remove(int val) {
    if (!cache_.count(val)) return false;
    int del_index = cache_[val];
    data_[del_index] = data_[data_.size() - 1];
    cache_[data_[data_.size() - 1]] = del_index;
    cache_.erase(val);
    data_.pop_back();
    return true;
  }

  int getRandom() {
    return data_[rand() % data_.size()];
  }

  vector<int> data_;
  unordered_map<int, int> cache_;
};

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * RandomizedSet* obj = new RandomizedSet();
 * bool param_1 = obj->insert(val);
 * bool param_2 = obj->remove(val);
 * int param_3 = obj->getRandom();
 */

