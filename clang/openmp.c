#include<stdio.h>
#include<omp.h>

int main() {
  int a[1000];
  int i;

  #pragma omp parallel for
  for (i = 0; i < 1000; i++) {
    a[i] = i;
    printf("a => %d\n", a[i]);
  }

  return 0;
}
