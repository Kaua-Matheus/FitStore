import type { ReactNode } from "react";
// import { useState } from "react";

interface ButtonProps {
    children?: ReactNode,
    onClick?: () => void,
    type?: "primary" | "secondary"
}

export default function Button({
    children,
    onClick,
    type,
}: ButtonProps) {

    if (type == "primary" || type == undefined) {
        return (

            <button className={`
                w-32
                text-custom-light bg-custom-primary font-bold
                text-center p-2 rounded-sm shadow-md
                active:shadow-sm active:bg-custom-secondary
                transition-colors duration-400
                `}
            onClick={onClick}>
                {children || "Primary"}
            </button>
        )
    } else {
        return (

        <button className={`
            w-32
            text-custom-dark bg-custom-light-gray font-bold
            text-center p-2 rounded-sm shadow-md
            active:shadow-sm active:bg-custom-secondary
            transition-colors duration-800
            `}
        onClick={onClick}>
            {children || "Secondary"}
        </button>
    )
    }
}