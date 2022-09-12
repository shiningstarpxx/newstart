//
// Created by 裴星鑫 on 2022/8/26.
//

#ifndef BASE_OPERATION_BASE_OPERATION_SEGMENT_TREE_SEGMENT_TREE_H_
#define BASE_OPERATION_BASE_OPERATION_SEGMENT_TREE_SEGMENT_TREE_H_

#include <math.h>
#include <vector>

using namespace std;

class SegmentTree {
 public:
  SegmentTree(const vector<int>& arr) {
    int n = arr.size();
    int x = (int)(ceil(log2(n)));
    int max_size = 2 * (int)pow(2, x) - 1;
    data_ = vector<int>(max_size);
  }
 private:
  vector<int> data_;
};

#endif //BASE_OPERATION_BASE_OPERATION_SEGMENT_TREE_SEGMENT_TREE_H_
