//
// Created by 裴星鑫 on 2022/5/25.
//
#include <vector>

using namespace std;

class Solution {
public:
    int jobScheduling(vector<int>& startTime, vector<int>& endTime, vector<int>& profit) {
        vector<tuple<int, int, int>> data;
        int n = startTime.size();
        for (int i = 0; i < n; i++) {
            data.push_back({endTime[i], startTime[i], profit[i]});
        }

        sort(data.begin(), data.end(), [](tuple<int, int, int>&a , tuple<int, int, int>& b) {
            return get<0>(a) < get<0>(b);
        });

        // TLE
        // vector<int> dp(n);
        // dp[0] = get<2>(data[0]);
        // for (int i = 0; i < n; i++) {
        //     dp[i] = get<2>(data[i]);
        //     for (int j = 0; j < i; j++) {
        //         if (get<0>(data[j]) <= get<1>(data[i])) dp[i] = max(dp[i], dp[j] + get<2>(data[i]));
        //     }
        //     dp[i] = max(i > 0 ? dp[i-1] : 0, dp[i]);
        // }
        // return dp.back();

        vector<int> dp(n);
        dp[0] = get<2>(data[0]);
        for (int i = 0; i < n; i++) {
            int l = 0, r = i-1;
            while (l <= r) {
                int mid = (l + r) / 2;
                if (get<0>(data[mid]) > get<1>(data[i])) r = mid -1;
                else l = mid + 1;
            }
            dp[i] = max(i > 0 ? dp[i-1] : 0,  l - 1 >= 0 ? dp[l-1] + get<2>(data[i]) : get<2>(data[i]));
        }
        return dp.back();
    }
};