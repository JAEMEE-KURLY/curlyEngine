from urllib.request import Request
from urllib.request import urlopen
from bs4 import BeautifulSoup

url = 'https://www.ssg.com/search.ssg?target=all&query=%EA%B9%80%EC%B9%98&src_area=late'
req = Request(url,headers={'User-Agent':'Mozila/5.0'})
webpage = urlopen(req)
soup = BeautifulSoup(webpage, 'html.parser')
resultList = soup.find_all("li")

# print(resultList)
classList = []
for i in resultList:
    # print(i)
    if 'class' in i.attrs.keys():
        print(i.attrs['class'])
        classList.append(i.attrs['class'])

print(classList)

for i in resultList:
    nameList = i.find("div", "title").find("a")
    name = nameList.find("em", "tx_ko")
    priceList = i.find("div", "cunit_price")
    price = priceList.find("em", "ssg_price")
    # print(name.text, price.text)

# url = 'https://search.shopping.naver.com/search/all?where=all&frm=NVSCTAB&query=%EA%B9%80%EC%B9%98'
# req = Request(url,headers={'User-Agent':'Mozila/5.0'})
# webpage = urlopen(req)
# soup = BeautifulSoup(webpage, 'html.parser')
# resultList = soup.find_all("li", "basicList_item__0T9JD")
#
# for i in resultList:
#     nameList = i.find("div", "basicList_title__VfX3c").find("a")
#     name = nameList
#     priceList = i.find("div", "basicList_price_area__K7DDT")
#     price = priceList.find("span", "price_num__S2p_v")
#     print(name.text, price.text)