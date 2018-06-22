defmodule Mammal do
  defmodule Human do
    defstruct name: "", age: 0
  end

  defmodule Cat do
    defstruct type: "", name: ""
  end

  def info do
    cat = %Cat{name: "sheez", type: "mia"}
    IO.inspect cat
  end
end

Mammal.info
