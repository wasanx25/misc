# IO.puts Enum.map ["kikuchi", "wataru", "23"], fn x -> x <> x end

# IO.puts ["kikuchi", "wataru", "23"]
#   |> Enum.map(fn x -> x <> x end)

IO.puts %{"kikuchi" => "wataru", "tanaka" => "takeshi", "sato" => "takeru"}
  |> Enum.map(fn {k, v}, acc -> acc <> v end)
