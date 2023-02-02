<!DOCTYPE html>
<html>
<head>
  <title>Currency API</title>
</head>
<body>
  <h1>Currency API</h1>
  <h2>Task</h2>
  <p>Implement a currency exchange service API</p>
  <h2>Requirements</h2>
  <ul>
    <li>The user can create a currency pair</li>
    <li>The user can transfer currency from one pair to another (only currencies that are in the system, otherwise an error must be thrown).</li>
  </ul>
  <p>At initialization, a worker (goroutine) must be launched, which must run once in N time ticks (for example, once an hour, etc.), go through all the records in the table and update their ratios (taken from the internet. You can parse the page, you can find the API. The second option is better).</p>
  <p>PostgreSQL should be used as the database.</p>
  <p>The table should look like this:</p>
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
  </ol>
</body>
</html>
