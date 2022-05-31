//
// Created by 裴星鑫 on 2022/5/31.
//
#include <queue>
#include <string>
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
public:
    unordered_map<char, vector<char>> graph_; // 构建有向图
    unordered_map<char, int> degree_;

    string alienOrder(vector<string>& words) {
        for (auto& s : words) {
            for (auto c : s) {
                if (!graph_.count(c)) graph_[c] = vector<char>();
            }
        }

        for (int i = 1; i < words.size(); i++) {
            if (!BuildGraph(words[i-1], words[i])) return "";
        }

        queue<char> q;  // 拓扑排序
        string res;
        for (auto it = graph_.begin(); it != graph_.end(); it++) {
            if (!degree_.count(it->first)) q.push(it->first);
        }

        while (!q.empty()) {
            char c = q.front();
            res.push_back(c);
            q.pop();
            auto& related_chars = graph_[c];
            for (auto v : related_chars) {
                degree_[v]--;
                if (degree_[v] == 0) q.push(v);
            }
        }

        return res.size() == graph_.size() ? res : "";
    }

    bool BuildGraph(string& a, string& b) {
        int i = 0, j = 0;
        int len = min(a.size(), b.size());

        while (i < len && j < len) {
            if (a[i] != b[j]) {
                graph_[a[i]].push_back(b[j]);
                degree_[b[j]]++;
                break;
            }
            i++;
            j++;
        }
        if (i == len && a.size() > b.size()) return false;
        return true;
    }
};