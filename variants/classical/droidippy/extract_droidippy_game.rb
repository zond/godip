#!/usr/bin/env ruby

fileReg = /^game_(\d+)\.txt$/
phaseReg = /\s*(\d+)\s*\|\s*(\S+),\s*(\d+)\s[^|]+\s*\|\s*(\S+)\s*$/
posReg = /\s*(\S+)\s*\|\s*(\S+)\s*\|\s*(\S+)\s*$/
orderReg = /\s*\S+:\s*(.*\S+)\s*$/

gameId = fileReg.match(ARGV[0])[1]

translations = {
  "lyo" => "gol",
  "mao" => "mid",
  "nwg" => "nrg",
  "nao" => "nat",
}

`echo 'select id, name, type from gamephase where game_id = #{gameId} order by ordinal' | psql -t dippy | grep -v '^$'`.each_line do |line|
  if (match = phaseReg.match(line)) != nil
    phaseId = match[1]
    season = match[2]
    year = match[3]
    type = match[4]
    puts "PHASE #{year} #{season} #{type}"
    puts "POSITIONS"
    `echo 'select type, power, province from gameposition where phase_id = #{phaseId}' | psql -t dippy | grep -v '^$'`.each_line do |line|
			translations.each do |k,v|
				line.gsub!(k,v)
			end
      if (m = posReg.match(line)) != nil
				puts "#{m[2]}: #{m[1]} #{m[3]}"
			else
				raise "Unknown line #{line}"
      end
    end
    puts "ORDERS"
    `echo 'select text from gameorder where phase_id = #{phaseId}' | psql -t dippy | grep -v '^$'`.each_line do |line|
      translations.each do |k,v| 
        line.gsub!(k,v)
      end
			if (match = orderReg.match(line)) != nil
				parts = match[1].split
				if !(/build/ =~ line) 
					parts = parts.reject! do |p|
						["Army", "Fleet"].include?(p)
					end
				end
				puts parts.join(" ")
			else 
        raise "Unknown line #{line}"
			end
    end
    
  end
end
