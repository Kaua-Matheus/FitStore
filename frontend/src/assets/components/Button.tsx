import type { ReactNode } from "react";
// import { useState } from "react";

interface ButtonProps {
    children?: ReactNode,
    onClick?: () => void,
}

export default function Button({
    children,
    onClick
}: ButtonProps) {

    // const [active, setActive] = useState(false);

    return (

        // Aumentar o tamanho dos textos
        <button className={`
            h-8 w-12 text-sm
            md:h-[4vh] md:w-[6vw] md:text-md
        
            text-custom-extra-dark hover:text-custom-extra-dark
            bg-custom-extra-light hover:bg-custom-primary
            active:bg-custom-secondary transition-colors duration-400
            text-center p-2 rounded
            `}
        onClick={onClick}>
            {children || "Default Button"}
        </button>
    )
}