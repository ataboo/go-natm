import React, { Component } from 'react';
import './App.css';
import {useAuth, AuthContext} from './context/auth'
import Main from './components/Main';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      userAuthenticated: false
    };
  }

  render() {
    return (
      <AuthContext.Provider value={this.state.user}>
        <Main/>
      </AuthContext.Provider>
    )
  }
}

export default App;
