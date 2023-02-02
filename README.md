<!DOCTYPE html>
<html>
<head>

</head>
<body>
  <h1>Currency API</h1>
<p><strong>Task:</strong> Implement a currency exchange API service</p>
<p><strong>Requirements:</strong></p>
<ul>
  <li>User can create a currency pair</li>
  <li>User can convert currency from one pair to another (only those currencies that are in the system, otherwise an error must be thrown).</li>
</ul>
<p>On initialization, a worker (goroutine) should start that runs once every N time units (for example, once an hour, etc.), and goes through all records in the table and updates their ratios (taken from the Internet. You can parse the page, you can find the API. The second option is better).</p>
<p>PostgreSQL should be used as the database.</p>
<p>The table should look like this:</p>
<table>
  <thead>
    <tr>
      <th>currency_from</th>
      <th>currency_to</th>
      <th>ratio</th>
      <th>updated_at</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>USD</td>
      <td>RUB</td>
      <td>75</td>
      <td>2020-09-23 09:13:00+00</td>
    </tr>
  </tbody>
</table>
<p>Fiber should be used as the HTTP server framework/library.</p>
<p><strong>API URLs:</strong></p>
<ul>
  <li>
    <p><strong>Create a record</strong></p>
    <p>POST /api/currency</p>
    <p>{ "currencyFrom": "USD", "currencyTo": "RUB" }</p>
  </li>
  <li>
    <p><strong>Convert value from one currency to another</strong></p>
    <p>PUT /api/currency</p>
    <p>{ "currencyFrom": "USD", "currencyTo": "RUB", "value": 1 }</p>
  </li>
  <li>
    <p><strong>Aggregate added currency pairs</strong></p>
    <p>GET  /api/currency</p>
    <p>[ { "currencyFrom": "USD", "currencyTo": "RUB" }, ... ]</p>
  </li>
</ul>
<p><strong>Tools used:</strong> Go, PostgreSQL, Fiber</p>
</body>
</html>
