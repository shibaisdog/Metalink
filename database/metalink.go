package database

import "fmt"

func New_MetaLink(title, description, sitename, siteurl, sitetype, color string, image []string) string {
	Image, Twitter_Image := "", ""
	for _, image_url := range image {
		Image += fmt.Sprintf(`<meta property="og:image" content="%s"/>\n`, image_url)
		Twitter_Image += fmt.Sprintf(`
			<meta property="twitter:image" content="%s"/>
			<meta name="twitter:card" content="summary_large_image">
		`, image_url)
	}
	html := fmt.Sprintf(`
	<!DOCTYPE html>
    <html lang="en">
		<head>
			<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
			<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<meta name="robots" content="all">

			<!-- Meta Tags -->
			<meta name="color" content="%s"/>
			<meta name="theme-color" content="%s"/>
			<title>%s</title>
			<meta name="title" content="%s"/>
			<meta name="description" content="%s"/>
			<meta name="url" content="%s">
			<meta name="site_name" content="%s">
			<link rel="canonical" href="%s">

			<!-- Open Graph -->
			<meta property="og:color" content="%s"/>
			<meta property="og:type" content="%s"/>
			<meta property="og:url" content="%s"/>
			<meta property="og:title" content="%s"/>
			<meta property="og:description" content="%s"/>
			<meta property="og:site_name" content="%s">
			%s

			<!-- Twitter -->
			<meta property="twitter:card" content="summary_large_image"/>
			<meta property="twitter:url" content="%s"/>
			<meta property="twitter:title" content="%s"/>
			<meta property="twitter:description" content="%s"/>
			%s

			<!-- Redirect -->
			%s
		</head>
		<body>
		    <h2>warning! When you click, you will be directed to the corresponding URL : %s</h2>
			<a href="%s">Redirect</a>
		</body>
    </html>
	`, color, color, title, title, description, siteurl, sitename, siteurl,
		color, sitetype, siteurl, title, description, sitename, Image,
		siteurl, title, description, Twitter_Image, siteurl, siteurl, siteurl)
	return html
}
