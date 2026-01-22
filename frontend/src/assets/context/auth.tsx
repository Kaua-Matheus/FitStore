import { useContext, createContext, type ReactNode, useState, useEffect } from "react";

interface User {
    userName: string;
}

interface AuthContextType {
    user: User | null;
    isLoading: boolean;
    checkAuth: () => Promise<void>
}

const AuthContext = createContext<AuthContextType | undefined>(undefined)

export function AuthProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<User | null>(null)
    const [isLoading, setIsLoading] = useState(true) // Usamos isLoading para adquirir informações sobre o carregamento e evitar fazer redenrização antes de carregar os dados

    const checkAuth = async () => {
        try {
            const response = await fetch("http://localhost:8080/user/auth", {
                method: "GET",
                credentials: "include",
            });
            const data = await response.json();

            if (data.authenticated) {
                setUser(data.user);
                console.log("Autenticado") // A ideia é utilizarmos a validação para identificar o que será renderizado
            } else {
                setUser(null);
                console.log("Não Autenticado")
            }

        } catch (e) {
            console.log("error trying to authenticate")
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

    return (
        <AuthContext.Provider value={{
            user,
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
        throw new Error('useAuth must be used within an AuthProvider')
    }
    return context;
}