
require 'nokogiri'
require 'open-uri'  

# Fetch and parse HTML document
name = gets
name = name.chomp
puts name

doc = Nokogiri::HTML(open('https://www.google.com/search?q=Linux'))
doc.css('div.ellip')collect do |qry|
    qry.css('h2').text


#doc = Nokogiri::HTML(open('https://nokogiri.org/tutorials/installing_nokogiri.html'))

#puts "---------------------------------------------"
#puts "### Search for nodes by css"
#doc.css('nav ul.menu li a', 'article h2').each do |link|
#  puts link.content
#end
#puts "--------------------------------------------"
#puts "### Search for nodes by xpath"
#doc.xpath('//nav//ul//li/a', '//article//h2').each do |link|
 # puts link.content
#end
#puts " ------------------------------------------- "
#puts "### Or mix and match."
#doc.search('nav ul.menu li a', '//article//h2').each do |link|
 # puts link.content
#end
