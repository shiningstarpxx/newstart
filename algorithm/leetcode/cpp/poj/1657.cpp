//
// Created by 裴星鑫 on 2022/10/5.
//

#include <stdio.h>
#include <math.h>

int main() {
  int test_nums;
  scanf("%d", &test_nums);

  for (int i = 0; i < test_nums; i++) {
    char begin[5], end[5];
    scanf("%s %s", begin, end);
    int x = abs(begin[0] - end[0]), y = abs(begin[1] - end[1]);
    if (x == 0 && y == 0) printf("0 0 0 0\n");
    else {
      printf("%d", x < y ? y : x);
      printf(" %d", (x == y || x == 0 || y == 0) ? 1 : 2);
      printf(" %d", (x == 0 || y == 0) ? 1 : 2);
      if (abs(x - y) % 2 != 0) printf(" Inf\n");
      else printf(" %d\n", x == y ? 1 : 2);
    }
  }
  return 0;
}