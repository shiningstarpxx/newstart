//
// Created by 裴星鑫 on 2022/4/4.
//

// Fenwick tree, also binary indexed tree
#include <iostream>
#include <vector>

using namespace std;

class BinaryIndexedTree {
 public:
  BinaryIndexedTree(vector<int>& d) {
    n_ = d.size();
    data_ = vector<int>(n_ + 1);
    for (int i = 0; i < n_; i++) {
      UpdateWithDelta(i, d[i]);
    }
  }
  int Query(int index) { // Query index, if query [left, right], Query(right) - Query(left - 1);
    int ans = 0;
    for (int i = index + 1; i > 0; i -= LowBit(i)) {
      ans += data_[i];
    }
    return ans;
  }
  void UpdateWithDelta(int index, int val) {
    index = index + 1;
    for (int i = index; i <= n_; i += LowBit(i)) data_[i] += val;
  }
 private:
  int LowBit(int x) {
   return x & (-x);
 }
  vector<int> data_;
  int n_;
};

int main() {
  vector<int> d= {1, 3, 5};
  BinaryIndexedTree bt(d);
  cout << bt.Query(2) << endl;
  return 0;
}