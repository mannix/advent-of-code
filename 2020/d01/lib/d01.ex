defmodule D01 do

  def findTheAnswer([]) do
    IO.puts("The answer you seek was not found.")
  end

  def findTheAnswer([head | tail]) do
    Enum.each(tail, fn x ->
      if head + x == 2020 do
        IO.puts("#{head} + #{x} = #{head + x}")
        IO.puts("#{head} * #{x} = #{head * x}")
        System.halt(0)
      end
    end)
    findTheAnswer(tail)
  end

  def main do
    {:ok, contents} = File.read("../input")
    contents
    |> String.split()
    |> Enum.map(fn x -> String.to_integer(x) end)
    |> findTheAnswer()
  end

end

D01.main()
