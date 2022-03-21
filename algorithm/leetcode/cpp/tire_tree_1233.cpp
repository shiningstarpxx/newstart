//
// Created by 裴星鑫 on 2022/3/21.
//
#include <string>
#include <vector>

using namespace std;

class Solution {
public:
    class Tire {
        public:
        Tire() {
            children_ = vector<Tire*> (27, nullptr);
            is_end_ = false;
        }
        void Insert(string& s) {
            auto t = this;
            for (auto c : s) {
                int index = 0;
                if (c == '/') index = 26;
                else index = c - 'a';
                if (t->children_[index] == nullptr) t->children_[index] = new Tire;
                t = t->children_[index];
            }
            t->is_end_ = true;
        }
        bool Prefix(string& s) {
            auto t = this;
            bool flag = false;
            for (int i = 0; i <s.size(); i++) {
                char c = s[i];
                int index = 0;
                if (c == '/') index = 26;
                else index = c - 'a';
                if (t->children_[index] == nullptr) return false;
                t = t->children_[index];
                // core match code: for root directory match
                if (t->is_end_ && s[i + 1] == '/') return true;
            }
            return false;
        }
        private:
        vector<Tire*> children_;
        bool is_end_;
    };

    vector<string> removeSubfolders(vector<string>& folder) {
        auto cmp = [](string& a, string& b) {
            return a.size() < b.size();
        };

        sort(folder.begin(), folder.end(), cmp);

        vector<string> res;
        Tire t;
        for (auto& s : folder) {
            if (t.Prefix(s)) continue;
            res.push_back(s);
            t.Insert(s);
        }
        return res;
    }
};