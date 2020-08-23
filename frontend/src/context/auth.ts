import { createContext } from "react";
import {User} from "../models/user"

interface AuthContextProps {
    currentUser: User|null;
    logout: () => Promise<boolean>
}

export const AuthContext = createContext<AuthContextProps>({
    currentUser: null,
    logout: async () => false
})