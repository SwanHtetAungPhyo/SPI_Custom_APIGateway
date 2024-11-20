const express = require('express');
const cors = require('cors');
const jwt = require('jsonwebtoken');
const bodyParser = require('body-parser');

const app = express();
const PORT = 3002;
const SECRET_KEY = 'your_secret_key'; // Replace with a secure key in a real-world scenario

// Middleware
app.use(cors({ origin: '*' })); // Allow only localhost:8081
app.use(bodyParser.json()); // Parse JSON request bodies

// Hardcoded credentials
const hardcodedUsername = 'admin';
const hardcodedPassword = 'password';

// Login endpoint
app.post('/user/login', (req, res) => {
  try {
    const { username, password } = req.body;

    // Check if credentials are correct
    if (username === hardcodedUsername && password === hardcodedPassword) {
      return res.status(200).send('Login successful');
    } else {
      return res.status(401).json({ message: 'Invalid credentials' });
    }
  } catch (error) {
    console.error("Error during login: ", error);
    return res.status(500).json({ message: 'Server error' });
  }
});

app.get("/user/info",(req, res) =>{
  return res.status(200).send("done");
})
// Start the server
app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
