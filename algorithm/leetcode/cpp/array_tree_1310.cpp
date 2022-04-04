//
// Created by 裴星鑫 on 2022/4/4.
//
#include <vector>

using namespace std;

class Solution {
 public:
  vector<int> xorQueries(vector<int>& arr, vector<vector<int>>& queries) {
    n_ = arr.size() + 1;
    data_ = vector<int> (n_);
    for (int i = 0; i < arr.size(); i++) {
      Update(i, arr[i]);
    }

    vector<int> ans;
    for (auto& v : queries) {
      ans.push_back(Query(v[1]) ^ Query(v[0] - 1));
    }
    return ans;
  }

  int Query(int index) {
    int ans = 0;
    for (int i = index + 1; i > 0; i-=LowBit(i)) {
      ans ^= data_[i];
    }
    return ans;
  }

  void Update(int index, int val) {
    for (int i = index + 1; i < n_; i += LowBit(i)) {
      data_[i] = data_[i] ^ val;
      // cout << data_[i] << endl;
    }
  }

  int LowBit(int x) {
    return x & -x;
  }

  vector<int> data_;
  int n_;
};