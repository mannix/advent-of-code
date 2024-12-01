test_input = [
"3   4",
"4   3",
"2   5",
"1   3",
"3   9",
"3   3"
]
case ARGV[0]
when "test"
  input = test_input
when "solve"
  input = File.read("input/day1_input.txt").split("\n")
else
  abort "test or solve?"
end

left_list = []
right_list = []

input.each do |line|
  left_list << Integer(line.split[0])
  right_list << Integer(line.split[1])
end

left_list.sort!
right_list.sort!

# Part one
total_distance = 0
left_list.each_index do |i|
  total_distance += (left_list[i] - right_list[i]).abs
end
puts "Total distance: #{total_distance}"

# Part two
similarity_sore = 0
left_list.each do |left_element|
  right_list_occurances = right_list.select {|right_element| left_element == right_element}
  similarity_sore += left_element * right_list_occurances.count
end
puts "Similarity score: #{similarity_sore}"