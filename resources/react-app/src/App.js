import React, { Component } from 'react';
import './App.css';
import axios from 'axios';

class App extends Component {

  constructor() {
    super();

    axios.get('/api/json')
      .then(function (response) {
        console.log(response.data.Hobbies[0] === "snowboarding");
        console.log(response.data.Hobbies[0]);
      })
      .catch(function (error) {
        console.log(error);
      });
  }

  render() {
    return (
      <div className="container">
        <div className="row">
          <span>hej</span>
        </div>
        <div className="row">
          <span>hej</span>
        </div>
      </div>
    );
  }
}

export default App;
