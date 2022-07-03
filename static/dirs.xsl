<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">

<xsl:template match="/">
<html>
    <head>
        <title>Gallery Directories</title>
        <link rel="stylesheet" href="/gallery-static/gallery.css"></link>
    </head>
    <body>
    <div class="container"><xsl:apply-templates/></div>
    </body>
</html>
</xsl:template>

<xsl:template match="Directory">
<a href="/gallery/{Name}"><span class="button"><xsl:value-of select="Name"/></span></a>
</xsl:template>

</xsl:stylesheet>
