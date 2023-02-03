<!DOCTYPE html>
<html>
<style>
	/*
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
		text-align: center;
		box-shadow: 0 0 10px #ddd;
	}

	h1 {
		margin-top: 0;
		margin-bottom: 20px;
		color: #333;
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
	*/
</style>
<head>
  <title>Currency API Guide</title>
</head>
<body>
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
			<p><strong>Note:</strong> These commands should be executed in the terminal while inside the project directory.</p>
		</div>
		<p>A big thank you to everyone who helped make this API possible, especially <a href="https://github.com/baibikov">https://github.com/baibikov</a> for giving this task to his brothers.</p>
<h4>
  <h2>Task</h2>
  <p>Implement a currency exchange service API</p>

<h2>Requirements</h2>
  <ul>
    <li>The user can create a currency pair</li>
    <li>The user can transfer currency from one pair to another (only currencies that are in the system, otherwise an error must be thrown).</li>
  </ul>

  <p></p>
  <p>At initialization, a worker (goroutine) must be launched, which must run once in N time ticks (for example, once an hour, etc.), go through all the records in the table and update their ratios (taken from the internet. You can parse the page, you can find the API. The second option is better).</p>
  <p>PostgreSQL should be used as the database.</p>
  <p>The table should look like this:</p>
  <img src="https://user-images.githubusercontent.com/80615643/216640477-76b2f241-244f-4330-bea8-7d501d62ab9a.png" alt="Image">


  <table>
    <tr>
      <th>currency_from</th>
      <th>currency_to</th>
      <th>rate</th>
      <th>updated_at</th>
    </tr>
    <tr>
      <td>USD</td>
      <td>RUB</td>
      <td>75</td>
      <td>2020-09-23 09:13:00+00</td>
    </tr>
  </table>
  <p>The fiber must be used as the HTTP server library/framework.</p>
  <h2>API URLs</h2>
  <ol>
    <li>
      <h3>Creating a record</h3>
      <p>POST /api/currency</p>
      <p>{ "currencyFrom": "USD", "currencyTo": "RUB" }</p>
    </li>
    <li>
      <h3>Converting a value from one currency to another</h3>
      <p>PUT /api/currency</p>
      <p>{ "currencyFrom": "USD", "currencyTo": "RUB", "value": 1 }</p>
    </li>
    <li>
      <h3>Aggregation of added currency pairs</h3>
      <p>GET /api/currency</p>
      <p>[{ "currencyFrom": "USD", "currencyTo": "RUB" }, ….]</p>
    </li>
  </ol></h4>
    
</body>
</html>
