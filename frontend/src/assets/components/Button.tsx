import type { ReactNode } from "react";
import { useState } from "react";

interface ButtonProps {
    children?: ReactNode,
}

export default function Button({
    children,
}: ButtonProps) {

    const [active, setActive] = useState(false);

    return (

        // Fazer alteração para adicionar sistema de clique e alteração
        <button className={`text-center p-1 rounded ${active ? "bg-custom-extra-light" : "bg-custom-extra-dark"}`}
        onClick={() => {
            setActive(!active);
        }}>
            {children || "Default Button"}
        </button>
    )
}