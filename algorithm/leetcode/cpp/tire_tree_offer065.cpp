//
// Created by 裴星鑫 on 2022/3/22.
//
#include <string>
#include <vector>

using namespace std;

class Solution {
public:
    class Tire {
        public:
        Tire() {
            children_ = vector<Tire*> (26, nullptr);
        }
        void Insert(string& s) {
            auto t = this;
            for (auto c : s) {
                if (t->children_[c - 'a'] == nullptr) t->children_[c - 'a'] = new Tire;
                t = t->children_[c - 'a'];
            }
        }
        bool NewWorld(string& s) {
            auto t = this;
            for (auto c : s) {
                if (t->children_[c - 'a'] == nullptr) return true;
                t = t->children_[c - 'a'];
            }
            return false;
        }
        private:
        vector<Tire*> children_;
    };
    int minimumLengthEncoding(vector<string>& words) {
        auto cmp = [](string& a, string& b) {
            return a.size() > b.size();
        };
        sort(words.begin(), words.end(), cmp);

        int res = 0;
        Tire t;
        for (auto& s : words) {
            reverse(s.begin(), s.end());
            if (!t.NewWorld(s)) continue;
            t.Insert(s);
            res += s.size() + 1;
        }
        return res;
    }
};