//
// Created by 裴星鑫 on 2021/11/27.
//

#include <random>
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
 public:
  Solution(int m, int n) {
    m_ = m;
    n_ = n;
    total_ = m * n;
    srand(unsigned(time(nullptr)));
  }

  vector<int> flip() {
    int index = rand() % total_;
    total_--;
    vector<int> ans;
    if (swap_index_.count(index)) {
      ans = {swap_index_[index] / n_, swap_index_[index] % n_};
    } else {
      ans = {index / n_, index % n_};
    }

    if (swap_index_.count(total_)) {
      swap_index_[index] = swap_index_[total_];
    } else {
      swap_index_[index] = total_;
    }
    return ans;
  }

  void reset() {
    total_ = m_ * n_;
    swap_index_.clear();
  }

 private:
  int m_;
  int n_;
  int total_;
  unordered_map<int, int> swap_index_;
};

/**
 * Your Solution object will be instantiated and called as such:
 * Solution* obj = new Solution(m, n);
 * vector<int> param_1 = obj->flip();
 * obj->reset();
 */