dir    = ARGV[0]
search = ARGV[1]
target = ARGV[2]
word   = ARGV[3]

value = `grep -lr "#{search}" #{dir}`
value = value.split("\n")

value.each do |file|
  n = 0
  next if file =~ /.swp\Z/ || File.extname(file).empty?
  File.open(file, 'r') do |f|
    f.scrub.each.with_index(1) do |line, i|
      if line =~ /#{target}/
        `sed -i -e "#{i + n}a #{word}" #{file}`
        n += 1
      end
    end
  end
end
