import {createContext, useContext} from 'react';

export const AuthContext = createContext({
    username: "",
    authenticated: false,
    login: () => {},
    logout: () => {},
    tryAuthenticate: () => {}
});

