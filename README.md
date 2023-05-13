<!DOCTYPE html>
<html>

<head>
  <title>Album CRUD API</title>
</head>

<body>
  <h1>Album CRUD API</h1>
  <p>This is an API application for an album CRUD.</p>

  <h2>Routes</h2>
  <ul>
    <li><strong>POST /albums</strong>: Create a new album.</li>
    <li><strong>GET /albums</strong>: Retrieve all albums.</li>
    <li><strong>GET /albums/{id}</strong>: Retrieve a specific album by ID.</li>
    <li><strong>PUT /albums/{id}</strong>: Update a specific album by ID.</li>
    <li><strong>DELETE /albums/{id}</strong>: Delete a specific album by ID.</li>
  </ul>

  <h2>Requests</h2>
  <p>Requests should be made using the JSON format.</p>

  <h3>DELETE Request Example:</h3>
<pre>
DELETE /albums/{id}
</pre>
<h3>GET All Albums Request Example:</h3>
<pre>
GET /albums
</pre>
<h3>GET Single Album Request Example:</h3>
<pre>
GET /albums/{id}
</pre>
<h3>POST Request Example:</h3>
<pre>
POST /albums
Content-Type: application/json
{
"title": "Album Title",
"artist": "Artist Name",
"price": 9.99
}
</pre>

<h3>PATCH Request Example:</h3>
<pre>
PATCH /albums/{id}
Content-Type: application/json
{
"title": "Updated Album Title",
"price": 12.99
}
</pre>

  <h2>Database Configuration</h2>
  <p>Before running the application, make sure to configure the database connection information in the <code>src/db/db.go</code> file.</p>

  <h2>Running the Application</h2>
  <p>To run the application, you need to have Go installed in your environment. Then, execute the following command:</p>
  <pre>
go run main.go
  </pre>

  <p>The application will run on port 5001 by default.</p>
</body>

</html>
