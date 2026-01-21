// React Icons
import { IoCloseCircleOutline } from "react-icons/io5";
import { useToast } from "../context/useToast";
import { useState } from "react";

interface ToastCardProps {
    id: string;
    message: string;
    type?: "success" | "error" | "warning" | "info"
}

function ToastCard({ id, message, type }: ToastCardProps) {
    const { removeToast } = useToast();
    const [isExisting, setIsExiting] = useState(true)

    setTimeout(()=>{
        setIsExiting(false)
    }, 4200) // Timeout para animação de fechamento de componente

    const getToastColor = () => {
        switch (type) {
            case 'success': return 'bg-green-500';
            case 'error': return 'bg-red-500';
            case 'warning': return 'bg-yellow-500';
            default: return 'bg-gray-500';
        }
    }

    return (
        <div className={`fixed top-[10vh] right-[1vw] h-10 w-80 ${getToastColor()} rounded-md text-center flex items-center justify-between px-4
            ${isExisting ? 'animate-slide-in-right' : 'animate-slide-out-right'}`}>
            <span className="text-white">{message}</span>
            <IoCloseCircleOutline
                size={25}
                onClick={() => removeToast(id)}
                className="cursor-pointer text-white hover:text-red-300 transition-colors" />
        </div>
    )
}

export default function ToastContainer() {
    const { toasts } = useToast();

    return (
        <div className="fixed top-0 right-0 z-50 space-y-2">
            {toasts.map((toast) => (
                <ToastCard
                    key={toast.id}
                    id={toast.id}
                    message={toast.message}
                    type={toast.type}
                 />
            ))}
        </div>
    )
}