```
git clone https://github.com/MichaelVareikis/file-emulator.git
cd file-emulator
docker build -t fileapi .
docker run --restart=unless-stopped -d -p 8080:8080 fileapi
```