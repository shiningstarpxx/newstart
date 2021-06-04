#include <stdio.h>
#include <algorithm>

using namespace std;

#define maxn 32010

int c[maxn] = {0};

int lowbit(int i) {
    return i & (-i);
}

void add(int i, int val) {
    while (i <= maxn) {
        c[i] += val;
        i += lowbit(i);
    }
}

int sum(int i) {
    int s = 0;
    while (i > 0) {
        s += c[i];
        i -= lowbit(i);
    }
    return s;
}

int main() {
    int ans[maxn];
    int max_x = 0;
    int n, x, y;
    scanf("%d", &n);
    for (int i = 0; i < n; ++i) {
        scanf("%d%d", &x, &y);
        ++x;
        ans[sum(x)]++;
        add(x, 1);
    }

    for (int i = 0; i < n; ++i)
        printf("%d\n", ans[i]);
    return 0;
}