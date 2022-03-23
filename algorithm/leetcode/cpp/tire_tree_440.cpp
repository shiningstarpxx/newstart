//
// Created by 裴星鑫 on 2022/3/23.
//

//给定整数 n 和 k，返回 [1, n] 中字典序第 k 小的数字。
//
//
//
// 示例 1:
//
//
//输入: n = 13, k = 2
//输出: 10
//解释: 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
//
//
// 示例 2:
//
//
//输入: n = 1, k = 1
//输出: 1
//
//
//
//
// 提示:
//
//
// 1 <= k <= n <= 10⁹
//
// Related Topics 字典树 👍 394 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
public:
    int findKthNumber(int n, int k) {
      int res = 1;
      k--;
      while (k > 0) {
        int steps = GetSteps(res, n);
        if (steps <= k) {
          k -= steps;
          res++;
        } else {
          k--;
          res *= 10;
        }
      }
      return res;
    }

    int GetSteps(int curr, long n) {
      int steps = 0;
      long first = curr;
      long end = curr;
      while (first <= n) {
        steps += min(end, n) - first + 1;
        first = first * 10;
        end = end * 10 + 9;
      }
      return steps;
    }

    int min(int a, int b) {
      return a <= b ? a : b;
    }
};
//leetcode submit region end(Prohibit modification and deletion)
