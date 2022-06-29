//
// Created by 裴星鑫 on 2022/6/7.
//

#include <vector>

using namespace std;

// 效率非常低，但是适用范围非常广
// 单点更新，区间更新 -》 区间求和，单点求值
class NumArray {
 public:
  struct Node {
    int l, r;
    int val;
    Node(int ll, int rr) {
      l = ll;
      r = rr;
      val = 0;
    }
    Node() {l=0, r = 0; val = 0;}
  };

  void Build(vector<int>& nums) {
    int n = nums.size();
    data_ = vector<Node> (4 * n);
    DoBuild(1, 1, n);
  }

  vector<Node> data_;

  void DoBuild(int k, int l, int r) {
    data_[k].l = l;
    data_[k].r = r;
    if (l == r) return;
    int mid = l + ((r - l) >> 1);
    DoBuild(k << 1, l, mid);
    DoBuild(k << 1 | 1, mid + 1, r);
  }

  void PushUp(int k) {
    data_[k].val = data_[k << 1].val + data_[k << 1 | 1].val;
  }
  void Update(int k, int index, int val) {
    if (data_[k].l == index && data_[k].r == index) {
      data_[k].val += val;
      return;
    }
    int mid = (data_[k].l + data_[k].r) >> 1;
    if (mid >= index) Update(k << 1, index, val);
    else Update(k << 1 | 1, index, val);
    PushUp(k);
  }

  int Query(int k, int l, int r) {
    if (data_[k].l >= l && data_[k].r <= r) return data_[k].val;
    int ans = 0;
    int mid = (data_[k].l + data_[k].r) >> 1;
    if (mid >= l) ans += Query(k << 1, l, r);
    if (mid < r) ans += Query(k << 1 | 1, l, r);
    return ans;
  }
  NumArray(vector<int>& nums) {
    nums_ = nums;
    Build(nums);

    for (int i = 0; i < nums.size(); i++) {
      Update(1, i + 1, nums[i]);
    }
  }

  void update(int index, int val) {
    Update(1, index + 1, val - nums_[index]);
    nums_[index] = val;
  }

  int sumRange(int left, int right) {
    return Query(1, left + 1, right + 1);
  }

  vector<int> nums_;
};
