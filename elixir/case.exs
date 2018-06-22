sample = fn a ->
  case a do
    1 -> IO.puts "one"
    2 -> IO.puts "two"
    3 -> IO.puts "three"
    4 -> IO.puts "four"
    _ -> IO.puts "not"
  end
end

sample.(5)
