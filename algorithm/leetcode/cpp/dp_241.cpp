//
// Created by Pei Xingxin on 1/7/2022.
//

class Solution {
public:
    vector<int> dfs(vector<vector<vector<int>>> f, int l, int r, vector<int>& ops) {
        if (f[l][r].empty()) {
            if (l == r) {
                f[l][r].push_back(ops[l]);
            } else {
                for (int i = l; i < r; i+=2) {
                    auto lv = dfs(f, l, i, ops);
                    auto rv = dfs(f, i+2, r, ops);
                    for (auto ll : lv) {
                        for (auto rr : rv) {
                            int d = 0;
                            if (ops[i+1] == -1) d = ll + rr;
                            else if (ops[i+1] == -2) d = ll - rr;
                            else d = ll * rr;
                            f[l][r].push_back(d);
                        }
                    }
                }
            }
        }
        return f[l][r];
    }
    vector<int> diffWaysToCompute(string expression) {
        vector<int> ops;
        // 解析字符串
        for (int i = 0; i < expression.size();) {
            int d = 0;
            while (isdigit(expression[i])) {
                d = d * 10 + (expression[i] - '0');
                i++;
            }
            ops.push_back(d);
            if (i < expression.size()) {
                if (expression[i] == '+') ops.push_back(-1);
                else if (expression[i] == '-') ops.push_back(-2);
                else ops.push_back(-3);
                i++;
            }
        }

        vector<vector<vector<int>>> f(ops.size(), vector<vector<int>>(ops.size()));
        return dfs(f, 0, ops.size() - 1, ops);
    }
};