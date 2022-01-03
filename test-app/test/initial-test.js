const assert = require('assert');
const request = require('request');
require('dotenv').config()

describe('Initial test', function () {
    it('test call service', function () {
        assert.equal(true, true);
        request.post({url: 'http://' + process.env.HOST_ORDER, formData: null}, function (error, response, body) {
            console.error('error:', error); // Print the error if one occurred
            console.log('statusCode:', response && response.statusCode); // Print the response status code if a response was received
            console.log('body:', body); // Print the HTML for the Google homepage.
        });
    });
});