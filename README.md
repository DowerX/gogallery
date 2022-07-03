# goGallery
Basic gallery host written in Golang.

## Setup
Environment:
- GALLERYROOT: path to the files you want to host ("/gallery")
- GALLERYPORT: port of the server (":6874")
- GALLERYDIRSTYLE: XSL template file for directories
- GALLERYIMAGESTYLE: XSL template file for files

You will need to host the static files separately, this program only generates the file listings in XML.
You can do that with [nginx](https://www.nginx.com/) for example.

I recommend using [docker](https://www.docker.com/), with the provided Dockerfile and docker-compose.yml.

## Example output
### Directories
```
<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" href="/static/dirs.xsl"?>
<dirs>
<Directory>
 <Name>My Images and videos</Name>
 <Count>3</Count>
</Directory>
<Directory>
 <Name>Other Folder</Name>
 <Count>2</Count>
</Directory>
</dirs>
```
### Files
```
<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" href="/static/images.xsl"?>
<images>
<File>
  <Path>Other Folder/some_image.jpg</Path>
  <Type>image</Type>
</File>
<File>
  <Path>Other Folder/some_video.jpg</Path>
  <Type>video</Type>
</File>
</images>
```