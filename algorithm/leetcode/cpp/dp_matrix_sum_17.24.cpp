//
// Created by 裴星鑫 on 2022/5/12.
//
#include <vector>

using namespace std;

class Solution {
public:
    vector<int> getMaxMatrix(vector<vector<int>>& matrix) {
        int row = matrix.size();
        int column = matrix[0].size();

        vector<vector<int>> presum = PreSum(matrix);

        int ans = matrix[0][0];
        vector<int> res = {0, 0, 0, 0};
        for (int bottom = 0; bottom < row; bottom++) {
            for (int top = 0; top <= bottom; top++) {
                int local_sum = 0, left = 0;
                for (int right = 0; right < column; right++) {
                    local_sum = presum[bottom + 1][right + 1] + presum[top][left] - presum[bottom+1][left] - presum[top][right+1];
                    if (local_sum > ans) {
                        ans = local_sum;
                        res = {top, left, bottom, right};
                    }
                    if (local_sum < 0) {
                        local_sum = 0;
                        left = right + 1;
                    }
                }
            }
        }
        return res;
    }

    vector<vector<int>> PreSum(vector<vector<int>>& matrix) {
        int row = matrix.size();
        int column = matrix[0].size();

        vector<vector<int>> presum(row + 1, vector<int>(column + 1));

        for (int i = 1; i <= row; i++) {
            for (int j = 1; j <= column; j++) {
                presum[i][j] = matrix[i-1][j-1] + presum[i-1][j] + presum[i][j-1] - presum[i-1][j-1];
            }
        }
        return std::move(presum);
    }
};
