//
// Created by 裴星鑫 on 2022/5/21.
//
#include <string>

using namespace std;

class Solution {
public:
    long long numberOfWays(string s) {  // 本质上是前缀和的题目
        long long ans = 0;
        int n = s.size();
        int count_1 = count(s.begin(), s.end(), '1');
        int current_1 = 0;
        for (int i = 0; i < s.size(); i++) {
            if (s[i] == '1') {  // 计算左边的0的个数，计算右边的0的个数
                ans += (i - current_1) * (n - i - (count_1 - current_1));
                current_1++;
            } else {
                ans += current_1 * (count_1 - current_1);
            }
        }
        return ans;
    }
};
