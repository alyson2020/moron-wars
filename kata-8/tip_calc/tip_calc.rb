ef calculate_tip(amount, rating)
  case rating.downcase
  when 'terrible'
    0
  when 'poor'
    (amount * 0.05).ceil
  when 'good'
    (amount * 0.1).ceil
  when 'great'
    (amount * 0.15).ceil
  when 'excellent'
    (amount * 0.2).ceil
  else
    'Rating not recognised'
  end
end

# Terrible: tip 0%
# Poor: tip 5%
# Good: tip 10%
# Great: tip 15%
# Excellent: tip 20%

puts(calculate_tip(30, "poor"), 2)
puts(calculate_tip(20, "Excellent"), 4)
puts(calculate_tip(20, "hi"), 'Rating not recognised')
puts(calculate_tip(107.65, "GReat"), 17)
puts(calculate_tip(20, "great!"), 'Rating not recognised')
