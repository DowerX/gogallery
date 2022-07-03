<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">

<xsl:template match="/">
<html>
    <head>
        <title>Gallery</title>
        <link rel="stylesheet" href="/gallery-static/gallery.css"></link>
        <script src="/jquery.js"/>
    </head>
    <body>
    <div class="container"><xsl:apply-templates/></div>
    <script src="/gallery-static/videos.js"/>
    </body>
</html>
</xsl:template>

<xsl:template match="File">
    <a href="/gallery-static/{Path}">
        <xsl:if test="Type='image'"><img src="/gallery-static/{Path}"/></xsl:if>
        <xsl:if test="Type='video'"><video allow="autoplay" muted="true" loop="true" preload="metadata"><source src="/gallery-static/{Path}#t=0.5"/></video></xsl:if>
        <span class="path"><xsl:value-of select="Path"/></span>
    </a>
</xsl:template>

</xsl:stylesheet>
