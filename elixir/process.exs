parent = self()

spawn fn -> send parent, {:hi, "woowwooow!"} end
receive do
  {:hi, msg} -> IO.puts msg
end
