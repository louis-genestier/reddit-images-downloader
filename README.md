# Reddit Images Downloader

Reddit Images Download is a program that allow users to download images from the top of all time in a given subreddit.

## Usage
You can set the subreddit with `-s` and the number of posts to get with `-l`.
```shell script
./reddit-images-downloader --help
Usage of ./reddit-images-downloader:
  -l int
        number of posts to get (default 100)
  -s string
        name of the subreddit to get images from (default "all")
```

```shell script
./reddit-images-downloader -s GlobalOffensive -l 3
Image from post FaZe Clan vs Cloud9 / ELEAGUE Major Boston 2018 - Grand-Final / Post-Match Discussion on /r/GlobalOffensive downloaded.
Image from post Man speaks the truth. on /r/GlobalOffensive downloaded.
Image from post On Anubis the sprayed B sign is older than the ancient hieroglyphs. on /r/GlobalOffensive downloaded.
```

Output directory is `./images/SUB_NAME/`.
## Contributing
Pull requests are welcome.