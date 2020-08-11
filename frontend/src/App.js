import React, { Component } from 'react';
import './App.css';
import {useAuth, AuthContext} from './context/auth'
import Main from './components/Main';
import axios from 'axios';

const client = axios.create({crossdomain: true, withCredentials: true});

export class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      authenticated: false,
      username: "None"
    };
  }

  tryAuthenticate() {
    client.get(
      "http://localhost:8080/api/userinfo", 
    ).then(res => {
        this.setState({
          username: res.data.name,
          authenticated: true
        });
    }).catch(err => {
      if (err.response && err.response.status === 401) {
        return;
      }
      console.error("Failed to authenticate: " + err);
    });
  };

  logout() {
    client.post(
      "http://localhost:8080/api/logout"
    ).then(res => {
      this.setState({
        username: "None",
        authenticated: false
      });
    }).catch(err => {
      console.error("Failed to logout: "+err);
    })
  }

  render() {
    

    return (
      <AuthContext.Provider value={{
        username: this.state.username,
        authenticated: this.state.authenticated,
        tryAuthenticate: this.tryAuthenticate.bind(this),
        logout: this.logout.bind(this)
      }}>
          <Main/>
      </AuthContext.Provider>
    )
  }
}

// export const withContext = Component => {
//   return props => {
//       return (
//           <AuthContext.Consumer>
//               {
//                   globalState => {
//                       return (
//                           <Component
//                               {...globalState}
//                               {...props}
//                           />
//                       )
//                   }
//               }
//           </AuthContext.Consumer>
//       )
//   }
// }
