const request = require('request');

// Example usage of the 'request' package to fetch a web page
request('https://jsonplaceholder.typicode.com/todos/1', (error, response, body) => {
  if (error) {
    console.error('Error:', error);
    return;
  }
  if (response && response.statusCode === 200) {
    console.log('Response Body:', body);
  } else {
    console.error('Failed to fetch. Status:', response && response.statusCode);
  }
});