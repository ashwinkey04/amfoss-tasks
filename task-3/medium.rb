require 'nokogiri'
require 'pry'
require 'open-uri'

doc = Nokogiri::HTML(open('https://www.yogajournal.com/poses/types'))

doc.css('.m-card--header').collect do |pose_type|
    pose_type.css('h2').text
end 
