import { User } from "../../models/user";

export interface IAuthService {
    tryAuthenticate(): Promise<User|null>

    logout(): Promise<boolean>
}