const express = require('express')
const cors = require('cors')
const app = express()
const port = 3000

app.use(express.static(__dirname + '/public'));
app.set('view engine', 'ejs');

app.get('/', cors(), function(req, res) {
  res.render('index');
});

app.listen(port, function () {
  console.log(`Server is running on port ${port}`);
});