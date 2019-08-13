import sys
sys.path.append('/home/ashwin/.local/bin')
import pytesseract
from PIL import Image
from pytesseract import image_to_string
str = image_to_string(Image.open('/home/ashwin/Pictures/cap.png')) #use your image's path
n1 = int(str[0:1])
n2 = int(str[2:3])
if '+' in str:
    print(n1+n2)
elif '-' in str:
    print(n1-n2)
elif '*' in str:
    print(n1*n2)
elif '/' in str:
    print(n1/n2)
else:
    print('unable')
