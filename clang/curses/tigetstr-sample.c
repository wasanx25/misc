#include <curses.h>
#include <term.h>

int main(void) {
  char *p;
  setupterm (NULL, 0, NULL);
  p = tigetstr ("cup");
  p = tparm (p, 10, 20);
  putp (p);
}
