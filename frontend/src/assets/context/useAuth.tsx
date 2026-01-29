// React
import { useContext, createContext, type ReactNode, useState, useEffect } from "react";

// Handler
import GetLocalIp from "../handler/http";

interface User {
    "user_login": string;
}

interface AuthContextType {
    user: User | null;
    isLoggedIn: boolean;
    isLoading: boolean;
    checkAuth: () => Promise<void>
}

const AuthContext = createContext<AuthContextType | undefined>(undefined)

export function AuthProvider({ children }: { children: ReactNode }) {
    
    const [user, setUser] = useState<User | null>(null)
    const [isLoading, setIsLoading] = useState(true)

    const checkAuth = async () => {
        try {
            const response = await fetch(`http://${GetLocalIp()}:8080/user/auth`, {
                method: "GET",
                credentials: "include",
            });
            const data = await response.json();

            if (data.authenticated) {
                setUser(data.user);
                return
            } else {
                setUser(null);
                return
            }

        } catch (e) {
            console.log(`error trying to authenticate: ${e}`)
        }
        
    }

    useEffect(() => {
        const initAuth = async () => {
            setIsLoading(true);
            await checkAuth();
            setIsLoading(false);
        };

        initAuth()
    }, [])

    const isLoggedIn = !!user

    return (
        <AuthContext.Provider value={{
            user,
            isLoggedIn,
            isLoading,
            checkAuth,
        }}>
            { children }
        </AuthContext.Provider>
    )
}

export function useAuth() {
    const context = useContext(AuthContext)
    if (!context) {
        throw new Error("useAuth must be used within an AuthProvider")
    }
    return context;
}