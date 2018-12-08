#include <curses.h>
#include <stdio.h>
#include <term.h>

int main(void) {
  char *cup, *p, *q;
  setupterm (NULL, 0, NULL);
  cup = tigetstr ("cup");
  p = tparm (cup, 10, 20);
  printf ("p=|%s|\n", p);
  q = tparm (cup, 30, 40);
  printf ("p=|%s|, q=|%s|\n", p, q);
}