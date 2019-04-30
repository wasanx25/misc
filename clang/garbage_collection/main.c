#include <stdio.h>
#include <stdlib.h>
#include <string.h>

const char *concat(const char *s1, const char *s2)
{
  size_t l1 = strlen(s1);
  size_t l2 = strlen(s1);

  char *result = malloc(l1 + l2 + 1);

  memcpy(result, s1, l1);
  memcpy(result + l1, s2, l2);

  result[l1 + l2] = '\0';

  return result;
}

int main(int argc, const char *argv[])
{
  const char *abc = "abc";
  const char *def = "def";

  puts(concat(abc, def));

  return 0;
}