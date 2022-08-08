


# fasthttp ![Stars](https://img.shields.io/github/stars/realTristan/GoQueue?color=brightgreen) 

![Go Queue Banner](https://user-images.githubusercontent.com/75189508/183435878-e5669071-df93-478a-a364-245862dadddb.png)

Manageable Queue System implementation for Go.

# fasthttp might not be for you!
fasthttp was design for some high performance edge cases. **Unless** your server/client needs to handle **thousands of small to medium requests per seconds** and needs a consistent low millisecond response time fasthttp might not be for you. **For most cases `net/http` is much better** as it's easier to use and can handle more cases. For most cases you won't even notice the performance difference.
