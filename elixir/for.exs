# IO.puts for {name, 2} <- [{"kikuchi", 2}, {"tanaka", 1}, {"sato", 2}], do: name <> name

# result =
#   for x <- ["kikuchi", "tanaka", "sato"],
#       y <- ["wataru", "taro", "saito"],
#       do: x <> y
#
# IO.inspect result


# result =
#   for x <- ["kikuchi", "tanaka", "sato"],
#       y <- ["wataru", "taro", "saito"],
#       into: %{},
#       do: {x, y}
#
# IO.inspect result

result =
  for x <- [1, 2, 3],
      y <- [4, 5, 6],
      z <- [7, 8, 9],
      into: %{},
      do: {x, y, z}

IO.inspect result
