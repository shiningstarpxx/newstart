//
// Created by 裴星鑫 on 2022/5/26.
//
#include <map>
#include <vector>

using namespace std;

class Solution {
public:
    vector<int> fallingSquares(vector<vector<int>>& positions) {
        int n = positions.size();
        map<int, vector<int>> data;

        vector<int> res;
        for (int i = 0; i < n; i++) {
            int l = positions[i][0], r = l + positions[i][1], height = positions[i][1];

            auto lt = data.lower_bound(l);
            auto rt = data.lower_bound(r);
            int h = 0;
            for (auto it = lt; it != rt; it++) {
                h = max(h, it->second[2] + height);
            }
            data.emplace(r, {l, r, height});
        }
        return res;
    }

    vector<int> BruteForce(vector<vector<int>>& positions) {
        int n = positions.size();
        vector<int> res;
        vector<int> height(n);
        int h = 0;
        for (int i = 0; i < n; i++){
            int left = positions[i][0], right = positions[i][0] + positions[i][1];
            int t = positions[i][1];
            for (int j = 0; j < i; j++) {
                if (positions[j][0] + positions[j][1] > left && right > positions[j][0]) {
                    t = max(t, positions[i][1] + height[j]);
                }
            }
            height[i] = t;
            h = max(h, height[i]);
            res.push_back(h);
        }
        return res;
    }
};
