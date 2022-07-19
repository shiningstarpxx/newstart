//
// Created by 裴星鑫 on 2022/6/6.
//
#include <vector>

using namespace std;

struct Node {
  int l, r;  // current index, left child, right child
  int val;
  Node(int ll, int rr) {l = ll, r = rr;}
};

class SegmentTree {
 public:
  void Build(int n) {
    data_ = vector<Node>(4 * n);
    DoBuild(1, 1, n);  // from 1
  }

  // [l, r] 需要注意是闭区间
  void DoBuild(int k, int l, int r) {
    data_[k] = Node(l, r);
    if (l == r) return;
    int mid = l + (r -l >> 1);
    DoBuild(k << 1, l, mid);
    DoBuild(k << 1 + 1, mid + 1, r);
  }

  void PushUp(int k) {
    data_[k].val = data_[k << 1].val + data_[k << 1 | 1].val;
  }

  int Calc(int k, int l, int r) {
    if (data_[k].l >= l && data_[k].r <= r) return data_[k].val;  // 如果在查询范围里，直接返回
    int mid = data_[k].l + (data_[k].r - data_[k].l >> 1);
    int ans = 0;
    if (mid >= l) ans += Calc(k >> 1, l, r);
    if (mid < r) ans += Calc(k >> 1 | 1, l, r);
    return ans;
  }

  void Update(int k, int index, int val) {
    if (data_[k].l == index && data_[k].r == index) {
      data_[k].val += val;
      return;
    }
    int mid = data_[k].l + (data_[k].r - data_[k].l >> 1);
    if (mid <= index) Update(mid, index, val);
    else Update(mid + 1, index, val);
  }

  vector<Node> data_;
};