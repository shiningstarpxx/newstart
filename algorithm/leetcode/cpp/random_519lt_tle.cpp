//
// Created by Pei Xingxin on 27/11/2021.
//

class Solution {
 public:
  Solution(int m, int n) {
    init(m, n);
  }

  vector<int> flip() {
    if (index_.empty()) return {};

    int index = rand() % index_.size();
    int val = index_[index];
    int row = val / n_;
    int column = val % n_;

    index_[index] = index_[index_.size() - 1];
    index_.pop_back();

    return {row, column};
  }

  void reset() {
    index_ = origin_;
  }

 private:
  void init(int m, int n) {
    m_ = m;
    n_ = n;
    index_.resize(m * n);
    for (int i = 0; i < m; i++) {
      for (int j = 0; j < n; j++)
        index_[i * n + j] = i * n + j;
    }
    origin_ = index_;
  }

  int m_;
  int n_;
  vector<int> index_;
  vector<int> origin_;
};

/**
 * Your Solution object will be instantiated and called as such:
 * Solution* obj = new Solution(m, n);
 * vector<int> param_1 = obj->flip();
 * obj->reset();
 */