//
// Created by 裴星鑫 on 2022/4/4.
//
#include <vector>
using namespace std;
class NumArray {
public:
    NumArray(vector<int>& nums) {
        size_ = nums.size() + 1;
        data_ = vector<int>(size_);
        origin_ = nums;
        for (int i = 0; i < origin_.size(); i++) {
            Update(i, origin_[i]);
        }
    }

    int Query(int index) {
        int ans = 0;
        for (int i = index + 1; i > 0; i -= LowBit(i)) ans += data_[i];
        return ans;
    }

    int LowBit(int x) { return x & -x; }

    void Update(int index, int delta) {
        for (int i = index + 1; i < size_; i+= LowBit(i)) data_[i] += delta;
    }

    void update(int index, int val) {
        int delta = val - origin_[index];
        Update(index, delta);
        origin_[index] = val;
    }

    int sumRange(int left, int right) {
        return Query(right) - Query(left - 1);
    }

    vector<int> origin_;
    vector<int> data_;
    int size_;

};

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray* obj = new NumArray(nums);
 * obj->update(index,val);
 * int param_2 = obj->sumRange(left,right);
 */

 // 前缀和，超时
 /*
     NumArray(vector<int>& nums) {
        nums_ = nums;
        sums_ = nums_;
        for (int i = 1; i < nums.size(); i++) {
            sums_[i] += sums_[i - 1];
        }
    }

    void update(int index, int val) {
        int diff = val - nums_[index];
        for (int i = index; i < nums_.size(); i++) {
            sums_[i] += diff;
        }
        nums_[index] = val;
    }

    int sumRange(int left, int right) {
        return sums_[right] - sums_[left] + nums_[left];
    }

    vector<int> nums_;
    vector<int> sums_;
    */