test_input = ['xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))']

case ARGV[0]
when 'test'
  input = test_input
when 'solve'
  input = File.read('input/day3_input.txt')
else
  abort 'test or solve?'
end

match = input.scan(/mul\((\d{1,3}),(\d{1,3})\)/)
total = match.sum { |a, b| a.to_i * b.to_i }
puts "Total: #{total}"
