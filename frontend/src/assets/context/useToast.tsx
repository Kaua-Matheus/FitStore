import { createContext, useContext, useState } from "react";
import type { ReactNode } from "react";

interface Toast {
    id: string;
    message: string;
    type?: "success" | "error" | "warning" | "info"
}

interface ToastContextType {
    toasts: Toast[];
    addToast: (message: string, type?: Toast["type"]) => void;
    removeToast: (id: string) => void;
}

const ToastContext = createContext<ToastContextType | undefined>(undefined);

export function ToastProvider({ children }: { children: ReactNode }) {
    const [toasts, setToasts] = useState<Toast[]>([])

    const addToast = (message: string, type: Toast["type"] = "info") => {
        const id = Math.random().toString(36).substring(2, 9);
        const newToast: Toast = { id, message, type }

        setToasts(prev => [...prev, newToast]);

        setTimeout(() => {
            removeToast(id);
        }, 5000) // 5s
    }

    const removeToast = (id: string) => {
        setToasts(prev => prev.filter(toast => toast.id !== id));
    }

    return (
        <ToastContext.Provider value={{ toasts, addToast, removeToast }}>
            {children}
        </ToastContext.Provider>
    )
}

export function useToast() {
    const context = useContext(ToastContext);

    if (!context) {
        throw new Error("useToast must be used within a ToastProvider");
    }
    return context;
}