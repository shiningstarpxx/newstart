//
// Created by 裴星鑫 on 2022/1/11.
//
#include <queue>
#include <unordered_set>
#include <vector>

using namespace std;

class Solution {
public:
    unordered_set<long> blocks_;
    const int D = 131;
    int size_ = 0;
    int dir[4][2] = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
    bool isEscapePossible(vector<vector<int>>& blocked, vector<int>& source, vector<int>& target) {
        for (auto v : blocked) {
            blocks_.insert(v[0] * D + v[1]);
        }
        size_ = blocks_.size();
        return check(source, target) && check(target, source);
    }

    bool check(vector<int>& source, vector<int>& target) {
        unordered_set<long> visited;
        deque<pair<int, int>> q;
        q.push_back({source[0], source[1]});
        visited.insert(source[0] * D + source[1]);

        while (!q.empty() && visited.size() <= (size_ * (size_ - 1) / 2)) {
            auto t = q.front();
            q.pop_front();
            for (int i = 0; i < 4; i++) {
                if (t.first + dir[i][0] < 0 || t.second + dir[i][1] < 0 || t.first + dir[i][0] >= 1e6 || t.second + dir[i][1] >= 1e6)
                    continue;
                if (t.first + dir[i][0] == target[0] && t.second + dir[i][1] == target[1]) return true;
                long d = (t.first + dir[i][0]) * D + (t.second + dir[i][1]);
                if (visited.count(d) || blocks_.count(d)) continue;
                q.push_back({t.first + dir[i][0], t.second + dir[i][1]});
                visited.emplace(d);
            }
        }
        return visited.size() > (size_ * (size_ - 1) / 2);
    }
};