import { IAuthService } from "../interface/iauthservice";
import { User } from "../../models/user";
import axios, { AxiosInstance } from "axios";

export class AuthService implements IAuthService {
    user: User | null;
    client: AxiosInstance;
    hostUri: string
    
    
    constructor() {
        this.user = null;
        this.client = axios.create({
            withCredentials: true,
        });

        this.hostUri = "http://localhost:8080/api/v1/";
    }

    async tryAuthenticate(): Promise<User|null> {       
        try {
            const response = await this.client.get(`${this.hostUri}userinfo`)

            return {
                id: response.data.id ?? "Unset",
                name: response.data.name ?? "Unset",
                email: response.data.email ?? "Unset"
            };
        } catch (err) {
            if (err.response && err.response.status === 401) {
                return null;
            }
            console.error("Failed to authenticate: " + err);
        }

        return null;
    }

    async logout(): Promise<boolean> {
        try {
            await this.client.post(`${this.hostUri}logout`);
            return true;
        } catch(err) {
            console.error("Failed to logout: "+err);
        }
        
        return false;
    }
}