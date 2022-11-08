const express = require('express')
const cors = require('cors')
const bodyParser = require('body-parser')
const app = express()
const port = 3000

var Pusher = require('pusher');
var pusher = new Pusher({
  appId: '1494118',
  key: 'e48e401e9da913009f95',
  secret: '8bf1185ce7e6c74b65dd',
  cluster: 'ap1'
});

app.use(express.static(__dirname + '/public'));
app.set('view engine', 'ejs');

app.use(cors())
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

app.get('/', function(req, res) {
  res.render('index');
});

app.post('/pusher/auth', function(req, res) {
  console.log(req.body);
  
  var presenceData = {
    user_id: req.body.socket_id,
  };
  var auth = pusher.authorizeChannel(
    req.body.socket_id,
    req.body.channel_name,
    presenceData
  );
  res.send(auth);
});

app.listen(port, function () {
  console.log(`Server is running on port ${port}`);
});