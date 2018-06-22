# フィボナッチ数列の
# n番目の値を求める

defmodule Fibonach do
  def run(num) do
    IO.puts calculate(num + 1)
  end

  defp calculate(num, a \\ 0 , b \\ 1) do
    case num do
      1 -> a
      _ -> calculate(num - 1, a + b, a)
    end
  end
end

Fibonach.run(10000)
