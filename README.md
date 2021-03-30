# Sitemap-Builder
Sitemap-Builder

A sitemap is basically a map of all of the pages within a specific domain. They are used by search engines and other tools to inform them of all of the pages on your domain.

One way these can be built is by first visiting the root page of the website and making a list of every link on that page that goes to a page on the same domain. For instance, on https://golang.org/ you might find a link to https://golang.org/doc/ along with several other links.

In this project the goal is to build a sitemap builder like the one described above. 

Once you have determined all of the pages of a site, sitemap builder should then output the data in the following XML format:

```
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>http://www.example.com/</loc>
  </url>
  <url>
    <loc>http://www.example.com/dogs</loc>
  </url>
</urlset>
```

Note: This is same as the standard sitemap protocol

Where each page is listed in its own <url> tag and includes the <loc> tag inside of it.
