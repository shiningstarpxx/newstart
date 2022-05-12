//
// Created by 裴星鑫 on 2022/5/12.
//
#include <vector>

using namespace std;

// 有趣的发现： 在使用vector的情况下，会超过限制，改成数组就不会了。说明大内存的申请释放很耗资源
class Solution {
public:
    int presum[101][101];
    int maxSumSubmatrix(vector<vector<int>>& matrix, int k) {
        int row = matrix.size();
        int column = matrix[0].size();

        CalcPreSum(matrix);
        // auto presum = CalcPreSum(matrix);

        int ans = INT_MIN;
        for (int bottom = 0; bottom < row; bottom++) {
            for (int top = 0; top <= bottom; top++) {
                for (int right = 0; right < column; right++) {
                    for (int left = 0; left <= right; left++) {
                        int local = presum[bottom+1][right+1] - presum[bottom+1][left] - presum[top][right+1] + presum[top][left];
                        if (local <= k) ans = max(local, ans);
                    }
                }
            }
        }
        return ans;
    }

    // vector<vector<int>> CalcPreSum(vector<vector<int>>& matrix) {
    void CalcPreSum(vector<vector<int>>& matrix) {

        int row = matrix.size();
        int column = matrix[0].size();

        // vector<vector<int>> data(row+1, vector<int>(column+1));

        for (int i = 1; i <= row; i++) {
            for (int j = 1; j <= column; j++) {
                presum[i][j] = matrix[i-1][j-1] + presum[i][j-1] + presum[i-1][j] - presum[i-1][j-1];
            }
        }
        // return std::move(data);
    }
};