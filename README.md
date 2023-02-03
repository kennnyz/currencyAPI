<!DOCTYPE html>
<html>
<head>
  <title>Currency API Guide</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			margin: 0;
			padding: 0;
			background-color: #f2f2f2;
		}
    	.container {
		max-width: 800px;
		margin: 0 auto;
		background-color: white;
		padding: 20px;
		text-align: left;
		box-shadow: 0 0 10px #ddd;
	}
		h1 {
	margin-top: 0;
	margin-bottom: 20px;
	color: #333;
	text-align: center;
}

p {
	margin-bottom: 20px;
	color: #333;
	line-height: 1.5;
}

pre {
	background-color: #f7f7f7;
	padding: 10px;
	border-radius: 5px;
	margin-bottom: 20px;
}

.note {
	background-color: #fcf8e3;
	border-left: 5px solid #faebcc;
	padding: 10px;
	margin-bottom: 20px;
}
table {
	width: 100%;
	margin-top: 20px;
	border-collapse: collapse;
}

th, td {
	border: 1px solid #ddd;
	padding: 10px;
	text-align: left;
}

th {
	background-color: #ddd;
}
</style>
</head>
<body>
  <div class="container">
	  <h1>Currency API</h1>
	    <p>CurrencyAPI is a simple exchange rate API service built with Go, using PostgreSQL as a database and Fiber as a http server framework. </p>
<p>Every exchange rate is retrieved from <a href="https://openexchangerates.org/">openexchangerates.org</a> and updated periodically with a worker running in the background. </p>
<p>To run the database, use the following command in the terminal:</p>
		<pre>make docker</pre>
		<p>This command will run the postgres container and make it ready to use.</p>
		<p>To create the database, use the following command:</p>
		<pre>make createdb</pre>
		<p>To delete the database, use the following command:</p>
		<pre>make dropdb</pre>
		<p>For database migration, use the following commands:</p>
		<pre>make migratedown</pre>
		<pre>make migrateup</pre>
		<div class="note">
			<p><strong>Note:</strong> These commands should be executed in the terminal while inside the project directory.</p
