
class Solution {
public:
    int arrangeCoins(int n) {
        return binarySearch(n);
    }

    int binarySearch(int n) {
        int left = 1, right = n;
        while (left < right) {
            int mid = (right - left + 1) / 2 + left;
            if ((long long) mid * (mid + 1) <= (long long) 2 * n) {
                left = mid;  // 如果left 就是等于 mid，也就是说(right - left + 1)/2 出现了没有step的情况，也就是 right <= left, 不成立
            } else {
                right = mid - 1;  // 必须左移一位，如果不移动，最坏情况是 right = mid, 也就说 right = (left + 1) 这个会陷入死循环
            }
        }
        return left;
    }

    int bruteForce(int n) {
         // 本质上求不大于n的数字 (k)(k-1)/2 <= n

        long ans = 0;
        for (long i = sqrt(n); i * (i + 1) / 2 <= (long)n; i++) {
            ans = i;
        }
        return (int)ans;
    }
};