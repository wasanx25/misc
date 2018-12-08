#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char const *argv[]) {
  int i = 0, n = atoi (argv [1]);
  char c [] = {'.', 'o', '0'};
  for (i = 0; i < n; i++) {
    printf("%c\033[D", c [i % 3]);
    fflush(stdout);
    sleep(1);
  }
}
