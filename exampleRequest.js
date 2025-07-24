const axios = require('axios');

// Example usage of the 'axios' package to fetch a web page
axios.get('https://jsonplaceholder.typicode.com/todos/1')
  .then(response => {
    if (response && response.status === 200) {
      console.log('Response Body:', response.data);
    } else {
      console.error('Failed to fetch. Status:', response && response.status);
    }
  })
  .catch(error => {
    console.error('Error:', error);
  });