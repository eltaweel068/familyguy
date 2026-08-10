import requests
from bs4 import BeautifulSoup
import json

headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
}

# proxies = {
#     "http": "http://",
#     "https": "http://"
# }
episodes = {}

for i in range(1, 24):
    season_key = f"season :{i}"
    episodes[season_key] = []
    
    url = f"https://freefamilyguy.com/category/season-{i}/"
    response = requests.get(url, timeout=10, headers=headers)
    soup = BeautifulSoup(response.content, "html.parser")
    # with open("response.html", "w", encoding="utf-8") as f:
    #     f.write(soup.prettify())

    
    for j, episode in enumerate(reversed(soup.find_all("div", class_="entry-header")), start=1):
        
        Episode_Name = episode.find_all("h2")[0].text.strip()
        link = episode.find("a", href=True)["href"]
        
        
        Episode_url = link
        
        response = requests.get(Episode_url, timeout=10, headers=headers)
        soup = BeautifulSoup(response.content, "html.parser")
        # video_element = soup.find('video', class_='first-video')
        video_element = soup.select_one('.wp-block-video video')
        
        if video_element and video_element.get('src'):
            Episode_MP4_link = video_element.get('src')
        else:
            Episode_MP4_link = None

        
        # with open("response_episodeMP4.html", "w", encoding="utf-8") as f:
        #     f.write(soup.prettify())
        
                
        episodes[season_key].append({"Episode": j, "Episode_Name": Episode_Name, "Episode_MP4_link": Episode_MP4_link})

with open("episodes.json", "w", encoding="utf-8") as f:
    json.dump(episodes, f, indent=4)
        