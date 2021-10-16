
// 代码写的实在是太丑了
class Solution {
public:
    vector<string> addOperators(string num, int target) {
        vector<string> ans;
        auto n = num.length();
        function<void(string, int, long, long)> dfs = [&](string expr, int index, long res, long mul) {
            if (index == n) {
                if (target == res) {
                    ans.push_back(expr);
                }
                return;
            }

            long val = 0;
            auto t = expr.length();
            if (index > 0) {
                expr += "0";
            }
            for (int i = index; i < n; i++) {
                if (num[index] == '0' && i > index) break;
                val = val * 10 + (num[i] - '0');
                expr += num[i];
                if (index == 0) {
                    dfs(expr, i + 1, val, val);
                } else {
                    expr[t] = '+';
                    dfs(expr, i + 1, res + val, val);
                    expr[t] = '-';
                    dfs(expr, i + 1, res - val, -val);
                    expr[t] = '*';
                    dfs(expr, i + 1, res - mul + mul * val, mul * val);
                }
            }
        };

        string expr;
        dfs(expr, 0, 0, 0);
        return ans;
    }
};