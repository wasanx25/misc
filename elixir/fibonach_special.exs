# 1, 1, 2, 3, 5, 8, 13, 18, 31, 44, 75, 106, 137, 168, 305

defmodule Fibonach do
  def run(num) do
    IO.puts calculate(num + 1)
  end

  defp calculate(num, a \\ 0 , b \\ 1, prime \\ 2) do
    case num do
      1 -> a
      _ -> calculate(num - 1, a + b, select_number(a, prime), select_number(a, prime))
    end
  end

  defp select_number(a, prime) do
    if prime?(a), do: a, else: prime
  end

  defp prime?(n) do
    if n == 1 | rem(n, 2) == 0 | rem(n, 3) == 0 | rem(n, 5) == 0 | rem(n, 7) == 0 | rem(n, 11) == 0 do
      false
    else
      check(n)
    end
  end

  defp check(n) do
    trunc(:math.sqrt(n))
  end
end

Fibonach.run(10000)
