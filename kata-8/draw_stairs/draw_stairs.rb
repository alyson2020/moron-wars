def draw_stairs(n)
  res = ''
  (0...n).each do |i|
    res << ' ' * i
    res << "I\n"
  end
  res.chomp
end

puts(draw_stairs(1), "I")
puts(draw_stairs(2), "I\n I")
puts(draw_stairs(3), "I\n I\n  I")
