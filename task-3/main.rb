
require 'nokogiri'
require 'open-uri'
require 'pry'


puts("Enter search query: ")
userin = gets
doc = Nokogiri::HTML(open('https://www.google.com/search?start=0&end=11&q='+userin))
doc.xpath('//div/a/div[text()]').each do |link|
    puts link.content
end

