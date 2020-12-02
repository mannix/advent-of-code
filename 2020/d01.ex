{:ok, contents} = File.read("input")

contents
|> String.split()
|> Enum.each(IO.puts())

