//
// Created by 裴星鑫 on 2022/3/23.
//
#include <vector>

using namespace std;

class Solution {
 public:

  static const int Num = 30;

  class Tire {
   public:
    Tire() {
      children_ = vector<Tire*> (2, nullptr);
    }
    void Insert(int n) {
      auto t = this;
      for (int i = Num; i >= 0; i--) {
        int index = (n >> i) & 1;
        if (t->children_[index] == nullptr) t->children_[index] = new Tire;
        t = t->children_[index];
      }
    }
    int GetMax(int n) {
      auto t = this;
      int res = 0;
      for (int i = Num; i >= 0; i--) {
        int index = (n >> i) & 1;
        if (index == 0) {
          if (t->children_[1] != nullptr) {
            t = t->children_[1];
            res = res * 2 + 1;
          } else {
            t = t->children_[0];
            res = res * 2;
          }
        } else {
          if (t->children_[0] != nullptr) {
            t = t->children_[0];
            res = res * 2 + 1;
          } else {
            t = t->children_[1];
            res = res * 2;
          }
        }
      }
      return res;
    }
   private:
    vector<Tire*> children_;
  };

  int findMaximumXOR(vector<int>& nums) {
    Tire t;
    int res = 0;
    for (int i = 1; i < nums.size(); i++) {
      t.Insert(nums[i - 1]);
      res = max(res, t.GetMax(nums[i]));
    }
    return res;
  }
};
