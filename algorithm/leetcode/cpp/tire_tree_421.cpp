//
// Created by 裴星鑫 on 2022/3/21.
//
#include <vector>

using namespace std;

class Solution {
public:
    int findMaximumXOR(vector<int>& nums) {
//         return BruteForce(nums);
        return HashCalc(nums);
    }
    int BruteForce(vector<int>& nums) { // TLE
        int res = 0;
        for (int i = 0; i < nums.size(); i++) {
            for (int j = i + 1; j < nums.size(); j++) {
                res = max(res, nums[i] ^ nums[j]);
            }
        }
        return res;
    }

    static const int BIT_NUM = 30;
    class Tire {
        public:
        Tire() {
            children_ = vector<Tire*> (2, nullptr);
        }
        void Insert(int num) {
            auto t = this;
            for (int i = BIT_NUM; i >= 0; i--) {
                int index = (num >> i) & 1;
                if (t->children_[index] == nullptr) t->children_[index] = new Tire;
                t = t->children_[index];
            }
        }
        int Calc(int num) {
            auto t = this;
            int res = 0;
            for (int i = BIT_NUM; i >= 0; i--) {
                int index = (num >> i) & 1;
                if (index == 0) {
                    if (t->children_[1]) {
                        t = t->children_[1];
                        res = res * 2 + 1;
                    } else {
                        t = t->children_[0];
                        res = res * 2;
                    }
                } else {
                    if (t->children_[0]) {
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

    int HashCalc(vector<int>& nums) {
        Tire t;
        int res = 0;
        for (int i = 1; i < nums.size(); i++) {
            t.Insert(nums[i - 1]);
            res = max(res, t.Calc(nums[i]));
        }
        return res;
    }
};