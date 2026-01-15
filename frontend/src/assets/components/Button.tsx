import type { ReactNode } from "react";
// import { useState } from "react";

interface ButtonProps {
    children?: ReactNode,
    onClick?: () => void,
    className?: string,
    type?: "primary" | "secondary" | "light" | "dark"
}

export default function Button({
    children,
    onClick,
    className,
    type,
}: ButtonProps) {

    switch (type) {
        case "primary":
            return (

            <button className={`
                    w-32
                    text-custom-light bg-custom-primary font-bold
                    text-center p-2 rounded-sm shadow-md
                    active:shadow-sm active:bg-custom-secondary
                    transition-colors duration-400
                    ` + ` ${className}`}
                onClick={onClick}>
                    {children || "Primary"}
                </button>
            )
            
        case "secondary":
            return (

            <button className={`
                    w-32
                    text-custom-dark bg-custom-light-gray font-bold
                    text-center p-2 rounded-sm
                    hover:shadow-md
                    active:shadow-sm active:bg-custom-secondary
                    transition-colors duration-800
                    ` + ` ${className}`}
                onClick={onClick}>
                    {children || "Secondary"}
                </button>
            )

        case "light":
            return (

            <button className={`
                    w-32
                    text-custom-dark bg-custom-light-gray font-bold
                    text-center p-2 rounded-sm shadow-md
                    border border-custom-gray
                    active:shadow-sm active:bg-custom-secondary
                    transition-colors duration-400
                    ` + ` ${className}`}
                onClick={onClick}>
                    {children || "Light"}
                </button>
            )

        case "dark":
            return (

            <button className={`
                    w-32
                    text-custom-light-gray bg-custom-dark font-bold
                    text-center p-2 rounded-sm shadow-md
                    border border-custom-gray
                    active:shadow-sm active:bg-custom-secondary
                    transition-colors duration-400
                    ` + ` ${className}`}
                onClick={onClick}>
                    {children || "Dark"}
                </button>
            )

        default:
            return (
            <button className={`
                    w-32
                    text-custom-light bg-custom-primary font-bold
                    text-center p-2 rounded-sm shadow-md
                    active:shadow-sm active:bg-custom-secondary
                    transition-colors duration-400
                    ` + ` ${className}`}
                onClick={onClick}>
                    {children || "Primary"}
                </button>
            )

    }
}