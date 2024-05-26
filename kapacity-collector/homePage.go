package main

func GetHomePage() string {
	htmlContent := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Kubernetes Capacity Viewer</title>
	</head>
	<body>
		<h1>Welcome to the Kubernetes Capacity Viewer</h1>
		<p>Choose an option to view data:</p>
		<ul>
			<li><a href="/json">View JSON Data</a></li>
		</ul>
	</body>
	</html>`

	return htmlContent
}
