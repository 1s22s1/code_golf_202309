ascii = Hash.new(0)

gets.chomp.split("").map{|c| ascii[c] += 1}

ascii.each do |key, value|
  puts "#{key} => #{value}"
end
