test_input = [
  "7 6 4 2 1",
  "1 2 7 8 9",
  "9 7 6 2 1",
  "1 3 2 4 5",
  "8 6 4 4 1",
  "1 3 6 7 9"
]
case ARGV[0]
when "test"
  input = test_input
when "solve"
  input = File.read("input/day2_input.txt").split("\n")
else
  abort "test or solve?"
end

def report_valid(levels)
  direction_determined = false
  is_increasing = false
  levels.each_index do |index|
    if index < levels.count - 1
      level = Integer(levels[index])
      next_level = Integer(levels[index + 1])

      if !direction_determined && level > next_level
        is_increasing = false
        direction_determined = true
      elsif !direction_determined && level < next_level
        is_increasing = true
        direction_determined = true
      end

      delta = (level - next_level).abs
      if delta < 1 || delta > 3
        return false
      elsif direction_determined && is_increasing && next_level < level
        return false
      elsif direction_determined && !is_increasing && next_level > level
        return false
      end
    end
  end
  return true
end

safe_report_count = 0

input.each do |report|
  levels = report.split
  if report_valid(levels)
    safe_report_count += 1
  else
    levels.each_index do |i|
      modified_levels = levels
      modified_levels.delete_at(i)
      if report_valid(modified_levels)
        safe_report_count += 1
        break
      end
    end
  end
end

puts "Number of safe reports: #{safe_report_count}"
