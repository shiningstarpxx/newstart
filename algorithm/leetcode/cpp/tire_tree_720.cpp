//
// Created by 裴星鑫 on 2022/3/17.
//
#include <string>
#include <vector>
#include <unordered_set>
using namespace std;

class Solution {
public:
    string longestWord(vector<string>& words) {
        // return HashHelper(words);
        return TrieHelper(words);
    }

    string HashHelper(vector<string>& words) {
        auto cmp = [&](string& a, string& b) -> bool {
            if (a.length() == b.length()) {
                return a > b;
            } else {
                return a.length() < b.length();
            }
        };

        sort(words.begin(), words.end(), cmp);

        string res;
        unordered_set<string> cache;
        for (auto &s : words) {
            if (s.size() == 1 || cache.count(s.substr(0, s.size() - 1))) {
                res = s;
                cache.emplace(s);
            }
        }
        return res;
    }

    class Trie {
      public:
        Trie() {
            children_ = vector<Trie*>(26, nullptr);
            is_end_ = false;
        }

        void Insert(string& s) {
            auto t = this;
            for (auto c : s) {
                if (t->children_[c - 'a'] == nullptr) {
                    t->children_[c - 'a'] = new Trie();
                }
                t = t->children_[c - 'a'];
            }
            t->is_end_ = true;
        }

        bool Query(string& s) {
            auto t = this;
            for (auto c : s) {
                if (t->children_[c - 'a'] == nullptr || t->children_[c-'a']->is_end_ == false) return false;
                t = t->children_[c - 'a'];
            }
            return t->is_end_;
        }

      private:
        vector<Trie *> children_;
        bool is_end_;
    };
    string TrieHelper(vector<string>& words) {
        Trie tire;
        for (auto& s : words) tire.Insert(s);

        string res;
        for (auto& s : words) {
            if (tire.Query(s)) {
                if (s.size() > res.size()) res = s;
                else if (s.size() == res.size() && s < res) res = s;
            }
        }
        return res;
    }
};