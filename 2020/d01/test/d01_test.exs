defmodule D01Test do
  use ExUnit.Case
  doctest D01

  test "greets the world" do
    assert D01.hello() == :world
  end
end
